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

// ERC721MockMetaData contains all meta data concerning the ERC721Mock contract.
var ERC721MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"getTokensByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokens_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001adc38038062001adc8339810160408190526200003491620001b1565b3382826000620000458382620002aa565b506001620000548282620002aa565b5050506001600160a01b0381166200008657604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b62000091816200009a565b50505062000376565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200011457600080fd5b81516001600160401b0380821115620001315762000131620000ec565b604051601f8301601f19908116603f011681019082821181831017156200015c576200015c620000ec565b816040528381526020925086838588010111156200017957600080fd5b600091505b838210156200019d57858201830151818301840152908201906200017e565b600093810190920192909252949350505050565b60008060408385031215620001c557600080fd5b82516001600160401b0380821115620001dd57600080fd5b620001eb8683870162000102565b935060208501519150808211156200020257600080fd5b50620002118582860162000102565b9150509250929050565b600181811c908216806200023057607f821691505b6020821081036200025157634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a557600081815260208120601f850160051c81016020861015620002805750805b601f850160051c820191505b81811015620002a1578281556001016200028c565b5050505b505050565b81516001600160401b03811115620002c657620002c6620000ec565b620002de81620002d784546200021b565b8462000257565b602080601f831160018114620003165760008415620002fd5750858301515b600019600386901b1c1916600185901b178555620002a1565b600085815260208120601f198616915b82811015620003475788860151825594840194600190910190840162000326565b5085821015620003665787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61175680620003866000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c80634f6ccce7116100b857806395d89b411161007c57806395d89b41146102a7578063a22cb465146102af578063b88d4fde146102c2578063c87b56dd146102d5578063e985e9c5146102e8578063f2fde38b146102fb57600080fd5b80634f6ccce7146102555780636352211e1461026857806370a082311461027b578063715018a61461028e5780638da5cb5b1461029657600080fd5b806323b872dd1161010a57806323b872dd146101d65780632f745c59146101e957806340398d67146101fc57806340c10f191461021c57806342842e0e1461022f57806342966c681461024257600080fd5b806301ffc9a71461014757806306fdde031461016f578063081812fc14610184578063095ea7b3146101af57806318160ddd146101c4575b600080fd5b61015a610155366004611314565b61030e565b60405190151581526020015b60405180910390f35b61017761031f565b6040516101669190611381565b610197610192366004611394565b6103b1565b6040516001600160a01b039091168152602001610166565b6101c26101bd3660046113c9565b6103da565b005b6008545b604051908152602001610166565b6101c26101e43660046113f3565b6103e9565b6101c86101f73660046113c9565b610479565b61020f61020a36600461142f565b6104de565b604051610166919061144a565b6101c261022a3660046113c9565b61057b565b6101c261023d3660046113f3565b610585565b6101c2610250366004611394565b6105a5565b6101c8610263366004611394565b6105b1565b610197610276366004611394565b61060a565b6101c861028936600461142f565b610615565b6101c261065d565b600b546001600160a01b0316610197565b610177610671565b6101c26102bd36600461148e565b610680565b6101c26102d03660046114e0565b61068b565b6101776102e3366004611394565b6106a2565b61015a6102f63660046115bc565b6106ad565b6101c261030936600461142f565b6106db565b600061031982610716565b92915050565b60606000805461032e906115ef565b80601f016020809104026020016040519081016040528092919081815260200182805461035a906115ef565b80156103a75780601f1061037c576101008083540402835291602001916103a7565b820191906000526020600020905b81548152906001019060200180831161038a57829003601f168201915b5050505050905090565b60006103bc8261073b565b506000828152600460205260409020546001600160a01b0316610319565b6103e5828233610774565b5050565b6001600160a01b03821661041857604051633250574960e11b8152600060048201526024015b60405180910390fd5b6000610425838333610781565b9050836001600160a01b0316816001600160a01b031614610473576040516364283d7b60e01b81526001600160a01b038086166004830152602482018490528216604482015260640161040f565b50505050565b600061048483610615565b82106104b55760405163295f44f760e21b81526001600160a01b03841660048201526024810183905260440161040f565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b606060006104eb83610615565b90508067ffffffffffffffff811115610506576105066114ca565b60405190808252806020026020018201604052801561052f578160200160208202803683370190505b50915060005b81811015610574576105478482610479565b83828151811061055957610559611629565b602090810291909101015261056d81611655565b9050610535565b5050919050565b6103e58282610796565b6105a08383836040518060200160405280600081525061068b565b505050565b6105ae816107fb565b50565b60006105bc60085490565b82106105e55760405163295f44f760e21b8152600060048201526024810183905260440161040f565b600882815481106105f8576105f8611629565b90600052602060002001549050919050565b60006103198261073b565b60006001600160a01b038216610641576040516322718ad960e21b81526000600482015260240161040f565b506001600160a01b031660009081526003602052604090205490565b610665610836565b61066f6000610863565b565b60606001805461032e906115ef565b6103e53383836108b5565b6106968484846103e9565b61047384848484610954565b606061031982610a7d565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6106e3610836565b6001600160a01b03811661070d57604051631e4fbdf760e01b81526000600482015260240161040f565b6105ae81610863565b60006001600160e01b03198216632483248360e11b1480610319575061031982610b86565b6000818152600260205260408120546001600160a01b03168061031957604051637e27328960e01b81526004810184905260240161040f565b6105a08383836001610bab565b600061078e848484610cb1565b949350505050565b6001600160a01b0382166107c057604051633250574960e11b81526000600482015260240161040f565b60006107ce83836000610781565b90506001600160a01b038116156105a0576040516339e3563760e11b81526000600482015260240161040f565b600061080a6000836000610781565b90506001600160a01b0381166103e557604051637e27328960e01b81526004810183905260240161040f565b600b546001600160a01b0316331461066f5760405163118cdaa760e01b815233600482015260240161040f565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b0382166108e757604051630b61174360e31b81526001600160a01b038316600482015260240161040f565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b1561047357604051630a85bd0160e11b81526001600160a01b0384169063150b7a029061099690339088908790879060040161166e565b6020604051808303816000875af19250505080156109d1575060408051601f3d908101601f191682019092526109ce918101906116ab565b60015b610a3a573d8080156109ff576040519150601f19603f3d011682016040523d82523d6000602084013e610a04565b606091505b508051600003610a3257604051633250574960e11b81526001600160a01b038516600482015260240161040f565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b14610a7657604051633250574960e11b81526001600160a01b038516600482015260240161040f565b5050505050565b6060610a888261073b565b506000828152600a602052604081208054610aa2906115ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610ace906115ef565b8015610b1b5780601f10610af057610100808354040283529160200191610b1b565b820191906000526020600020905b815481529060010190602001808311610afe57829003601f168201915b505050505090506000610b3960408051602081019091526000815290565b90508051600003610b4b575092915050565b815115610b7d578082604051602001610b659291906116c8565b60405160208183030381529060405292505050919050565b61078e84610d7e565b60006001600160e01b0319821663780e9d6360e01b1480610319575061031982610df3565b8080610bbf57506001600160a01b03821615155b15610c81576000610bcf8461073b565b90506001600160a01b03831615801590610bfb5750826001600160a01b0316816001600160a01b031614155b8015610c0e5750610c0c81846106ad565b155b15610c375760405163a9fbf51f60e01b81526001600160a01b038416600482015260240161040f565b8115610c7f5783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5050600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b600080610cbf858585610e43565b90506001600160a01b038116610d1c57610d1784600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610d3f565b846001600160a01b0316816001600160a01b031614610d3f57610d3f8185610f3c565b6001600160a01b038516610d5b57610d5684610fcd565b61078e565b846001600160a01b0316816001600160a01b03161461078e5761078e858561107c565b6060610d898261073b565b506000610da160408051602081019091526000815290565b90506000815111610dc15760405180602001604052806000815250610dec565b80610dcb846110cc565b604051602001610ddc9291906116c8565b6040516020818303038152906040525b9392505050565b60006001600160e01b031982166380ac58cd60e01b1480610e2457506001600160e01b03198216635b5e139f60e01b145b8061031957506301ffc9a760e01b6001600160e01b0319831614610319565b6000828152600260205260408120546001600160a01b0390811690831615610e7057610e7081848661115f565b6001600160a01b03811615610eae57610e8d600085600080610bab565b6001600160a01b038116600090815260036020526040902080546000190190555b6001600160a01b03851615610edd576001600160a01b0385166000908152600360205260409020805460010190555b60008481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b6000610f4783610615565b600083815260076020526040902054909150808214610f9a576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b600854600090610fdf906001906116f7565b6000838152600960205260408120546008805493945090928490811061100757611007611629565b90600052602060002001549050806008838154811061102857611028611629565b60009182526020808320909101929092558281526009909152604080822084905585825281205560088054806110605761106061170a565b6001900381819060005260206000200160009055905550505050565b6000600161108984610615565b61109391906116f7565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b606060006110d9836111c3565b600101905060008167ffffffffffffffff8111156110f9576110f96114ca565b6040519080825280601f01601f191660200182016040528015611123576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461112d57509392505050565b61116a83838361129b565b6105a0576001600160a01b03831661119857604051637e27328960e01b81526004810182905260240161040f565b60405163177e802f60e01b81526001600160a01b03831660048201526024810182905260440161040f565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106112025772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef8100000000831061122e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061124c57662386f26fc10000830492506010015b6305f5e1008310611264576305f5e100830492506008015b612710831061127857612710830492506004015b6064831061128a576064830492506002015b600a83106103195760010192915050565b60006001600160a01b0383161580159061078e5750826001600160a01b0316846001600160a01b031614806112d557506112d584846106ad565b8061078e5750506000908152600460205260409020546001600160a01b03908116911614919050565b6001600160e01b0319811681146105ae57600080fd5b60006020828403121561132657600080fd5b8135610dec816112fe565b60005b8381101561134c578181015183820152602001611334565b50506000910152565b6000815180845261136d816020860160208601611331565b601f01601f19169290920160200192915050565b602081526000610dec6020830184611355565b6000602082840312156113a657600080fd5b5035919050565b80356001600160a01b03811681146113c457600080fd5b919050565b600080604083850312156113dc57600080fd5b6113e5836113ad565b946020939093013593505050565b60008060006060848603121561140857600080fd5b611411846113ad565b925061141f602085016113ad565b9150604084013590509250925092565b60006020828403121561144157600080fd5b610dec826113ad565b6020808252825182820181905260009190848201906040850190845b8181101561148257835183529284019291840191600101611466565b50909695505050505050565b600080604083850312156114a157600080fd5b6114aa836113ad565b9150602083013580151581146114bf57600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156114f657600080fd5b6114ff856113ad565b935061150d602086016113ad565b925060408501359150606085013567ffffffffffffffff8082111561153157600080fd5b818701915087601f83011261154557600080fd5b813581811115611557576115576114ca565b604051601f8201601f19908116603f0116810190838211818310171561157f5761157f6114ca565b816040528281528a602084870101111561159857600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b600080604083850312156115cf57600080fd5b6115d8836113ad565b91506115e6602084016113ad565b90509250929050565b600181811c9082168061160357607f821691505b60208210810361162357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016116675761166761163f565b5060010190565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906116a190830184611355565b9695505050505050565b6000602082840312156116bd57600080fd5b8151610dec816112fe565b600083516116da818460208801611331565b8351908301906116ee818360208801611331565b01949350505050565b818103818111156103195761031961163f565b634e487b7160e01b600052603160045260246000fdfea26469706673582212205d637c50630fcfe1a78f4481d363b7cb78960a952b80737a51f5665f34a6f75164736f6c63430008140033",
}

// ERC721MockABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MockMetaData.ABI instead.
var ERC721MockABI = ERC721MockMetaData.ABI

// ERC721MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC721MockMetaData.Bin instead.
var ERC721MockBin = ERC721MockMetaData.Bin

// DeployERC721Mock deploys a new Ethereum contract, binding an instance of ERC721Mock to it.
func DeployERC721Mock(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC721Mock, error) {
	parsed, err := ERC721MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC721MockBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Mock{ERC721MockCaller: ERC721MockCaller{contract: contract}, ERC721MockTransactor: ERC721MockTransactor{contract: contract}, ERC721MockFilterer: ERC721MockFilterer{contract: contract}}, nil
}

// ERC721Mock is an auto generated Go binding around an Ethereum contract.
type ERC721Mock struct {
	ERC721MockCaller     // Read-only binding to the contract
	ERC721MockTransactor // Write-only binding to the contract
	ERC721MockFilterer   // Log filterer for contract events
}

// ERC721MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MockSession struct {
	Contract     *ERC721Mock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MockCallerSession struct {
	Contract *ERC721MockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC721MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MockTransactorSession struct {
	Contract     *ERC721MockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC721MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MockRaw struct {
	Contract *ERC721Mock // Generic contract binding to access the raw methods on
}

// ERC721MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MockCallerRaw struct {
	Contract *ERC721MockCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MockTransactorRaw struct {
	Contract *ERC721MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Mock creates a new instance of ERC721Mock, bound to a specific deployed contract.
func NewERC721Mock(address common.Address, backend bind.ContractBackend) (*ERC721Mock, error) {
	contract, err := bindERC721Mock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Mock{ERC721MockCaller: ERC721MockCaller{contract: contract}, ERC721MockTransactor: ERC721MockTransactor{contract: contract}, ERC721MockFilterer: ERC721MockFilterer{contract: contract}}, nil
}

// NewERC721MockCaller creates a new read-only instance of ERC721Mock, bound to a specific deployed contract.
func NewERC721MockCaller(address common.Address, caller bind.ContractCaller) (*ERC721MockCaller, error) {
	contract, err := bindERC721Mock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MockCaller{contract: contract}, nil
}

// NewERC721MockTransactor creates a new write-only instance of ERC721Mock, bound to a specific deployed contract.
func NewERC721MockTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MockTransactor, error) {
	contract, err := bindERC721Mock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MockTransactor{contract: contract}, nil
}

// NewERC721MockFilterer creates a new log filterer instance of ERC721Mock, bound to a specific deployed contract.
func NewERC721MockFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MockFilterer, error) {
	contract, err := bindERC721Mock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MockFilterer{contract: contract}, nil
}

// bindERC721Mock binds a generic wrapper to an already deployed contract.
func bindERC721Mock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721MockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Mock *ERC721MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Mock.Contract.ERC721MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Mock *ERC721MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mock.Contract.ERC721MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Mock *ERC721MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Mock.Contract.ERC721MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Mock *ERC721MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Mock *ERC721MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Mock *ERC721MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Mock.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mock *ERC721MockCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mock *ERC721MockSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Mock.Contract.BalanceOf(&_ERC721Mock.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC721Mock *ERC721MockCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Mock.Contract.BalanceOf(&_ERC721Mock.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mock.Contract.GetApproved(&_ERC721Mock.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mock.Contract.GetApproved(&_ERC721Mock.CallOpts, tokenId)
}

// GetTokensByOwner is a free data retrieval call binding the contract method 0x40398d67.
//
// Solidity: function getTokensByOwner(address owner_) view returns(uint256[] tokens_)
func (_ERC721Mock *ERC721MockCaller) GetTokensByOwner(opts *bind.CallOpts, owner_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "getTokensByOwner", owner_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokensByOwner is a free data retrieval call binding the contract method 0x40398d67.
//
// Solidity: function getTokensByOwner(address owner_) view returns(uint256[] tokens_)
func (_ERC721Mock *ERC721MockSession) GetTokensByOwner(owner_ common.Address) ([]*big.Int, error) {
	return _ERC721Mock.Contract.GetTokensByOwner(&_ERC721Mock.CallOpts, owner_)
}

// GetTokensByOwner is a free data retrieval call binding the contract method 0x40398d67.
//
// Solidity: function getTokensByOwner(address owner_) view returns(uint256[] tokens_)
func (_ERC721Mock *ERC721MockCallerSession) GetTokensByOwner(owner_ common.Address) ([]*big.Int, error) {
	return _ERC721Mock.Contract.GetTokensByOwner(&_ERC721Mock.CallOpts, owner_)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mock *ERC721MockCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mock *ERC721MockSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Mock.Contract.IsApprovedForAll(&_ERC721Mock.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ERC721Mock *ERC721MockCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Mock.Contract.IsApprovedForAll(&_ERC721Mock.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Mock *ERC721MockCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Mock *ERC721MockSession) Name() (string, error) {
	return _ERC721Mock.Contract.Name(&_ERC721Mock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC721Mock *ERC721MockCallerSession) Name() (string, error) {
	return _ERC721Mock.Contract.Name(&_ERC721Mock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721Mock *ERC721MockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721Mock *ERC721MockSession) Owner() (common.Address, error) {
	return _ERC721Mock.Contract.Owner(&_ERC721Mock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721Mock *ERC721MockCallerSession) Owner() (common.Address, error) {
	return _ERC721Mock.Contract.Owner(&_ERC721Mock.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mock.Contract.OwnerOf(&_ERC721Mock.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ERC721Mock *ERC721MockCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Mock.Contract.OwnerOf(&_ERC721Mock.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mock *ERC721MockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mock *ERC721MockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Mock.Contract.SupportsInterface(&_ERC721Mock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC721Mock *ERC721MockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Mock.Contract.SupportsInterface(&_ERC721Mock.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Mock *ERC721MockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Mock *ERC721MockSession) Symbol() (string, error) {
	return _ERC721Mock.Contract.Symbol(&_ERC721Mock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC721Mock *ERC721MockCallerSession) Symbol() (string, error) {
	return _ERC721Mock.Contract.Symbol(&_ERC721Mock.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Mock.Contract.TokenByIndex(&_ERC721Mock.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Mock.Contract.TokenByIndex(&_ERC721Mock.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Mock.Contract.TokenOfOwnerByIndex(&_ERC721Mock.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ERC721Mock *ERC721MockCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Mock.Contract.TokenOfOwnerByIndex(&_ERC721Mock.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Mock *ERC721MockCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Mock *ERC721MockSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Mock.Contract.TokenURI(&_ERC721Mock.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ERC721Mock *ERC721MockCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Mock.Contract.TokenURI(&_ERC721Mock.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Mock *ERC721MockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Mock.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Mock *ERC721MockSession) TotalSupply() (*big.Int, error) {
	return _ERC721Mock.Contract.TotalSupply(&_ERC721Mock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC721Mock *ERC721MockCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Mock.Contract.TotalSupply(&_ERC721Mock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Approve(&_ERC721Mock.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Approve(&_ERC721Mock.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockTransactor) Burn(opts *bind.TransactOpts, tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "burn", tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Burn(&_ERC721Mock.TransactOpts, tokenId_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockTransactorSession) Burn(tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Burn(&_ERC721Mock.TransactOpts, tokenId_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to_, uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockTransactor) Mint(opts *bind.TransactOpts, to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "mint", to_, tokenId_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to_, uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockSession) Mint(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Mint(&_ERC721Mock.TransactOpts, to_, tokenId_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to_, uint256 tokenId_) returns()
func (_ERC721Mock *ERC721MockTransactorSession) Mint(to_ common.Address, tokenId_ *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.Mint(&_ERC721Mock.TransactOpts, to_, tokenId_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721Mock *ERC721MockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721Mock *ERC721MockSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC721Mock.Contract.RenounceOwnership(&_ERC721Mock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721Mock *ERC721MockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC721Mock.Contract.RenounceOwnership(&_ERC721Mock.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SafeTransferFrom(&_ERC721Mock.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SafeTransferFrom(&_ERC721Mock.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721Mock *ERC721MockTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721Mock *ERC721MockSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SafeTransferFrom0(&_ERC721Mock.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ERC721Mock *ERC721MockTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SafeTransferFrom0(&_ERC721Mock.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mock *ERC721MockTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mock *ERC721MockSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SetApprovalForAll(&_ERC721Mock.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ERC721Mock *ERC721MockTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Mock.Contract.SetApprovalForAll(&_ERC721Mock.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.TransferFrom(&_ERC721Mock.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Mock *ERC721MockTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Mock.Contract.TransferFrom(&_ERC721Mock.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721Mock *ERC721MockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC721Mock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721Mock *ERC721MockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC721Mock.Contract.TransferOwnership(&_ERC721Mock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721Mock *ERC721MockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC721Mock.Contract.TransferOwnership(&_ERC721Mock.TransactOpts, newOwner)
}

// ERC721MockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Mock contract.
type ERC721MockApprovalIterator struct {
	Event *ERC721MockApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockApproval)
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
		it.Event = new(ERC721MockApproval)
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
func (it *ERC721MockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockApproval represents a Approval event raised by the ERC721Mock contract.
type ERC721MockApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721MockApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MockApprovalIterator{contract: _ERC721Mock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MockApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockApproval)
				if err := _ERC721Mock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) ParseApproval(log types.Log) (*ERC721MockApproval, error) {
	event := new(ERC721MockApproval)
	if err := _ERC721Mock.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MockApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Mock contract.
type ERC721MockApprovalForAllIterator struct {
	Event *ERC721MockApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MockApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockApprovalForAll)
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
		it.Event = new(ERC721MockApprovalForAll)
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
func (it *ERC721MockApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockApprovalForAll represents a ApprovalForAll event raised by the ERC721Mock contract.
type ERC721MockApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mock *ERC721MockFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721MockApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MockApprovalForAllIterator{contract: _ERC721Mock.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mock *ERC721MockFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MockApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockApprovalForAll)
				if err := _ERC721Mock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Mock *ERC721MockFilterer) ParseApprovalForAll(log types.Log) (*ERC721MockApprovalForAll, error) {
	event := new(ERC721MockApprovalForAll)
	if err := _ERC721Mock.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MockBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ERC721Mock contract.
type ERC721MockBatchMetadataUpdateIterator struct {
	Event *ERC721MockBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC721MockBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockBatchMetadataUpdate)
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
		it.Event = new(ERC721MockBatchMetadataUpdate)
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
func (it *ERC721MockBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ERC721Mock contract.
type ERC721MockBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721Mock *ERC721MockFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ERC721MockBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC721MockBatchMetadataUpdateIterator{contract: _ERC721Mock.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721Mock *ERC721MockFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC721MockBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockBatchMetadataUpdate)
				if err := _ERC721Mock.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ERC721Mock *ERC721MockFilterer) ParseBatchMetadataUpdate(log types.Log) (*ERC721MockBatchMetadataUpdate, error) {
	event := new(ERC721MockBatchMetadataUpdate)
	if err := _ERC721Mock.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MockMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ERC721Mock contract.
type ERC721MockMetadataUpdateIterator struct {
	Event *ERC721MockMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ERC721MockMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockMetadataUpdate)
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
		it.Event = new(ERC721MockMetadataUpdate)
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
func (it *ERC721MockMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockMetadataUpdate represents a MetadataUpdate event raised by the ERC721Mock contract.
type ERC721MockMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721Mock *ERC721MockFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ERC721MockMetadataUpdateIterator, error) {

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ERC721MockMetadataUpdateIterator{contract: _ERC721Mock.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721Mock *ERC721MockFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ERC721MockMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockMetadataUpdate)
				if err := _ERC721Mock.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ERC721Mock *ERC721MockFilterer) ParseMetadataUpdate(log types.Log) (*ERC721MockMetadataUpdate, error) {
	event := new(ERC721MockMetadataUpdate)
	if err := _ERC721Mock.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC721Mock contract.
type ERC721MockOwnershipTransferredIterator struct {
	Event *ERC721MockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC721MockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockOwnershipTransferred)
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
		it.Event = new(ERC721MockOwnershipTransferred)
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
func (it *ERC721MockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockOwnershipTransferred represents a OwnershipTransferred event raised by the ERC721Mock contract.
type ERC721MockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC721Mock *ERC721MockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC721MockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MockOwnershipTransferredIterator{contract: _ERC721Mock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC721Mock *ERC721MockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC721MockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockOwnershipTransferred)
				if err := _ERC721Mock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC721Mock *ERC721MockFilterer) ParseOwnershipTransferred(log types.Log) (*ERC721MockOwnershipTransferred, error) {
	event := new(ERC721MockOwnershipTransferred)
	if err := _ERC721Mock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Mock contract.
type ERC721MockTransferIterator struct {
	Event *ERC721MockTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MockTransfer)
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
		it.Event = new(ERC721MockTransfer)
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
func (it *ERC721MockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MockTransfer represents a Transfer event raised by the ERC721Mock contract.
type ERC721MockTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721MockTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mock.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MockTransferIterator{contract: _ERC721Mock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MockTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Mock.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MockTransfer)
				if err := _ERC721Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Mock *ERC721MockFilterer) ParseTransfer(log types.Log) (*ERC721MockTransfer, error) {
	event := new(ERC721MockTransfer)
	if err := _ERC721Mock.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
