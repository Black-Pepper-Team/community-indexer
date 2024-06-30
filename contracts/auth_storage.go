// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// AuthenticationStorageMetaData contains all meta data concerning the AuthenticationStorage contract.
var AuthenticationStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"credentialId\",\"type\":\"bytes32\"}],\"name\":\"CredentialExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDeadline\",\"type\":\"uint256\"}],\"name\":\"DeadlineExceedsMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZKProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNFTOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"credentialId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_DEADLINE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credentialRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"verifier_\",\"type\":\"address\"}],\"name\":\"__AuthenticationStorage_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credentialRegistry\",\"outputs\":[{\"internalType\":\"contractPoseidonSMT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftOwner_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"credentialId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credentialRegistry_\",\"type\":\"address\"}],\"name\":\"setCredentialRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier_\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516116726100fd600039600081816106cb015281816106f4015261083a01526116726000f3fe6080604052600436106100c25760003560e01c80638da5cb5b1161007f578063b948c3d411610059578063b948c3d414610223578063bbb9d4b014610243578063c93a985c14610263578063f2fde38b1461028357600080fd5b80638da5cb5b14610188578063a7013707146101c5578063ad3cb1cc146101e557600080fd5b80632b7ac3f3146100c75780634f1ef2861461010457806352d1902d146101195780635437988d1461013c5780636f0e530f1461015c578063715018a614610173575b600080fd5b3480156100d357600080fd5b506001546100e7906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b610117610112366004611156565b6102a3565b005b34801561012557600080fd5b5061012e6102c2565b6040519081526020016100fb565b34801561014857600080fd5b506101176101573660046111fe565b6102df565b34801561016857600080fd5b5061012e622a300081565b34801561017f57600080fd5b50610117610309565b34801561019457600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166100e7565b3480156101d157600080fd5b506101176101e03660046111fe565b61031d565b3480156101f157600080fd5b50610216604051806040016040528060058152602001640352e302e360dc1b81525081565b6040516100fb919061123f565b34801561022f57600080fd5b506000546100e7906001600160a01b031681565b34801561024f57600080fd5b5061011761025e366004611272565b610347565b34801561026f57600080fd5b5061011761027e3660046112ab565b610489565b34801561028f57600080fd5b5061011761029e3660046111fe565b610682565b6102ab6106c0565b6102b482610765565b6102be828261076d565b5050565b60006102cc61082f565b5060008051602061161d83398151915290565b6102e7610878565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b610311610878565b61031b60006108d3565b565b610325610878565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff1660008115801561038d5750825b905060008267ffffffffffffffff1660011480156103aa5750303b155b9050811580156103b8575080155b156103d65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561040057845460ff60401b1916600160401b1785555b61040933610944565b600180546001600160a01b038089166001600160a01b03199283161790925560008054928a1692909116919091179055831561048057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2906020015b60405180910390a15b50505050505050565b610496622a30004261131c565b8211156104c7576040516379b11a2f60e11b815260048101839052622a300060248201526044015b60405180910390fd5b6104d18686610955565b814211156104fb5760405163fb98715560e01b8152336004820152602481018490526044016104be565b610509868686868686610a56565b610526576040516303b2487b60e11b815260040160405180910390fd5b604080516060810182526001600160a01b0380891682526020820188905286168183015290516304b98e1d60e31b815260009173__$03320550cd1b629da90608251571b2532e$__916325cc70e8916105819160040161133d565b602060405180830381865af415801561059e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c2919061136e565b6000546040516368ef2c9560e11b815260048101879052602481018390529192506001600160a01b03169063d1de592a90604401600060405180830381600087803b15801561061057600080fd5b505af1158015610624573d6000803e3d6000fd5b5050604080516001600160a01b03808c168252602082018b905289169181019190915260608101879052608081018690527fdc72f447eb6c32bde8c6dcd2d09f5b297b4d780aaf6ab9fb99d3d083b4b0058c925060a0019050610477565b61068a610878565b6001600160a01b0381166106b457604051631e4fbdf760e01b8152600060048201526024016104be565b6106bd816108d3565b50565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061074757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661073b60008051602061161d833981519152546001600160a01b031690565b6001600160a01b031614155b1561031b5760405163703e46dd60e11b815260040160405180910390fd5b6106bd610878565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156107c7575060408051601f3d908101601f191682019092526107c49181019061136e565b60015b6107ef57604051634c9c8ce360e01b81526001600160a01b03831660048201526024016104be565b60008051602061161d833981519152811461082057604051632a87526960e21b8152600481018290526024016104be565b61082a8383610b62565b505050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461031b5760405163703e46dd60e11b815260040160405180910390fd5b336108aa7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461031b5760405163118cdaa760e01b81523360048201526024016104be565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b61094c610bb8565b6106bd81610c01565b600080836001600160a01b0316636352211e60e01b8460405160240161097d91815260200190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516109bb9190611387565b600060405180830381855afa9150503d80600081146109f6576040519150601f19603f3d011682016040523d82523d6000602084013e6109fb565b606091505b5091509150811580610a325750336001600160a01b031681806020019051810190610a2691906113a3565b6001600160a01b031614155b15610a5057604051631022318760e21b815260040160405180910390fd5b50505050565b60408051600580825260c082019092526000918291906020820160a0803683370190505090508460001c81600081518110610a9357610a936113c0565b602002602001018181525050876001600160a01b031681600181518110610abc57610abc6113c0565b6020026020010181815250508681600281518110610adc57610adc6113c0565b602002602001018181525050856001600160a01b031681600381518110610b0557610b056113c0565b6020026020010181815250508381600481518110610b2557610b256113c0565b6020908102919091010152610b5681610b4336869003860186611426565b6001546001600160a01b03169190610c09565b98975050505050505050565b610b6b82610c30565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115610bb05761082a8282610c95565b6102be610d0d565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661031b57604051631afcd79f60e31b815260040160405180910390fd5b61068a610bb8565b6000610c2684836000015184602001518560400151878851610d2c565b90505b9392505050565b806001600160a01b03163b600003610c6657604051634c9c8ce360e01b81526001600160a01b03821660048201526024016104be565b60008051602061161d83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051610cb29190611387565b600060405180830381855af49150503d8060008114610ced576040519150601f19603f3d011682016040523d82523d6000602084013e610cf2565b606091505b5091509150610d02858383610ebe565b925050505b92915050565b341561031b5760405163b398979f60e01b815260040160405180910390fd5b600080610d3883610f1a565b604051602001610d4891906114c0565b6040516020818303038152906040529050600080896001600160a01b0316838a8a8a604051602401610d7c9392919061155a565b60408051601f198184030181529082905291610d9791611387565b60405180910390206001600160e01b0319166020820180516001600160e01b03838183161783525050505087604051602001610dd49291906115b2565b60408051601f1981840301815290829052610dee91611387565b600060405180830381855afa9150503d8060008114610e29576040519150601f19603f3d011682016040523d82523d6000602084013e610e2e565b606091505b509150915081610e9c5760405162461bcd60e51b815260206004820152603360248201527f566572696669657248656c7065723a206661696c656420746f2063616c6c207660448201527232b934b33ca83937b7b310333ab731ba34b7b760691b60648201526084016104be565b80806020019051810190610eb091906115fa565b9a9950505050505050505050565b606082610ed357610ece82610fad565b610c29565b8151158015610eea57506001600160a01b0384163b155b15610f1357604051639996b31560e01b81526001600160a01b03851660048201526024016104be565b5080610c29565b60606000610f2783610fd6565b600101905060008167ffffffffffffffff811115610f4757610f476110c3565b6040519080825280601f01601f191660200182016040528015610f71576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610f7b57509392505050565b805115610fbd5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106110155772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611041576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061105f57662386f26fc10000830492506010015b6305f5e1008310611077576305f5e100830492506008015b612710831061108b57612710830492506004015b6064831061109d576064830492506002015b600a8310610d075760010192915050565b6001600160a01b03811681146106bd57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156110fc576110fc6110c3565b60405290565b6040805190810167ffffffffffffffff811182821017156110fc576110fc6110c3565b604051601f8201601f1916810167ffffffffffffffff8111828210171561114e5761114e6110c3565b604052919050565b6000806040838503121561116957600080fd5b8235611174816110ae565b915060208381013567ffffffffffffffff8082111561119257600080fd5b818601915086601f8301126111a657600080fd5b8135818111156111b8576111b86110c3565b6111ca601f8201601f19168501611125565b915080825287848285010111156111e057600080fd5b80848401858401376000848284010152508093505050509250929050565b60006020828403121561121057600080fd5b8135610c29816110ae565b60005b8381101561123657818101518382015260200161121e565b50506000910152565b602081526000825180602084015261125e81604085016020870161121b565b601f01601f19169190910160400192915050565b6000806040838503121561128557600080fd5b8235611290816110ae565b915060208301356112a0816110ae565b809150509250929050565b6000806000806000808688036101a08112156112c657600080fd5b87356112d1816110ae565b96506020880135955060408801356112e8816110ae565b94506060880135935060808801359250610100609f198201121561130b57600080fd5b5060a0870190509295509295509295565b80820180821115610d0757634e487b7160e01b600052601160045260246000fd5b60608101818360005b6003811015611365578151835260209283019290910190600101611346565b50505092915050565b60006020828403121561138057600080fd5b5051919050565b6000825161139981846020870161121b565b9190910192915050565b6000602082840312156113b557600080fd5b8151610c29816110ae565b634e487b7160e01b600052603260045260246000fd5b600082601f8301126113e757600080fd5b6113ef611102565b80604084018581111561140157600080fd5b845b8181101561141b578035845260209384019301611403565b509095945050505050565b6000610100828403121561143957600080fd5b6114416110d9565b61144b84846113d6565b8152604084605f85011261145e57600080fd5b611466611102565b8060c086018781111561147857600080fd5b8387015b8181101561149d5761148e89826113d6565b8452602090930192840161147c565b508160208601526114ae88826113d6565b84860152505050508091505092915050565b7f76657269667950726f6f662875696e743235365b325d2c75696e743235365b3281527f5d5b325d2c75696e743235365b325d2c75696e743235365b000000000000000060208201526000825161151e81603885016020870161121b565b615d2960f01b6038939091019283015250603a01919050565b8060005b6002811015610a5057815184526020938401939091019060010161153b565b61010081016115698286611537565b60408083018560005b600281101561159957611586838351611537565b9183019160209190910190600101611572565b505050506115aa60c0830184611537565b949350505050565b6000835160206115c5828583890161121b565b84519184019181860160005b828110156115ed578151855293830193908301906001016115d1565b5092979650505050505050565b60006020828403121561160c57600080fd5b81518015158114610c2957600080fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212200d8ee8d9c2a3d875a90ae9da34cd9dc506f7bbde82ba643b8d5a53ceda5fcdf964736f6c63430008140033",
}

// AuthenticationStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthenticationStorageMetaData.ABI instead.
var AuthenticationStorageABI = AuthenticationStorageMetaData.ABI

// AuthenticationStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthenticationStorageMetaData.Bin instead.
var AuthenticationStorageBin = AuthenticationStorageMetaData.Bin

// DeployAuthenticationStorage deploys a new Ethereum contract, binding an instance of AuthenticationStorage to it.
func DeployAuthenticationStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuthenticationStorage, error) {
	parsed, err := AuthenticationStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthenticationStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuthenticationStorage{AuthenticationStorageCaller: AuthenticationStorageCaller{contract: contract}, AuthenticationStorageTransactor: AuthenticationStorageTransactor{contract: contract}, AuthenticationStorageFilterer: AuthenticationStorageFilterer{contract: contract}}, nil
}

// AuthenticationStorage is an auto generated Go binding around an Ethereum contract.
type AuthenticationStorage struct {
	AuthenticationStorageCaller     // Read-only binding to the contract
	AuthenticationStorageTransactor // Write-only binding to the contract
	AuthenticationStorageFilterer   // Log filterer for contract events
}

// AuthenticationStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthenticationStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthenticationStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthenticationStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthenticationStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthenticationStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthenticationStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthenticationStorageSession struct {
	Contract     *AuthenticationStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AuthenticationStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthenticationStorageCallerSession struct {
	Contract *AuthenticationStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AuthenticationStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthenticationStorageTransactorSession struct {
	Contract     *AuthenticationStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AuthenticationStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthenticationStorageRaw struct {
	Contract *AuthenticationStorage // Generic contract binding to access the raw methods on
}

// AuthenticationStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthenticationStorageCallerRaw struct {
	Contract *AuthenticationStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AuthenticationStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthenticationStorageTransactorRaw struct {
	Contract *AuthenticationStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthenticationStorage creates a new instance of AuthenticationStorage, bound to a specific deployed contract.
func NewAuthenticationStorage(address common.Address, backend bind.ContractBackend) (*AuthenticationStorage, error) {
	contract, err := bindAuthenticationStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorage{AuthenticationStorageCaller: AuthenticationStorageCaller{contract: contract}, AuthenticationStorageTransactor: AuthenticationStorageTransactor{contract: contract}, AuthenticationStorageFilterer: AuthenticationStorageFilterer{contract: contract}}, nil
}

// NewAuthenticationStorageCaller creates a new read-only instance of AuthenticationStorage, bound to a specific deployed contract.
func NewAuthenticationStorageCaller(address common.Address, caller bind.ContractCaller) (*AuthenticationStorageCaller, error) {
	contract, err := bindAuthenticationStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageCaller{contract: contract}, nil
}

// NewAuthenticationStorageTransactor creates a new write-only instance of AuthenticationStorage, bound to a specific deployed contract.
func NewAuthenticationStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthenticationStorageTransactor, error) {
	contract, err := bindAuthenticationStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageTransactor{contract: contract}, nil
}

// NewAuthenticationStorageFilterer creates a new log filterer instance of AuthenticationStorage, bound to a specific deployed contract.
func NewAuthenticationStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthenticationStorageFilterer, error) {
	contract, err := bindAuthenticationStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageFilterer{contract: contract}, nil
}

// bindAuthenticationStorage binds a generic wrapper to an already deployed contract.
func bindAuthenticationStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuthenticationStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthenticationStorage *AuthenticationStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthenticationStorage.Contract.AuthenticationStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthenticationStorage *AuthenticationStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.AuthenticationStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthenticationStorage *AuthenticationStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.AuthenticationStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthenticationStorage *AuthenticationStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthenticationStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthenticationStorage *AuthenticationStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthenticationStorage *AuthenticationStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.contract.Transact(opts, method, params...)
}

// MAXDEADLINE is a free data retrieval call binding the contract method 0x6f0e530f.
//
// Solidity: function MAX_DEADLINE() view returns(uint256)
func (_AuthenticationStorage *AuthenticationStorageCaller) MAXDEADLINE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "MAX_DEADLINE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEADLINE is a free data retrieval call binding the contract method 0x6f0e530f.
//
// Solidity: function MAX_DEADLINE() view returns(uint256)
func (_AuthenticationStorage *AuthenticationStorageSession) MAXDEADLINE() (*big.Int, error) {
	return _AuthenticationStorage.Contract.MAXDEADLINE(&_AuthenticationStorage.CallOpts)
}

// MAXDEADLINE is a free data retrieval call binding the contract method 0x6f0e530f.
//
// Solidity: function MAX_DEADLINE() view returns(uint256)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) MAXDEADLINE() (*big.Int, error) {
	return _AuthenticationStorage.Contract.MAXDEADLINE(&_AuthenticationStorage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuthenticationStorage *AuthenticationStorageCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuthenticationStorage *AuthenticationStorageSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AuthenticationStorage.Contract.UPGRADEINTERFACEVERSION(&_AuthenticationStorage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AuthenticationStorage.Contract.UPGRADEINTERFACEVERSION(&_AuthenticationStorage.CallOpts)
}

// CredentialRegistry is a free data retrieval call binding the contract method 0xb948c3d4.
//
// Solidity: function credentialRegistry() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCaller) CredentialRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "credentialRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredentialRegistry is a free data retrieval call binding the contract method 0xb948c3d4.
//
// Solidity: function credentialRegistry() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageSession) CredentialRegistry() (common.Address, error) {
	return _AuthenticationStorage.Contract.CredentialRegistry(&_AuthenticationStorage.CallOpts)
}

// CredentialRegistry is a free data retrieval call binding the contract method 0xb948c3d4.
//
// Solidity: function credentialRegistry() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) CredentialRegistry() (common.Address, error) {
	return _AuthenticationStorage.Contract.CredentialRegistry(&_AuthenticationStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageSession) Owner() (common.Address, error) {
	return _AuthenticationStorage.Contract.Owner(&_AuthenticationStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) Owner() (common.Address, error) {
	return _AuthenticationStorage.Contract.Owner(&_AuthenticationStorage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuthenticationStorage *AuthenticationStorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuthenticationStorage *AuthenticationStorageSession) ProxiableUUID() ([32]byte, error) {
	return _AuthenticationStorage.Contract.ProxiableUUID(&_AuthenticationStorage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AuthenticationStorage.Contract.ProxiableUUID(&_AuthenticationStorage.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuthenticationStorage.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageSession) Verifier() (common.Address, error) {
	return _AuthenticationStorage.Contract.Verifier(&_AuthenticationStorage.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_AuthenticationStorage *AuthenticationStorageCallerSession) Verifier() (common.Address, error) {
	return _AuthenticationStorage.Contract.Verifier(&_AuthenticationStorage.CallOpts)
}

// AuthenticationStorageInit is a paid mutator transaction binding the contract method 0xbbb9d4b0.
//
// Solidity: function __AuthenticationStorage_init(address credentialRegistry_, address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) AuthenticationStorageInit(opts *bind.TransactOpts, credentialRegistry_ common.Address, verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "__AuthenticationStorage_init", credentialRegistry_, verifier_)
}

// AuthenticationStorageInit is a paid mutator transaction binding the contract method 0xbbb9d4b0.
//
// Solidity: function __AuthenticationStorage_init(address credentialRegistry_, address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageSession) AuthenticationStorageInit(credentialRegistry_ common.Address, verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.AuthenticationStorageInit(&_AuthenticationStorage.TransactOpts, credentialRegistry_, verifier_)
}

// AuthenticationStorageInit is a paid mutator transaction binding the contract method 0xbbb9d4b0.
//
// Solidity: function __AuthenticationStorage_init(address credentialRegistry_, address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) AuthenticationStorageInit(credentialRegistry_ common.Address, verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.AuthenticationStorageInit(&_AuthenticationStorage.TransactOpts, credentialRegistry_, verifier_)
}

// Register is a paid mutator transaction binding the contract method 0xc93a985c.
//
// Solidity: function register(address nft_, uint256 tokenId_, address nftOwner_, bytes32 credentialId_, uint256 deadline_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) Register(opts *bind.TransactOpts, nft_ common.Address, tokenId_ *big.Int, nftOwner_ common.Address, credentialId_ [32]byte, deadline_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "register", nft_, tokenId_, nftOwner_, credentialId_, deadline_, zkPoints_)
}

// Register is a paid mutator transaction binding the contract method 0xc93a985c.
//
// Solidity: function register(address nft_, uint256 tokenId_, address nftOwner_, bytes32 credentialId_, uint256 deadline_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_AuthenticationStorage *AuthenticationStorageSession) Register(nft_ common.Address, tokenId_ *big.Int, nftOwner_ common.Address, credentialId_ [32]byte, deadline_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.Register(&_AuthenticationStorage.TransactOpts, nft_, tokenId_, nftOwner_, credentialId_, deadline_, zkPoints_)
}

// Register is a paid mutator transaction binding the contract method 0xc93a985c.
//
// Solidity: function register(address nft_, uint256 tokenId_, address nftOwner_, bytes32 credentialId_, uint256 deadline_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) Register(nft_ common.Address, tokenId_ *big.Int, nftOwner_ common.Address, credentialId_ [32]byte, deadline_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.Register(&_AuthenticationStorage.TransactOpts, nft_, tokenId_, nftOwner_, credentialId_, deadline_, zkPoints_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthenticationStorage *AuthenticationStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.RenounceOwnership(&_AuthenticationStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.RenounceOwnership(&_AuthenticationStorage.TransactOpts)
}

// SetCredentialRegistry is a paid mutator transaction binding the contract method 0xa7013707.
//
// Solidity: function setCredentialRegistry(address credentialRegistry_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) SetCredentialRegistry(opts *bind.TransactOpts, credentialRegistry_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "setCredentialRegistry", credentialRegistry_)
}

// SetCredentialRegistry is a paid mutator transaction binding the contract method 0xa7013707.
//
// Solidity: function setCredentialRegistry(address credentialRegistry_) returns()
func (_AuthenticationStorage *AuthenticationStorageSession) SetCredentialRegistry(credentialRegistry_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.SetCredentialRegistry(&_AuthenticationStorage.TransactOpts, credentialRegistry_)
}

// SetCredentialRegistry is a paid mutator transaction binding the contract method 0xa7013707.
//
// Solidity: function setCredentialRegistry(address credentialRegistry_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) SetCredentialRegistry(credentialRegistry_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.SetCredentialRegistry(&_AuthenticationStorage.TransactOpts, credentialRegistry_)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) SetVerifier(opts *bind.TransactOpts, verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "setVerifier", verifier_)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageSession) SetVerifier(verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.SetVerifier(&_AuthenticationStorage.TransactOpts, verifier_)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address verifier_) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) SetVerifier(verifier_ common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.SetVerifier(&_AuthenticationStorage.TransactOpts, verifier_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthenticationStorage *AuthenticationStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.TransferOwnership(&_AuthenticationStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.TransferOwnership(&_AuthenticationStorage.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthenticationStorage *AuthenticationStorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthenticationStorage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthenticationStorage *AuthenticationStorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.UpgradeToAndCall(&_AuthenticationStorage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AuthenticationStorage *AuthenticationStorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AuthenticationStorage.Contract.UpgradeToAndCall(&_AuthenticationStorage.TransactOpts, newImplementation, data)
}

// AuthenticationStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AuthenticationStorage contract.
type AuthenticationStorageInitializedIterator struct {
	Event *AuthenticationStorageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuthenticationStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthenticationStorageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuthenticationStorageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuthenticationStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthenticationStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthenticationStorageInitialized represents a Initialized event raised by the AuthenticationStorage contract.
type AuthenticationStorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuthenticationStorage *AuthenticationStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuthenticationStorageInitializedIterator, error) {

	logs, sub, err := _AuthenticationStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageInitializedIterator{contract: _AuthenticationStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuthenticationStorage *AuthenticationStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuthenticationStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _AuthenticationStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthenticationStorageInitialized)
				if err := _AuthenticationStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuthenticationStorage *AuthenticationStorageFilterer) ParseInitialized(log types.Log) (*AuthenticationStorageInitialized, error) {
	event := new(AuthenticationStorageInitialized)
	if err := _AuthenticationStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthenticationStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuthenticationStorage contract.
type AuthenticationStorageOwnershipTransferredIterator struct {
	Event *AuthenticationStorageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuthenticationStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthenticationStorageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuthenticationStorageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuthenticationStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthenticationStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthenticationStorageOwnershipTransferred represents a OwnershipTransferred event raised by the AuthenticationStorage contract.
type AuthenticationStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthenticationStorage *AuthenticationStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuthenticationStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthenticationStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageOwnershipTransferredIterator{contract: _AuthenticationStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthenticationStorage *AuthenticationStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuthenticationStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuthenticationStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthenticationStorageOwnershipTransferred)
				if err := _AuthenticationStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuthenticationStorage *AuthenticationStorageFilterer) ParseOwnershipTransferred(log types.Log) (*AuthenticationStorageOwnershipTransferred, error) {
	event := new(AuthenticationStorageOwnershipTransferred)
	if err := _AuthenticationStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthenticationStorageRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the AuthenticationStorage contract.
type AuthenticationStorageRegisteredIterator struct {
	Event *AuthenticationStorageRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuthenticationStorageRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthenticationStorageRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuthenticationStorageRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuthenticationStorageRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthenticationStorageRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthenticationStorageRegistered represents a Registered event raised by the AuthenticationStorage contract.
type AuthenticationStorageRegistered struct {
	Nft          common.Address
	TokenId      *big.Int
	NftOwner     common.Address
	CredentialId [32]byte
	Deadline     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xdc72f447eb6c32bde8c6dcd2d09f5b297b4d780aaf6ab9fb99d3d083b4b0058c.
//
// Solidity: event Registered(address nft, uint256 tokenId, address nftOwner, bytes32 credentialId, uint256 deadline)
func (_AuthenticationStorage *AuthenticationStorageFilterer) FilterRegistered(opts *bind.FilterOpts) (*AuthenticationStorageRegisteredIterator, error) {

	logs, sub, err := _AuthenticationStorage.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageRegisteredIterator{contract: _AuthenticationStorage.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xdc72f447eb6c32bde8c6dcd2d09f5b297b4d780aaf6ab9fb99d3d083b4b0058c.
//
// Solidity: event Registered(address nft, uint256 tokenId, address nftOwner, bytes32 credentialId, uint256 deadline)
func (_AuthenticationStorage *AuthenticationStorageFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *AuthenticationStorageRegistered) (event.Subscription, error) {

	logs, sub, err := _AuthenticationStorage.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthenticationStorageRegistered)
				if err := _AuthenticationStorage.contract.UnpackLog(event, "Registered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistered is a log parse operation binding the contract event 0xdc72f447eb6c32bde8c6dcd2d09f5b297b4d780aaf6ab9fb99d3d083b4b0058c.
//
// Solidity: event Registered(address nft, uint256 tokenId, address nftOwner, bytes32 credentialId, uint256 deadline)
func (_AuthenticationStorage *AuthenticationStorageFilterer) ParseRegistered(log types.Log) (*AuthenticationStorageRegistered, error) {
	event := new(AuthenticationStorageRegistered)
	if err := _AuthenticationStorage.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthenticationStorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AuthenticationStorage contract.
type AuthenticationStorageUpgradedIterator struct {
	Event *AuthenticationStorageUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuthenticationStorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuthenticationStorageUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuthenticationStorageUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuthenticationStorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuthenticationStorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuthenticationStorageUpgraded represents a Upgraded event raised by the AuthenticationStorage contract.
type AuthenticationStorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuthenticationStorage *AuthenticationStorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AuthenticationStorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuthenticationStorage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AuthenticationStorageUpgradedIterator{contract: _AuthenticationStorage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuthenticationStorage *AuthenticationStorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AuthenticationStorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AuthenticationStorage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuthenticationStorageUpgraded)
				if err := _AuthenticationStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AuthenticationStorage *AuthenticationStorageFilterer) ParseUpgraded(log types.Log) (*AuthenticationStorageUpgraded, error) {
	event := new(AuthenticationStorageUpgraded)
	if err := _AuthenticationStorage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
