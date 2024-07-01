package core

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wazero"

	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/black-pepper-team/community-indexer/contracts"
	"github.com/black-pepper-team/community-indexer/internal/data"
)

var (
	ErrContractNotFound        = fmt.Errorf("contract not found")
	ErrCreateWitnessCalculator = errors.New("failed to create witness calculator")
)

const (
	Deadline  = 1720092721
	Timestamp = 1720092421
)

func (c *Core) GetCommunitiesList() ([]data.Community, error) {
	communities, err := c.db.New().CommunitiesQ().Select()
	if err != nil {
		c.log.WithError(err).Error("Failed to get communities list")
		return nil, err
	}

	return communities, nil
}

func (c *Core) GetCommunityById(communityId uuid.UUID) (*data.Community, error) {
	community, err := c.db.New().CommunitiesQ().WhereID(communityId).Get()
	if err != nil {
		c.log.WithError(err).Error("Failed to get community by ID")
		return nil, err
	}

	if community == nil {
		return nil, nil
	}

	return community, nil
}

func (c *Core) CreateCommunity(
	collectionName string,
	collectionSymbol string,
	privKey *ecdsa.PrivateKey,
) (*data.Community, error) {
	var err error
	var tx *types.Transaction
	var address common.Address
	err = c.retryChainCall(c.ctx, privKey, func(signer *bind.TransactOpts) error {
		address, tx, _, err = contracts.DeployERC721Mock(signer, c.ethClient, collectionName, collectionSymbol)
		if err != nil {
			return fmt.Errorf("failed to deploy new community: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	communityId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	newCommunity := data.Community{
		ID:              communityId,
		Status:          data.Deploying,
		Name:            collectionName,
		Symbol:          collectionSymbol,
		ContractAddress: address,
		OwnerAddress:    crypto.PubkeyToAddress(privKey.PublicKey),
	}

	if err = c.db.New().CommunitiesQ().Insert(&newCommunity); err != nil {
		return nil, fmt.Errorf("failed to insert new community: %w", err)
	}

	go func() {
		nextStatus := data.Ready

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = data.FailedToDeploy
		}

		newCommunity.Status = nextStatus
		if newErr := c.db.New().CommunitiesQ().Update(&newCommunity); newErr != nil {
			c.log.
				WithField("reason", newErr.Error()).
				Errorf("failed to update the community status to failed community")
		}

		return
	}()

	return &newCommunity, nil
}

func (c *Core) ImportCommunity(contractAddress common.Address) (*data.Community, error) {
	existedCommunity, err := c.db.CommunitiesQ().WhereContractAddress(contractAddress).Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get community by contract address: %w", err)
	}

	if existedCommunity != nil {
		return existedCommunity, nil
	}

	collectionContract, err := contracts.NewERC721Mock(contractAddress, c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get ERC721 contract: %w", err)
	}

	collectionName, err := collectionContract.Name(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	collectionSymbol, err := collectionContract.Symbol(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	contractOwner, err := collectionContract.Owner(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	communityId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	newCommunity := &data.Community{
		ID:              communityId,
		Status:          data.Ready,
		Name:            collectionName,
		Symbol:          collectionSymbol,
		ContractAddress: contractAddress,
		OwnerAddress:    contractOwner,
	}

	if err = c.db.New().CommunitiesQ().Insert(newCommunity); err != nil {
		return nil, fmt.Errorf("failed to insert new community: %w", err)
	}

	return newCommunity, nil
}

func (c *Core) RegisterInCommunity(
	bjjPublicKey babyjub.PublicKey,
	nftOwner common.Address,
	contractId common.Address,
	nftId big.Int,
	privKey *ecdsa.PrivateKey,
) (*RegisterRequest, error) {
	if c.registeredUsers[nftOwner] != nil {
		registerMetadata := c.registeredUsers[nftOwner][contractId]
		if registerMetadata != nil {
			registryRequest := c.registerRequests[*registerMetadata.Id]

			if registryRequest.Status == Processing || registryRequest.Status == Registered {
				return registryRequest, nil
			}
		}
	}

	wtnsCalculator, err := witness.NewCalculator(
		c.circuits[VerifiableCommitmentCircuitName][WASM],
		witness.WithWasmEngine(wazero.NewCircom2WZWitnessCalculator),
	)
	if err != nil {
		return nil, ErrCreateWitnessCalculator
	}

	transitionInputs := VarifiableCommitmentInputs{
		ContractId:     contractId.String(),
		NftId:          nftId.String(),
		NftOwner:       nftOwner.String(),
		Deadline:       strconv.FormatInt(Deadline, 10),
		Timestamp:      strconv.FormatInt(Timestamp, 10),
		BabyJubJubPKAx: bjjPublicKey.X.String(),
		BabyJubJubPKAy: bjjPublicKey.Y.String(),
	}

	transitionInputsRaw, err := json.Marshal(transitionInputs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transition inputs: %w", err)
	}

	parsedInputs, err := witness.ParseInputs(transitionInputsRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse witness inputs: %w", err)
	}

	wtnsBytes, err := wtnsCalculator.CalculateWTNSBin(parsedInputs, true)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate witnesses: %w", err)
	}

	rapidProof, err := prover.Groth16Prover(c.circuits[VerifiableCommitmentCircuitName][ZKEY], wtnsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prov with groth16: %w", err)
	}

	contractZKP, err := parseZKPArgs(&ZKProof{
		A:        rapidProof.Proof.A,
		B:        rapidProof.Proof.B,
		C:        rapidProof.Proof.C,
		Protocol: rapidProof.Proof.Protocol,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse ZKProof to contract readable: %w", err)
	}

	credentialIdBigInt, ok := new(big.Int).SetString(rapidProof.PubSignals[0], 10)
	if !ok {
		return nil, errors.New("failed to parse credential ID")
	}

	var tx *types.Transaction
	err = c.retryChainCall(c.ctx, privKey, func(signer *bind.TransactOpts) error {
		tx, err = c.authStorageContract.Register(
			signer,
			contractId,
			&nftId,
			nftOwner,
			[32]byte(credentialIdBigInt.Bytes()),
			big.NewInt(Deadline),
			*contractZKP,
		)
		if err != nil {
			return fmt.Errorf("failed to call Register: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	requestId := uuid.New()
	registerRequest := RegisterRequest{
		Id:     requestId,
		Status: Processing,
	}

	c.registerRequests[requestId] = &registerRequest

	if c.registeredUsers[nftOwner] == nil {
		c.registeredUsers[nftOwner] = make(map[common.Address]*Metadata)
	}
	c.registeredUsers[nftOwner][contractId] = &Metadata{
		Id:           &requestId,
		CredentialID: credentialIdBigInt,
	}

	go func() {
		nextStatus := Registered

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = FailedRegister
		}

		registerRequest.Status = nextStatus
		c.registerRequests[requestId] = &registerRequest

		return
	}()

	return &registerRequest, nil
}

func (c *Core) GetRegister(registerRequestId uuid.UUID) (*RegisterRequest, error) {
	registerRequest := c.registerRequests[registerRequestId]

	if registerRequest == nil {
		return nil, nil
	}

	return registerRequest, nil
}

func (c *Core) AddCommunityParticipant(
	privKey *ecdsa.PrivateKey,
	participantAddress common.Address,
	contractAddress common.Address,
) (*RegisterRequest, error) {
	community, err := c.db.CommunitiesQ().WhereContractAddress(contractAddress).Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get community by contract address: %w", err)
	}

	if community == nil {
		return nil, ErrContractNotFound
	}

	collectionContract, err := contracts.NewERC721Mock(contractAddress, c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get ERC721 contract: %w", err)
	}

	var tx *types.Transaction
	err = c.retryChainCall(c.ctx, privKey, func(signer *bind.TransactOpts) error {
		newRandId := babyjub.NewRandPrivKey()
		tx, err = collectionContract.Mint(signer, participantAddress, (newRandId).Scalar().BigInt())
		if err != nil {
			return fmt.Errorf("failed to call Register: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	requestId := uuid.New()
	mintRequest := RegisterRequest{
		Id:     requestId,
		Status: Processing,
	}

	c.registerRequests[requestId] = &mintRequest

	go func() {
		nextStatus := Registered

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = FailedRegister
		}

		mintRequest.Status = nextStatus
		c.registerRequests[requestId] = &mintRequest

		return
	}()

	return &mintRequest, nil
}

func (c *Core) SendMessage(
	bjjPrivKey babyjub.PrivateKey,
	nftId *big.Int,
	nftOwner common.Address,
	contractId common.Address,
	message string,
) (*RegisterRequest, error) {
	wtnsCalculator, err := witness.NewCalculator(
		c.circuits[PostMessageCircuitName][WASM],
		witness.WithWasmEngine(wazero.NewCircom2WZWitnessCalculator),
	)
	if err != nil {
		return nil, ErrCreateWitnessCalculator
	}

	hashInt := new(big.Int).SetBytes(crypto.Keccak256([]byte(message)))

	mask := new(big.Int).SetBytes([]byte{0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	result := new(big.Int).And(hashInt, mask)

	if c.registeredUsers[nftOwner] == nil {
		return nil, errors.New("user is not registered")
	}

	if c.registeredUsers[nftOwner][contractId] == nil {
		return nil, errors.New("user is not registered in this community")
	}

	credentialID, err := buildCredentialId(
		contractId,
		nftId,
		nftOwner,
		bjjPrivKey.Public().X,
		bjjPrivKey.Public().Y,
		Timestamp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build credential ID: %w", err)
	}

	hashedCredentialId, err := poseidon.Hash([]*big.Int{credentialID})
	if err != nil {
		return nil, fmt.Errorf("failed to hash credential ID: %w", err)
	}

	// Convert hash to hexadecimal string
	hashHex := fmt.Sprintf("%064x", hashedCredentialId)

	bytesHashedCredentialId, err := hex.DecodeString(hashHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hash to bytes: %w", err)
	}

	smtProof, err := c.poseidonSMTContract.GetProof(nil, [32]byte(bytesHashedCredentialId))
	if err != nil {
		return nil, fmt.Errorf("failed to get proof from PoseidonSMT contract: %w", err)
	}

	fmt.Println(smtProof.Existence)
	fmt.Println("build credential: ", [32]byte(bytesHashedCredentialId))
	fmt.Println("build credential: ", bytesHashedCredentialId)

	var siblingsHex []string
	for _, sibling := range smtProof.Siblings {
		siblingsHex = append(siblingsHex, hexutil.Encode(sibling[:]))
	}

	auxExistence := "0"
	if smtProof.AuxExistence {
		auxExistence = "1"
	}

	signature := bjjPrivKey.SignPoseidon(result)

	expectedTimestamp := time.Now().Add(500 * time.Second).Unix()
	postInputs := PostMessageInputs{
		ContractId: contractId.String(),
		NftId:      nftId.String(),
		NftOwner:   nftOwner.String(),
		Timestamp:  strconv.FormatInt(Timestamp, 10),

		BabyJubJubPKAx: bjjPrivKey.Public().X.String(),
		BabyJubJubPKAy: bjjPrivKey.Public().Y.String(),

		MessageHash:              result.String(),
		ExpectedMessageTimestamp: strconv.FormatInt(expectedTimestamp, 10),

		Root:        hexutil.Encode(smtProof.Root[:]),
		Siblings:    siblingsHex,
		AuxKey:      hexutil.Encode(smtProof.AuxKey[:]),
		AuxValue:    hexutil.Encode(smtProof.AuxValue[:]),
		AuxIsEmpty:  auxExistence,
		IsExclusion: "0",

		MessageSignatureR8x: signature.R8.X.String(),
		MessageSignatureR8y: signature.R8.Y.String(),
		MessageSignatureS:   signature.S.String(),
	}

	marshalled, _ := json.Marshal(postInputs)
	fmt.Println("Transition inputs:", string(marshalled))

	postInputsInputsRaw, err := json.Marshal(postInputs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transition inputs: %w", err)
	}

	parsedInputs, err := witness.ParseInputs(postInputsInputsRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse witness inputs: %w", err)
	}

	wtnsBytes, err := wtnsCalculator.CalculateWTNSBin(parsedInputs, true)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate witnesses: %w", err)
	}

	rapidProof, err := prover.Groth16Prover(c.circuits[PostMessageCircuitName][ZKEY], wtnsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prov with groth16: %w", err)
	}

	contractZKP, err := parseZKPArgs(&ZKProof{
		A:        rapidProof.Proof.A,
		B:        rapidProof.Proof.B,
		C:        rapidProof.Proof.C,
		Protocol: rapidProof.Proof.Protocol,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse ZKProof to contract readable: %w", err)
	}

	var tx *types.Transaction
	err = c.retryChainCall(c.ctx, c.privateKey, func(signer *bind.TransactOpts) error {
		tx, err = c.chatContract.PostMessage(
			signer,
			contractId,
			message,
			smtProof.Root,
			big.NewInt(expectedTimestamp),
			*contractZKP,
		)
		if err != nil {
			return fmt.Errorf("failed to call Chat: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	requestId := uuid.New()
	registerRequest := RegisterRequest{
		Id:     requestId,
		Status: Processing,
	}

	c.registerRequests[requestId] = &registerRequest

	go func() {
		nextStatus := Registered

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = FailedRegister
		}

		registerRequest.Status = nextStatus
		c.registerRequests[requestId] = &registerRequest

		return
	}()

	return &registerRequest, nil
}

func buildCredentialId(
	contractId common.Address,
	nftId *big.Int,
	nftOwner common.Address,
	babyJubJubPKAx *big.Int,
	babyJubJubPKAy *big.Int,
	timestamp int64,
) (*big.Int, error) {
	// Convert inputs to big.Int
	contractIdBigInt := new(big.Int).SetBytes(contractId.Bytes())

	nftOwnerBigInt := new(big.Int).SetBytes(nftOwner.Bytes())

	timestampBigInt := big.NewInt(timestamp)

	fmt.Println("contractIdBigInt: ", contractIdBigInt)
	fmt.Println("nftId: ", nftId)
	fmt.Println("nftOwnerBigInt: ", nftOwnerBigInt)
	fmt.Println("babyJubJubPKAx: ", babyJubJubPKAx)
	fmt.Println("babyJubJubPKAy: ", babyJubJubPKAy)
	fmt.Println("timestampBigInt: ", timestampBigInt)

	// Poseidon hash
	inputs := []*big.Int{contractIdBigInt, nftId, nftOwnerBigInt, babyJubJubPKAx, babyJubJubPKAy, timestampBigInt}
	hash, err := poseidon.Hash(inputs)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
