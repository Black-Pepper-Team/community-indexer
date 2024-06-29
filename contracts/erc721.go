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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620019af380380620019af8339810160408190526200003491620001b1565b3382826000620000458382620002aa565b506001620000548282620002aa565b5050506001600160a01b0381166200008657604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b62000091816200009a565b50505062000376565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200011457600080fd5b81516001600160401b0380821115620001315762000131620000ec565b604051601f8301601f19908116603f011681019082821181831017156200015c576200015c620000ec565b816040528381526020925086838588010111156200017957600080fd5b600091505b838210156200019d57858201830151818301840152908201906200017e565b600093810190920192909252949350505050565b60008060408385031215620001c557600080fd5b82516001600160401b0380821115620001dd57600080fd5b620001eb8683870162000102565b935060208501519150808211156200020257600080fd5b50620002118582860162000102565b9150509250929050565b600181811c908216806200023057607f821691505b6020821081036200025157634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a557600081815260208120601f850160051c81016020861015620002805750805b601f850160051c820191505b81811015620002a1578281556001016200028c565b5050505b505050565b81516001600160401b03811115620002c657620002c6620000ec565b620002de81620002d784546200021b565b8462000257565b602080601f831160018114620003165760008415620002fd5750858301515b600019600386901b1c1916600185901b178555620002a1565b600085815260208120601f198616915b82811015620003475788860151825594840194600190910190840162000326565b5085821015620003665787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61162980620003866000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80634f6ccce7116100b857806395d89b411161007c57806395d89b411461027c578063a22cb46514610284578063b88d4fde14610297578063c87b56dd146102aa578063e985e9c5146102bd578063f2fde38b146102d057600080fd5b80634f6ccce71461022a5780636352211e1461023d57806370a0823114610250578063715018a6146102635780638da5cb5b1461026b57600080fd5b806323b872dd116100ff57806323b872dd146101cb5780632f745c59146101de57806340c10f19146101f157806342842e0e1461020457806342966c681461021757600080fd5b806301ffc9a71461013c57806306fdde0314610164578063081812fc14610179578063095ea7b3146101a457806318160ddd146101b9575b600080fd5b61014f61014a36600461124c565b6102e3565b60405190151581526020015b60405180910390f35b61016c6102f4565b60405161015b91906112b9565b61018c6101873660046112cc565b610386565b6040516001600160a01b03909116815260200161015b565b6101b76101b2366004611301565b6103af565b005b6008545b60405190815260200161015b565b6101b76101d936600461132b565b6103be565b6101bd6101ec366004611301565b61044e565b6101b76101ff366004611301565b6104b3565b6101b761021236600461132b565b6104bd565b6101b76102253660046112cc565b6104dd565b6101bd6102383660046112cc565b6104e9565b61018c61024b3660046112cc565b610542565b6101bd61025e366004611367565b61054d565b6101b7610595565b600b546001600160a01b031661018c565b61016c6105a9565b6101b7610292366004611382565b6105b8565b6101b76102a53660046113d4565b6105c3565b61016c6102b83660046112cc565b6105da565b61014f6102cb3660046114b0565b6105e5565b6101b76102de366004611367565b610613565b60006102ee8261064e565b92915050565b606060008054610303906114e3565b80601f016020809104026020016040519081016040528092919081815260200182805461032f906114e3565b801561037c5780601f106103515761010080835404028352916020019161037c565b820191906000526020600020905b81548152906001019060200180831161035f57829003601f168201915b5050505050905090565b600061039182610673565b506000828152600460205260409020546001600160a01b03166102ee565b6103ba8282336106ac565b5050565b6001600160a01b0382166103ed57604051633250574960e11b8152600060048201526024015b60405180910390fd5b60006103fa8383336106b9565b9050836001600160a01b0316816001600160a01b031614610448576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016103e4565b50505050565b60006104598361054d565b821061048a5760405163295f44f760e21b81526001600160a01b0384166004820152602481018390526044016103e4565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b6103ba82826106ce565b6104d8838383604051806020016040528060008152506105c3565b505050565b6104e681610733565b50565b60006104f460085490565b821061051d5760405163295f44f760e21b815260006004820152602481018390526044016103e4565b600882815481106105305761053061151d565b90600052602060002001549050919050565b60006102ee82610673565b60006001600160a01b038216610579576040516322718ad960e21b8152600060048201526024016103e4565b506001600160a01b031660009081526003602052604090205490565b61059d61076e565b6105a7600061079b565b565b606060018054610303906114e3565b6103ba3383836107ed565b6105ce8484846103be565b6104488484848461088c565b60606102ee826109b5565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b61061b61076e565b6001600160a01b03811661064557604051631e4fbdf760e01b8152600060048201526024016103e4565b6104e68161079b565b60006001600160e01b03198216632483248360e11b14806102ee57506102ee82610abe565b6000818152600260205260408120546001600160a01b0316806102ee57604051637e27328960e01b8152600481018490526024016103e4565b6104d88383836001610ae3565b60006106c6848484610be9565b949350505050565b6001600160a01b0382166106f857604051633250574960e11b8152600060048201526024016103e4565b6000610706838360006106b9565b90506001600160a01b038116156104d8576040516339e3563760e11b8152600060048201526024016103e4565b600061074260008360006106b9565b90506001600160a01b0381166103ba57604051637e27328960e01b8152600481018390526024016103e4565b600b546001600160a01b031633146105a75760405163118cdaa760e01b81523360048201526024016103e4565b600b80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03821661081f57604051630b61174360e31b81526001600160a01b03831660048201526024016103e4565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b1561044857604051630a85bd0160e11b81526001600160a01b0384169063150b7a02906108ce903390889087908790600401611533565b6020604051808303816000875af1925050508015610909575060408051601f3d908101601f1916820190925261090691810190611570565b60015b610972573d808015610937576040519150601f19603f3d011682016040523d82523d6000602084013e61093c565b606091505b50805160000361096a57604051633250574960e11b81526001600160a01b03851660048201526024016103e4565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b146109ae57604051633250574960e11b81526001600160a01b03851660048201526024016103e4565b5050505050565b60606109c082610673565b506000828152600a6020526040812080546109da906114e3565b80601f0160208091040260200160405190810160405280929190818152602001828054610a06906114e3565b8015610a535780601f10610a2857610100808354040283529160200191610a53565b820191906000526020600020905b815481529060010190602001808311610a3657829003601f168201915b505050505090506000610a7160408051602081019091526000815290565b90508051600003610a83575092915050565b815115610ab5578082604051602001610a9d92919061158d565b60405160208183030381529060405292505050919050565b6106c684610cb6565b60006001600160e01b0319821663780e9d6360e01b14806102ee57506102ee82610d2b565b8080610af757506001600160a01b03821615155b15610bb9576000610b0784610673565b90506001600160a01b03831615801590610b335750826001600160a01b0316816001600160a01b031614155b8015610b465750610b4481846105e5565b155b15610b6f5760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016103e4565b8115610bb75783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5050600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b600080610bf7858585610d7b565b90506001600160a01b038116610c5457610c4f84600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b610c77565b846001600160a01b0316816001600160a01b031614610c7757610c778185610e74565b6001600160a01b038516610c9357610c8e84610f05565b6106c6565b846001600160a01b0316816001600160a01b0316146106c6576106c68585610fb4565b6060610cc182610673565b506000610cd960408051602081019091526000815290565b90506000815111610cf95760405180602001604052806000815250610d24565b80610d0384611004565b604051602001610d1492919061158d565b6040516020818303038152906040525b9392505050565b60006001600160e01b031982166380ac58cd60e01b1480610d5c57506001600160e01b03198216635b5e139f60e01b145b806102ee57506301ffc9a760e01b6001600160e01b03198316146102ee565b6000828152600260205260408120546001600160a01b0390811690831615610da857610da8818486611097565b6001600160a01b03811615610de657610dc5600085600080610ae3565b6001600160a01b038116600090815260036020526040902080546000190190555b6001600160a01b03851615610e15576001600160a01b0385166000908152600360205260409020805460010190555b60008481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b6000610e7f8361054d565b600083815260076020526040902054909150808214610ed2576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b600854600090610f17906001906115bc565b60008381526009602052604081205460088054939450909284908110610f3f57610f3f61151d565b906000526020600020015490508060088381548110610f6057610f6061151d565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480610f9857610f986115dd565b6001900381819060005260206000200160009055905550505050565b60006001610fc18461054d565b610fcb91906115bc565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b60606000611011836110fb565b600101905060008167ffffffffffffffff811115611031576110316113be565b6040519080825280601f01601f19166020018201604052801561105b576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461106557509392505050565b6110a28383836111d3565b6104d8576001600160a01b0383166110d057604051637e27328960e01b8152600481018290526024016103e4565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016103e4565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b831061113a5772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611166576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061118457662386f26fc10000830492506010015b6305f5e100831061119c576305f5e100830492506008015b61271083106111b057612710830492506004015b606483106111c2576064830492506002015b600a83106102ee5760010192915050565b60006001600160a01b038316158015906106c65750826001600160a01b0316846001600160a01b0316148061120d575061120d84846105e5565b806106c65750506000908152600460205260409020546001600160a01b03908116911614919050565b6001600160e01b0319811681146104e657600080fd5b60006020828403121561125e57600080fd5b8135610d2481611236565b60005b8381101561128457818101518382015260200161126c565b50506000910152565b600081518084526112a5816020860160208601611269565b601f01601f19169290920160200192915050565b602081526000610d24602083018461128d565b6000602082840312156112de57600080fd5b5035919050565b80356001600160a01b03811681146112fc57600080fd5b919050565b6000806040838503121561131457600080fd5b61131d836112e5565b946020939093013593505050565b60008060006060848603121561134057600080fd5b611349846112e5565b9250611357602085016112e5565b9150604084013590509250925092565b60006020828403121561137957600080fd5b610d24826112e5565b6000806040838503121561139557600080fd5b61139e836112e5565b9150602083013580151581146113b357600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b600080600080608085870312156113ea57600080fd5b6113f3856112e5565b9350611401602086016112e5565b925060408501359150606085013567ffffffffffffffff8082111561142557600080fd5b818701915087601f83011261143957600080fd5b81358181111561144b5761144b6113be565b604051601f8201601f19908116603f01168101908382118183101715611473576114736113be565b816040528281528a602084870101111561148c57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b600080604083850312156114c357600080fd5b6114cc836112e5565b91506114da602084016112e5565b90509250929050565b600181811c908216806114f757607f821691505b60208210810361151757634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906115669083018461128d565b9695505050505050565b60006020828403121561158257600080fd5b8151610d2481611236565b6000835161159f818460208801611269565b8351908301906115b3818360208801611269565b01949350505050565b818103818111156102ee57634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603160045260246000fdfea26469706673582212203fdfdffd4078a26d4987d84ddf1a9883d07f7f78a1729a04baa987b803a305a464736f6c63430008140033",
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
