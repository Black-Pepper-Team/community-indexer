package core

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/contracts"
)

type Metadata struct {
	CredentialID *big.Int
	Id           *uuid.UUID
}

type Circuits map[string]map[circuitKey][]byte
type RegisterStorage map[common.Address]map[common.Address]*Metadata

type RegisterStatus string

const (
	Registered     RegisterStatus = "registered"
	Processing     RegisterStatus = "processing"
	FailedRegister RegisterStatus = "failed-register"
)

type VarifiableCommitmentInputs struct {
	ContractId     string `json:"contractId"`
	NftId          string `json:"nftId"`
	NftOwner       string `json:"nftOwner"`
	Deadline       string `json:"deadline"`
	BabyJubJubPKAx string `json:"babyJubJubPK_Ax"`
	BabyJubJubPKAy string `json:"babyJubJubPK_Ay"`
	Timestamp      string `json:"timestamp"`
}

type PostMessageInputs struct {
	ContractId string `json:"contractId"`
	NftId      string `json:"nftId"`
	NftOwner   string `json:"nftOwner"`
	Timestamp  string `json:"timestamp"`

	BabyJubJubPKAx string `json:"babyJubJubPK_Ax"`
	BabyJubJubPKAy string `json:"babyJubJubPK_Ay"`

	MessageHash              string `json:"messageHash"`
	ExpectedMessageTimestamp string `json:"expectedMessageTimestamp"`

	Root        string   `json:"root"`
	Siblings    []string `json:"siblings"`
	AuxKey      string   `json:"auxKey"`
	AuxValue    string   `json:"auxValue"`
	AuxIsEmpty  string   `json:"auxIsEmpty"`
	IsExclusion string   `json:"isExclusion"`

	MessageSignatureR8x string `json:"messageSignatureR8x"`
	MessageSignatureR8y string `json:"messageSignatureR8y"`
	MessageSignatureS   string `json:"messageSignatureS"`
}

type RegisterRequest struct {
	Id     uuid.UUID      `json:"id"`
	Status RegisterStatus `json:"status"`
}

// ZKProof is structure that represents SnarkJS library result of proof generation
type ZKProof struct {
	A        []string   `json:"pi_a"`
	B        [][]string `json:"pi_b"`
	C        []string   `json:"pi_c"`
	Protocol string     `json:"protocol"`
}

// FullProof is ZKP proof with public signals
type FullProof struct {
	Proof      *ZKProof `json:"proof"`
	PubSignals []string `json:"pub_signals"`
}

func parseZKPArgs(zkp *ZKProof) (*contracts.VerifierHelperProofPoints, error) {
	a, b, c, err := zkp.ProofToBigInts()
	if err != nil {
		return nil, fmt.Errorf("failed to parse ZKProof to big ints: %w", err)
	}

	// Here we swap array's elements to be
	// compatible with verifier contract
	return &contracts.VerifierHelperProofPoints{
		A: [2]*big.Int{a[0], a[1]},
		B: [2][2]*big.Int{
			{b[0][1], b[0][0]},
			{b[1][1], b[1][0]},
		},
		C: [2]*big.Int{c[0], c[1]},
	}, nil
}

func (p *ZKProof) ProofToBigInts() (a []*big.Int, b [][]*big.Int, c []*big.Int, err error) {
	a, err = arrayStringToBigInt(p.A)
	if err != nil {
		return nil, nil, nil, err
	}

	b = make([][]*big.Int, len(p.B))
	for i, v := range p.B {
		b[i], err = arrayStringToBigInt(v)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	c, err = arrayStringToBigInt(p.C)
	if err != nil {
		return nil, nil, nil, err
	}

	return a, b, c, nil
}

func arrayStringToBigInt(str []string) ([]*big.Int, error) {
	var result []*big.Int

	for i := 0; i < len(str); i++ {
		si, err := stringToBigInt(str[i])
		if err != nil {
			return nil, err
		}

		result = append(result, si)
	}

	return result, nil
}

func stringToBigInt(str string) (*big.Int, error) {
	base := 10
	if bytes.HasPrefix([]byte(str), []byte("0x")) {
		base = 16
		str = strings.TrimPrefix(str, "0x")
	}

	n, ok := new(big.Int).SetString(str, base)
	if !ok {
		return nil, fmt.Errorf("can not parse string to *big.Int: %s", str)
	}

	return n, nil
}
