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

// ChatMessage is an auto generated low-level Go binding around an user-defined struct.
type ChatMessage struct {
	Message   string
	Timestamp *big.Int
}

// ChatMetaData contains all meta data concerning the Chat contract.
var ChatMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CredentialRootInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currectTime_\",\"type\":\"uint256\"}],\"name\":\"DeadlineNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZKProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC721\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"MessagePosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEADLINE_VALIDITY_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authenticationStorage_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"postMessageVerifier_\",\"type\":\"address\"}],\"name\":\"__Chat_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credentialStorage\",\"outputs\":[{\"internalType\":\"contractPoseidonSMT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"listMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structChat.Message[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"message_\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expectedMessageTime_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"}],\"name\":\"postMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postMessageVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credentialStorage_\",\"type\":\"address\"}],\"name\":\"setCredentialStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"postMessageVerifier_\",\"type\":\"address\"}],\"name\":\"setPostMessageVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b608051611c126100fd600039600081816109a6015281816109cf0152610b150152611c126000f3fe6080604052600436106100dd5760003560e01c80637524ac3d1161007f578063ad3cb1cc11610059578063ad3cb1cc1461025a578063d09ce7cc14610298578063e0d30bb3146102b8578063f2fde38b146102ce57600080fd5b80637524ac3d146101dd5780637873dd5f146101fd5780638da5cb5b1461021d57600080fd5b806351d85c10116100bb57806351d85c101461014d57806352d1902d146101855780635af8bfc6146101a8578063715018a6146101c857600080fd5b806306b7702f146100e25780631e9d3959146101185780634f1ef2861461013a575b600080fd5b3480156100ee57600080fd5b506101026100fd366004611436565b6102ee565b60405161010f91906114bb565b60405180910390f35b34801561012457600080fd5b50610138610133366004611531565b610433565b005b610138610148366004611626565b61045d565b34801561015957600080fd5b5060015461016d906001600160a01b031681565b6040516001600160a01b03909116815260200161010f565b34801561019157600080fd5b5061019a61047c565b60405190815260200161010f565b3480156101b457600080fd5b506101386101c336600461168a565b610499565b3480156101d457600080fd5b50610138610610565b3480156101e957600080fd5b506101386101f8366004611720565b610624565b34801561020957600080fd5b50610138610218366004611531565b610765565b34801561022957600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661016d565b34801561026657600080fd5b5061028b604051806040016040528060058152602001640352e302e360dc1b81525081565b60405161010f9190611759565b3480156102a457600080fd5b5060005461016d906001600160a01b031681565b3480156102c457600080fd5b5061019a610e1081565b3480156102da57600080fd5b506101386102e9366004611531565b61078f565b6001600160a01b0383166000908152600260205260408120606091906103159085856107cd565b6001600160a01b03861660009081526003602052604081209192509061033c908686610894565b90506000825167ffffffffffffffff81111561035a5761035a61154e565b6040519080825280602002602001820160405280156103a057816020015b6040805180820190915260608152600060208201528152602001906001900390816103785790505b50905060005b83518110156104265760405180604001604052808583815181106103cc576103cc61176c565b602002602001015181526020018483815181106103eb576103eb61176c565b60200260200101518152508282815181106104085761040861176c565b6020026020010181905250808061041e90611798565b9150506103a6565b50925050505b9392505050565b61043b610940565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b61046561099b565b61046e82610a40565b6104788282610a48565b5050565b6000610486610b0a565b50600080516020611bbd83398151915290565b600154604051630c3bd06d60e21b8152600481018590526001600160a01b03909116906330ef41b490602401602060405180830381865afa1580156104e2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050691906117b1565b6105235760405163a0e51db960e01b815260040160405180910390fd5b42821015610552576040516339b0da6b60e01b8152600481018390524260248201526044015b60405180910390fd5b61056d8585858561056836879003870187611823565b610b53565b61058a576040516303b2487b60e11b815260040160405180910390fd5b6001600160a01b03851660009081526002602052604090206105ac9085610c5b565b506001600160a01b03851660009081526003602052604090206105cf9042610c70565b507ffe7cdc2459269b5bc10478528a21029a11424d593a2c88471f4f5cfb582765b785856040516106019291906118bd565b60405180910390a15050505050565b610618610940565b6106226000610c7c565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff1660008115801561066a5750825b905060008267ffffffffffffffff1660011480156106875750303b155b905081158015610695575080155b156106b35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106dd57845460ff60401b1916600160401b1785555b6106e633610ced565b600080546001600160a01b038089166001600160a01b03199283161790925560018054928a1692909116919091179055831561075c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b61076d610940565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b610797610940565b6001600160a01b0381166107c157604051631e4fbdf760e01b815260006004820152602401610549565b6107ca81610c7c565b50565b606060006107e46107dd86610cfe565b8585610d08565b90506107f084826118e1565b67ffffffffffffffff8111156108085761080861154e565b60405190808252806020026020018201604052801561083b57816020015b60608152602001906001900390816108265790505b509150835b8181101561088b576108528682610d31565b8361085d87846118e1565b8151811061086d5761086d61176c565b6020026020010181905250808061088390611798565b915050610840565b50509392505050565b606060006108a46107dd86610cfe565b90506108b084826118e1565b67ffffffffffffffff8111156108c8576108c861154e565b6040519080825280602002602001820160405280156108f1578160200160208202803683370190505b509150835b8181101561088b576109088682610d3d565b8361091387846118e1565b815181106109235761092361176c565b60209081029190910101528061093881611798565b9150506108f6565b336109727f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106225760405163118cdaa760e01b8152336004820152602401610549565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610a2257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a16600080516020611bbd833981519152546001600160a01b031690565b6001600160a01b031614155b156106225760405163703e46dd60e11b815260040160405180910390fd5b6107ca610940565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610aa2575060408051601f3d908101601f19168201909252610a9f918101906118f4565b60015b610aca57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610549565b600080516020611bbd8339815191528114610afb57604051632a87526960e21b815260048101829052602401610549565b610b058383610d49565b505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106225760405163703e46dd60e11b815260040160405180910390fd5b60408051600480825260a0820190925260009182919060208201608080368337019050509050866001600160a01b031681600081518110610b9657610b9661176c565b6020026020010181815250508460001c81600181518110610bb957610bb961176c565b60200260200101818152505085604051602001610bd6919061190d565b604051602081830303815290604052805190602001206001600160f81b0360001b1660001c81600281518110610c0e57610c0e61176c565b6020026020010181815250508381600381518110610c2e57610c2e61176c565b6020908102919091010152600054610c50906001600160a01b03168285610d9f565b979650505050505050565b6000610c678383610dc4565b90505b92915050565b6000610c678383610e31565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b610cf5610e78565b6107ca81610ec1565b6000610c6a825490565b6000610d148284611929565b905083811115610d215750825b8083111561042c57509092915050565b6060610c678383610ec9565b6000610c678383610f7b565b610d5282610fa5565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115610d9757610b05828261100a565b610478611080565b6000610dbc8483600001518460200151856040015187885161109f565b949350505050565b80516020808301919091206000908152600184019091526040812054610e2957825460018101845560008481526020902001610e0083826119c4565b505081548151602080840191909120600090815260018086019092526040902091909155610c6a565b506000610c6a565b6000818152600183016020526040812054610e2957508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610c6a565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661062257604051631afcd79f60e31b815260040160405180910390fd5b610797610e78565b6060826000018281548110610ee057610ee061176c565b906000526020600020018054610ef59061193c565b80601f0160208091040260200160405190810160405280929190818152602001828054610f219061193c565b8015610f6e5780601f10610f4357610100808354040283529160200191610f6e565b820191906000526020600020905b815481529060010190602001808311610f5157829003601f168201915b5050505050905092915050565b6000826000018281548110610f9257610f9261176c565b9060005260206000200154905092915050565b806001600160a01b03163b600003610fdb57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610549565b600080516020611bbd83398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051611027919061190d565b600060405180830381855af49150503d8060008114611062576040519150601f19603f3d011682016040523d82523d6000602084013e611067565b606091505b5091509150611077858383611231565b95945050505050565b34156106225760405163b398979f60e01b815260040160405180910390fd5b6000806110ab8361128d565b6040516020016110bb9190611a84565b6040516020818303038152906040529050600080896001600160a01b0316838a8a8a6040516024016110ef93929190611b24565b60408051601f19818403018152908290529161110a9161190d565b60405180910390206001600160e01b0319166020820180516001600160e01b03838183161783525050505087604051602001611147929190611b74565b60408051601f19818403018152908290526111619161190d565b600060405180830381855afa9150503d806000811461119c576040519150601f19603f3d011682016040523d82523d6000602084013e6111a1565b606091505b50915091508161120f5760405162461bcd60e51b815260206004820152603360248201527f566572696669657248656c7065723a206661696c656420746f2063616c6c207660448201527232b934b33ca83937b7b310333ab731ba34b7b760691b6064820152608401610549565b8080602001905181019061122391906117b1565b9a9950505050505050505050565b6060826112465761124182611320565b61042c565b815115801561125d57506001600160a01b0384163b155b1561128657604051639996b31560e01b81526001600160a01b0385166004820152602401610549565b508061042c565b6060600061129a83611349565b600101905060008167ffffffffffffffff8111156112ba576112ba61154e565b6040519080825280601f01601f1916602001820160405280156112e4576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846112ee57509392505050565b8051156113305780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106113885772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef810000000083106113b4576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc1000083106113d257662386f26fc10000830492506010015b6305f5e10083106113ea576305f5e100830492506008015b61271083106113fe57612710830492506004015b60648310611410576064830492506002015b600a8310610c6a5760010192915050565b6001600160a01b03811681146107ca57600080fd5b60008060006060848603121561144b57600080fd5b833561145681611421565b95602085013595506040909401359392505050565b60005b8381101561148657818101518382015260200161146e565b50506000910152565b600081518084526114a781602086016020860161146b565b601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561152357888303603f19018552815180518785526115068886018261148f565b9189015194890194909452948701949250908601906001016114e2565b509098975050505050505050565b60006020828403121561154357600080fd5b813561042c81611421565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156115875761158761154e565b60405290565b6040805190810167ffffffffffffffff811182821017156115875761158761154e565b600067ffffffffffffffff808411156115cb576115cb61154e565b604051601f8501601f19908116603f011681019082821181831017156115f3576115f361154e565b8160405280935085815286868601111561160c57600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561163957600080fd5b823561164481611421565b9150602083013567ffffffffffffffff81111561166057600080fd5b8301601f8101851361167157600080fd5b611680858235602084016115b0565b9150509250929050565b60008060008060008587036101808112156116a457600080fd5b86356116af81611421565b9550602087013567ffffffffffffffff8111156116cb57600080fd5b8701601f810189136116dc57600080fd5b6116eb898235602084016115b0565b9550506040870135935060608701359250610100607f198201121561170f57600080fd5b506080860190509295509295909350565b6000806040838503121561173357600080fd5b823561173e81611421565b9150602083013561174e81611421565b809150509250929050565b602081526000610c67602083018461148f565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016117aa576117aa611782565b5060010190565b6000602082840312156117c357600080fd5b8151801515811461042c57600080fd5b600082601f8301126117e457600080fd5b6117ec61158d565b8060408401858111156117fe57600080fd5b845b81811015611818578035845260209384019301611800565b509095945050505050565b6000610100828403121561183657600080fd5b61183e611564565b61184884846117d3565b8152604084605f85011261185b57600080fd5b61186361158d565b8060c086018781111561187557600080fd5b8387015b8181101561189a5761188b89826117d3565b84526020909301928401611879565b508160208601526118ab88826117d3565b84860152505050508091505092915050565b6001600160a01b0383168152604060208201819052600090610dbc9083018461148f565b81810381811115610c6a57610c6a611782565b60006020828403121561190657600080fd5b5051919050565b6000825161191f81846020870161146b565b9190910192915050565b80820180821115610c6a57610c6a611782565b600181811c9082168061195057607f821691505b60208210810361197057634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610b0557600081815260208120601f850160051c8101602086101561199d5750805b601f850160051c820191505b818110156119bc578281556001016119a9565b505050505050565b815167ffffffffffffffff8111156119de576119de61154e565b6119f2816119ec845461193c565b84611976565b602080601f831160018114611a275760008415611a0f5750858301515b600019600386901b1c1916600185901b1785556119bc565b600085815260208120601f198616915b82811015611a5657888601518255948401946001909101908401611a37565b5085821015611a745787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b7f76657269667950726f6f662875696e743235365b325d2c75696e743235365b3281527f5d5b325d2c75696e743235365b325d2c75696e743235365b0000000000000000602082015260008251611ae281603885016020870161146b565b615d2960f01b6038939091019283015250603a01919050565b8060005b6002811015611b1e578151845260209384019390910190600101611aff565b50505050565b6101008101611b338286611afb565b60408083018560005b6002811015611b6357611b50838351611afb565b9183019160209190910190600101611b3c565b50505050610dbc60c0830184611afb565b600083516020611b87828583890161146b565b84519184019181860160005b82811015611baf57815185529383019390830190600101611b93565b509297965050505050505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca26469706673582212209f20d16ba3e09539803ee8ac866314bb0c9aa11e7ee02bb7d754202c7591571f64736f6c63430008140033",
}

// ChatABI is the input ABI used to generate the binding from.
// Deprecated: Use ChatMetaData.ABI instead.
var ChatABI = ChatMetaData.ABI

// ChatBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChatMetaData.Bin instead.
var ChatBin = ChatMetaData.Bin

// DeployChat deploys a new Ethereum contract, binding an instance of Chat to it.
func DeployChat(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Chat, error) {
	parsed, err := ChatMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChatBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Chat{ChatCaller: ChatCaller{contract: contract}, ChatTransactor: ChatTransactor{contract: contract}, ChatFilterer: ChatFilterer{contract: contract}}, nil
}

// Chat is an auto generated Go binding around an Ethereum contract.
type Chat struct {
	ChatCaller     // Read-only binding to the contract
	ChatTransactor // Write-only binding to the contract
	ChatFilterer   // Log filterer for contract events
}

// ChatCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChatSession struct {
	Contract     *Chat             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChatCallerSession struct {
	Contract *ChatCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChatTransactorSession struct {
	Contract     *ChatTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChatRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChatRaw struct {
	Contract *Chat // Generic contract binding to access the raw methods on
}

// ChatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChatCallerRaw struct {
	Contract *ChatCaller // Generic read-only contract binding to access the raw methods on
}

// ChatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChatTransactorRaw struct {
	Contract *ChatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChat creates a new instance of Chat, bound to a specific deployed contract.
func NewChat(address common.Address, backend bind.ContractBackend) (*Chat, error) {
	contract, err := bindChat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chat{ChatCaller: ChatCaller{contract: contract}, ChatTransactor: ChatTransactor{contract: contract}, ChatFilterer: ChatFilterer{contract: contract}}, nil
}

// NewChatCaller creates a new read-only instance of Chat, bound to a specific deployed contract.
func NewChatCaller(address common.Address, caller bind.ContractCaller) (*ChatCaller, error) {
	contract, err := bindChat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChatCaller{contract: contract}, nil
}

// NewChatTransactor creates a new write-only instance of Chat, bound to a specific deployed contract.
func NewChatTransactor(address common.Address, transactor bind.ContractTransactor) (*ChatTransactor, error) {
	contract, err := bindChat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChatTransactor{contract: contract}, nil
}

// NewChatFilterer creates a new log filterer instance of Chat, bound to a specific deployed contract.
func NewChatFilterer(address common.Address, filterer bind.ContractFilterer) (*ChatFilterer, error) {
	contract, err := bindChat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChatFilterer{contract: contract}, nil
}

// bindChat binds a generic wrapper to an already deployed contract.
func bindChat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chat *ChatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chat.Contract.ChatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chat *ChatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chat.Contract.ChatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chat *ChatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chat.Contract.ChatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Chat *ChatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Chat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Chat *ChatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Chat *ChatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Chat.Contract.contract.Transact(opts, method, params...)
}

// DEADLINEVALIDITYWINDOW is a free data retrieval call binding the contract method 0xe0d30bb3.
//
// Solidity: function DEADLINE_VALIDITY_WINDOW() view returns(uint256)
func (_Chat *ChatCaller) DEADLINEVALIDITYWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "DEADLINE_VALIDITY_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEADLINEVALIDITYWINDOW is a free data retrieval call binding the contract method 0xe0d30bb3.
//
// Solidity: function DEADLINE_VALIDITY_WINDOW() view returns(uint256)
func (_Chat *ChatSession) DEADLINEVALIDITYWINDOW() (*big.Int, error) {
	return _Chat.Contract.DEADLINEVALIDITYWINDOW(&_Chat.CallOpts)
}

// DEADLINEVALIDITYWINDOW is a free data retrieval call binding the contract method 0xe0d30bb3.
//
// Solidity: function DEADLINE_VALIDITY_WINDOW() view returns(uint256)
func (_Chat *ChatCallerSession) DEADLINEVALIDITYWINDOW() (*big.Int, error) {
	return _Chat.Contract.DEADLINEVALIDITYWINDOW(&_Chat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Chat *ChatCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Chat *ChatSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Chat.Contract.UPGRADEINTERFACEVERSION(&_Chat.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Chat *ChatCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Chat.Contract.UPGRADEINTERFACEVERSION(&_Chat.CallOpts)
}

// CredentialStorage is a free data retrieval call binding the contract method 0x51d85c10.
//
// Solidity: function credentialStorage() view returns(address)
func (_Chat *ChatCaller) CredentialStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "credentialStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredentialStorage is a free data retrieval call binding the contract method 0x51d85c10.
//
// Solidity: function credentialStorage() view returns(address)
func (_Chat *ChatSession) CredentialStorage() (common.Address, error) {
	return _Chat.Contract.CredentialStorage(&_Chat.CallOpts)
}

// CredentialStorage is a free data retrieval call binding the contract method 0x51d85c10.
//
// Solidity: function credentialStorage() view returns(address)
func (_Chat *ChatCallerSession) CredentialStorage() (common.Address, error) {
	return _Chat.Contract.CredentialStorage(&_Chat.CallOpts)
}

// ListMessages is a free data retrieval call binding the contract method 0x06b7702f.
//
// Solidity: function listMessages(address nft_, uint256 offset_, uint256 limit_) view returns((string,uint256)[])
func (_Chat *ChatCaller) ListMessages(opts *bind.CallOpts, nft_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]ChatMessage, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "listMessages", nft_, offset_, limit_)

	if err != nil {
		return *new([]ChatMessage), err
	}

	out0 := *abi.ConvertType(out[0], new([]ChatMessage)).(*[]ChatMessage)

	return out0, err

}

// ListMessages is a free data retrieval call binding the contract method 0x06b7702f.
//
// Solidity: function listMessages(address nft_, uint256 offset_, uint256 limit_) view returns((string,uint256)[])
func (_Chat *ChatSession) ListMessages(nft_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]ChatMessage, error) {
	return _Chat.Contract.ListMessages(&_Chat.CallOpts, nft_, offset_, limit_)
}

// ListMessages is a free data retrieval call binding the contract method 0x06b7702f.
//
// Solidity: function listMessages(address nft_, uint256 offset_, uint256 limit_) view returns((string,uint256)[])
func (_Chat *ChatCallerSession) ListMessages(nft_ common.Address, offset_ *big.Int, limit_ *big.Int) ([]ChatMessage, error) {
	return _Chat.Contract.ListMessages(&_Chat.CallOpts, nft_, offset_, limit_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chat *ChatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chat *ChatSession) Owner() (common.Address, error) {
	return _Chat.Contract.Owner(&_Chat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Chat *ChatCallerSession) Owner() (common.Address, error) {
	return _Chat.Contract.Owner(&_Chat.CallOpts)
}

// PostMessageVerifier is a free data retrieval call binding the contract method 0xd09ce7cc.
//
// Solidity: function postMessageVerifier() view returns(address)
func (_Chat *ChatCaller) PostMessageVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "postMessageVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PostMessageVerifier is a free data retrieval call binding the contract method 0xd09ce7cc.
//
// Solidity: function postMessageVerifier() view returns(address)
func (_Chat *ChatSession) PostMessageVerifier() (common.Address, error) {
	return _Chat.Contract.PostMessageVerifier(&_Chat.CallOpts)
}

// PostMessageVerifier is a free data retrieval call binding the contract method 0xd09ce7cc.
//
// Solidity: function postMessageVerifier() view returns(address)
func (_Chat *ChatCallerSession) PostMessageVerifier() (common.Address, error) {
	return _Chat.Contract.PostMessageVerifier(&_Chat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Chat *ChatCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Chat.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Chat *ChatSession) ProxiableUUID() ([32]byte, error) {
	return _Chat.Contract.ProxiableUUID(&_Chat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Chat *ChatCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Chat.Contract.ProxiableUUID(&_Chat.CallOpts)
}

// ChatInit is a paid mutator transaction binding the contract method 0x7524ac3d.
//
// Solidity: function __Chat_init(address authenticationStorage_, address postMessageVerifier_) returns()
func (_Chat *ChatTransactor) ChatInit(opts *bind.TransactOpts, authenticationStorage_ common.Address, postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "__Chat_init", authenticationStorage_, postMessageVerifier_)
}

// ChatInit is a paid mutator transaction binding the contract method 0x7524ac3d.
//
// Solidity: function __Chat_init(address authenticationStorage_, address postMessageVerifier_) returns()
func (_Chat *ChatSession) ChatInit(authenticationStorage_ common.Address, postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.ChatInit(&_Chat.TransactOpts, authenticationStorage_, postMessageVerifier_)
}

// ChatInit is a paid mutator transaction binding the contract method 0x7524ac3d.
//
// Solidity: function __Chat_init(address authenticationStorage_, address postMessageVerifier_) returns()
func (_Chat *ChatTransactorSession) ChatInit(authenticationStorage_ common.Address, postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.ChatInit(&_Chat.TransactOpts, authenticationStorage_, postMessageVerifier_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x5af8bfc6.
//
// Solidity: function postMessage(address nft_, string message_, bytes32 root_, uint256 expectedMessageTime_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Chat *ChatTransactor) PostMessage(opts *bind.TransactOpts, nft_ common.Address, message_ string, root_ [32]byte, expectedMessageTime_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "postMessage", nft_, message_, root_, expectedMessageTime_, zkPoints_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x5af8bfc6.
//
// Solidity: function postMessage(address nft_, string message_, bytes32 root_, uint256 expectedMessageTime_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Chat *ChatSession) PostMessage(nft_ common.Address, message_ string, root_ [32]byte, expectedMessageTime_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Chat.Contract.PostMessage(&_Chat.TransactOpts, nft_, message_, root_, expectedMessageTime_, zkPoints_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x5af8bfc6.
//
// Solidity: function postMessage(address nft_, string message_, bytes32 root_, uint256 expectedMessageTime_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_) returns()
func (_Chat *ChatTransactorSession) PostMessage(nft_ common.Address, message_ string, root_ [32]byte, expectedMessageTime_ *big.Int, zkPoints_ VerifierHelperProofPoints) (*types.Transaction, error) {
	return _Chat.Contract.PostMessage(&_Chat.TransactOpts, nft_, message_, root_, expectedMessageTime_, zkPoints_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chat *ChatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chat *ChatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chat.Contract.RenounceOwnership(&_Chat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Chat *ChatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Chat.Contract.RenounceOwnership(&_Chat.TransactOpts)
}

// SetCredentialStorage is a paid mutator transaction binding the contract method 0x1e9d3959.
//
// Solidity: function setCredentialStorage(address credentialStorage_) returns()
func (_Chat *ChatTransactor) SetCredentialStorage(opts *bind.TransactOpts, credentialStorage_ common.Address) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "setCredentialStorage", credentialStorage_)
}

// SetCredentialStorage is a paid mutator transaction binding the contract method 0x1e9d3959.
//
// Solidity: function setCredentialStorage(address credentialStorage_) returns()
func (_Chat *ChatSession) SetCredentialStorage(credentialStorage_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.SetCredentialStorage(&_Chat.TransactOpts, credentialStorage_)
}

// SetCredentialStorage is a paid mutator transaction binding the contract method 0x1e9d3959.
//
// Solidity: function setCredentialStorage(address credentialStorage_) returns()
func (_Chat *ChatTransactorSession) SetCredentialStorage(credentialStorage_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.SetCredentialStorage(&_Chat.TransactOpts, credentialStorage_)
}

// SetPostMessageVerifier is a paid mutator transaction binding the contract method 0x7873dd5f.
//
// Solidity: function setPostMessageVerifier(address postMessageVerifier_) returns()
func (_Chat *ChatTransactor) SetPostMessageVerifier(opts *bind.TransactOpts, postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "setPostMessageVerifier", postMessageVerifier_)
}

// SetPostMessageVerifier is a paid mutator transaction binding the contract method 0x7873dd5f.
//
// Solidity: function setPostMessageVerifier(address postMessageVerifier_) returns()
func (_Chat *ChatSession) SetPostMessageVerifier(postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.SetPostMessageVerifier(&_Chat.TransactOpts, postMessageVerifier_)
}

// SetPostMessageVerifier is a paid mutator transaction binding the contract method 0x7873dd5f.
//
// Solidity: function setPostMessageVerifier(address postMessageVerifier_) returns()
func (_Chat *ChatTransactorSession) SetPostMessageVerifier(postMessageVerifier_ common.Address) (*types.Transaction, error) {
	return _Chat.Contract.SetPostMessageVerifier(&_Chat.TransactOpts, postMessageVerifier_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chat *ChatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chat *ChatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chat.Contract.TransferOwnership(&_Chat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Chat *ChatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Chat.Contract.TransferOwnership(&_Chat.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Chat *ChatTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Chat.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Chat *ChatSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Chat.Contract.UpgradeToAndCall(&_Chat.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Chat *ChatTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Chat.Contract.UpgradeToAndCall(&_Chat.TransactOpts, newImplementation, data)
}

// ChatInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Chat contract.
type ChatInitializedIterator struct {
	Event *ChatInitialized // Event containing the contract specifics and raw log

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
func (it *ChatInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatInitialized)
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
		it.Event = new(ChatInitialized)
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
func (it *ChatInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatInitialized represents a Initialized event raised by the Chat contract.
type ChatInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Chat *ChatFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChatInitializedIterator, error) {

	logs, sub, err := _Chat.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChatInitializedIterator{contract: _Chat.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Chat *ChatFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChatInitialized) (event.Subscription, error) {

	logs, sub, err := _Chat.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatInitialized)
				if err := _Chat.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Chat *ChatFilterer) ParseInitialized(log types.Log) (*ChatInitialized, error) {
	event := new(ChatInitialized)
	if err := _Chat.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChatMessagePostedIterator is returned from FilterMessagePosted and is used to iterate over the raw logs and unpacked data for MessagePosted events raised by the Chat contract.
type ChatMessagePostedIterator struct {
	Event *ChatMessagePosted // Event containing the contract specifics and raw log

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
func (it *ChatMessagePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatMessagePosted)
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
		it.Event = new(ChatMessagePosted)
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
func (it *ChatMessagePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatMessagePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatMessagePosted represents a MessagePosted event raised by the Chat contract.
type ChatMessagePosted struct {
	Nft     common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessagePosted is a free log retrieval operation binding the contract event 0xfe7cdc2459269b5bc10478528a21029a11424d593a2c88471f4f5cfb582765b7.
//
// Solidity: event MessagePosted(address nft, string message)
func (_Chat *ChatFilterer) FilterMessagePosted(opts *bind.FilterOpts) (*ChatMessagePostedIterator, error) {

	logs, sub, err := _Chat.contract.FilterLogs(opts, "MessagePosted")
	if err != nil {
		return nil, err
	}
	return &ChatMessagePostedIterator{contract: _Chat.contract, event: "MessagePosted", logs: logs, sub: sub}, nil
}

// WatchMessagePosted is a free log subscription operation binding the contract event 0xfe7cdc2459269b5bc10478528a21029a11424d593a2c88471f4f5cfb582765b7.
//
// Solidity: event MessagePosted(address nft, string message)
func (_Chat *ChatFilterer) WatchMessagePosted(opts *bind.WatchOpts, sink chan<- *ChatMessagePosted) (event.Subscription, error) {

	logs, sub, err := _Chat.contract.WatchLogs(opts, "MessagePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatMessagePosted)
				if err := _Chat.contract.UnpackLog(event, "MessagePosted", log); err != nil {
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

// ParseMessagePosted is a log parse operation binding the contract event 0xfe7cdc2459269b5bc10478528a21029a11424d593a2c88471f4f5cfb582765b7.
//
// Solidity: event MessagePosted(address nft, string message)
func (_Chat *ChatFilterer) ParseMessagePosted(log types.Log) (*ChatMessagePosted, error) {
	event := new(ChatMessagePosted)
	if err := _Chat.contract.UnpackLog(event, "MessagePosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChatOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Chat contract.
type ChatOwnershipTransferredIterator struct {
	Event *ChatOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChatOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatOwnershipTransferred)
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
		it.Event = new(ChatOwnershipTransferred)
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
func (it *ChatOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatOwnershipTransferred represents a OwnershipTransferred event raised by the Chat contract.
type ChatOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chat *ChatFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChatOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chat.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChatOwnershipTransferredIterator{contract: _Chat.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Chat *ChatFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChatOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Chat.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatOwnershipTransferred)
				if err := _Chat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Chat *ChatFilterer) ParseOwnershipTransferred(log types.Log) (*ChatOwnershipTransferred, error) {
	event := new(ChatOwnershipTransferred)
	if err := _Chat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChatUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Chat contract.
type ChatUpgradedIterator struct {
	Event *ChatUpgraded // Event containing the contract specifics and raw log

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
func (it *ChatUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChatUpgraded)
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
		it.Event = new(ChatUpgraded)
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
func (it *ChatUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChatUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChatUpgraded represents a Upgraded event raised by the Chat contract.
type ChatUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Chat *ChatFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChatUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Chat.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChatUpgradedIterator{contract: _Chat.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Chat *ChatFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChatUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Chat.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChatUpgraded)
				if err := _Chat.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Chat *ChatFilterer) ParseUpgraded(log types.Log) (*ChatUpgraded, error) {
	event := new(ChatUpgraded)
	if err := _Chat.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
