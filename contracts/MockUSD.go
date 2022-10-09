// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MockUSDABI is the input ABI used to generate the binding from.
const MockUSDABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MockUSDBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const MockUSDBinRuntime = ``

// MockUSDBin is the compiled bytecode used for deploying new contracts.
var MockUSDBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600781526020017f4d6f636b555344000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f635553440000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620005db565b508060049081620000a19190620005db565b505050620000c4620000b86200010860201b60201c565b6200011060201b60201c565b6200010233620000d9620001d660201b60201c565b600a620000e7919062000852565b6103e8620000f69190620008a3565b620001df60201b60201c565b62000a12565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000251576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002489062000965565b60405180910390fd5b62000265600083836200035760201b60201c565b806002600082825462000279919062000987565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620002d0919062000987565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620003379190620009f5565b60405180910390a362000353600083836200035c60201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620003e357607f821691505b602082108103620003f957620003f86200039b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004637fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000424565b6200046f868362000424565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620004bc620004b6620004b08462000487565b62000491565b62000487565b9050919050565b6000819050919050565b620004d8836200049b565b620004f0620004e782620004c3565b84845462000431565b825550505050565b600090565b62000507620004f8565b62000514818484620004cd565b505050565b5b818110156200053c5762000530600082620004fd565b6001810190506200051a565b5050565b601f8211156200058b576200055581620003ff565b620005608462000414565b8101602085101562000570578190505b620005886200057f8562000414565b83018262000519565b50505b505050565b600082821c905092915050565b6000620005b06000198460080262000590565b1980831691505092915050565b6000620005cb83836200059d565b9150826002028217905092915050565b620005e68262000361565b67ffffffffffffffff8111156200060257620006016200036c565b5b6200060e8254620003ca565b6200061b82828562000540565b600060209050601f8311600181146200065357600084156200063e578287015190505b6200064a8582620005bd565b865550620006ba565b601f1984166200066386620003ff565b60005b828110156200068d5784890151825560018201915060208501945060208101905062000666565b86831015620006ad5784890151620006a9601f8916826200059d565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156200075057808604811115620007285762000727620006c2565b5b6001851615620007385780820291505b80810290506200074885620006f1565b945062000708565b94509492505050565b6000826200076b57600190506200083e565b816200077b57600090506200083e565b81600181146200079457600281146200079f57620007d5565b60019150506200083e565b60ff841115620007b457620007b3620006c2565b5b8360020a915084821115620007ce57620007cd620006c2565b5b506200083e565b5060208310610133831016604e8410600b84101617156200080f5782820a905083811115620008095762000808620006c2565b5b6200083e565b6200081e8484846001620006fe565b92509050818404811115620008385762000837620006c2565b5b81810290505b9392505050565b600060ff82169050919050565b60006200085f8262000487565b91506200086c8362000845565b92506200089b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848462000759565b905092915050565b6000620008b08262000487565b9150620008bd8362000487565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620008f957620008f8620006c2565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200094d601f8362000904565b91506200095a8262000915565b602082019050919050565b6000602082019050818103600083015262000980816200093e565b9050919050565b6000620009948262000487565b9150620009a18362000487565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620009d957620009d8620006c2565b5b828201905092915050565b620009ef8162000487565b82525050565b600060208201905062000a0c6000830184620009e4565b92915050565b61181d8062000a226000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d714610276578063a9059cbb146102a6578063dd62ed3e146102d6578063f2fde38b14610306576100f5565b806370a0823114610200578063715018a6146102305780638da5cb5b1461023a57806395d89b4114610258576100f5565b806323b872dd116100d357806323b872dd14610166578063313ce5671461019657806339509351146101b457806340c10f19146101e4576100f5565b806306fdde03146100fa578063095ea7b31461011857806318160ddd14610148575b600080fd5b610102610322565b60405161010f9190610f44565b60405180910390f35b610132600480360381019061012d9190610fff565b6103b4565b60405161013f919061105a565b60405180910390f35b6101506103d7565b60405161015d9190611084565b60405180910390f35b610180600480360381019061017b919061109f565b6103e1565b60405161018d919061105a565b60405180910390f35b61019e610410565b6040516101ab919061110e565b60405180910390f35b6101ce60048036038101906101c99190610fff565b610419565b6040516101db919061105a565b60405180910390f35b6101fe60048036038101906101f99190610fff565b610450565b005b61021a60048036038101906102159190611129565b610466565b6040516102279190611084565b60405180910390f35b6102386104ae565b005b6102426104c2565b60405161024f9190611165565b60405180910390f35b6102606104ec565b60405161026d9190610f44565b60405180910390f35b610290600480360381019061028b9190610fff565b61057e565b60405161029d919061105a565b60405180910390f35b6102c060048036038101906102bb9190610fff565b6105f5565b6040516102cd919061105a565b60405180910390f35b6102f060048036038101906102eb9190611180565b610618565b6040516102fd9190611084565b60405180910390f35b610320600480360381019061031b9190611129565b61069f565b005b606060038054610331906111ef565b80601f016020809104026020016040519081016040528092919081815260200182805461035d906111ef565b80156103aa5780601f1061037f576101008083540402835291602001916103aa565b820191906000526020600020905b81548152906001019060200180831161038d57829003601f168201915b5050505050905090565b6000806103bf610722565b90506103cc81858561072a565b600191505092915050565b6000600254905090565b6000806103ec610722565b90506103f98582856108f3565b61040485858561097f565b60019150509392505050565b60006012905090565b600080610424610722565b90506104458185856104368589610618565b610440919061124f565b61072a565b600191505092915050565b610458610bfe565b6104628282610c7c565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6104b6610bfe565b6104c06000610ddb565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546104fb906111ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610527906111ef565b80156105745780601f1061054957610100808354040283529160200191610574565b820191906000526020600020905b81548152906001019060200180831161055757829003601f168201915b5050505050905090565b600080610589610722565b905060006105978286610618565b9050838110156105dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d390611317565b60405180910390fd5b6105e9828686840361072a565b60019250505092915050565b600080610600610722565b905061060d81858561097f565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6106a7610bfe565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070d906113a9565b60405180910390fd5b61071f81610ddb565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610799576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107909061143b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ff906114cd565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516108e69190611084565b60405180910390a3505050565b60006108ff8484610618565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610979578181101561096b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096290611539565b60405180910390fd5b610978848484840361072a565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e5906115cb565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a549061165d565b60405180910390fd5b610a68838383610ea1565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610aee576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae5906116ef565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b81919061124f565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610be59190611084565b60405180910390a3610bf8848484610ea6565b50505050565b610c06610722565b73ffffffffffffffffffffffffffffffffffffffff16610c246104c2565b73ffffffffffffffffffffffffffffffffffffffff1614610c7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c719061175b565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610ceb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce2906117c7565b60405180910390fd5b610cf760008383610ea1565b8060026000828254610d09919061124f565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d5e919061124f565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610dc39190611084565b60405180910390a3610dd760008383610ea6565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ee5578082015181840152602081019050610eca565b83811115610ef4576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f1682610eab565b610f208185610eb6565b9350610f30818560208601610ec7565b610f3981610efa565b840191505092915050565b60006020820190508181036000830152610f5e8184610f0b565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f9682610f6b565b9050919050565b610fa681610f8b565b8114610fb157600080fd5b50565b600081359050610fc381610f9d565b92915050565b6000819050919050565b610fdc81610fc9565b8114610fe757600080fd5b50565b600081359050610ff981610fd3565b92915050565b6000806040838503121561101657611015610f66565b5b600061102485828601610fb4565b925050602061103585828601610fea565b9150509250929050565b60008115159050919050565b6110548161103f565b82525050565b600060208201905061106f600083018461104b565b92915050565b61107e81610fc9565b82525050565b60006020820190506110996000830184611075565b92915050565b6000806000606084860312156110b8576110b7610f66565b5b60006110c686828701610fb4565b93505060206110d786828701610fb4565b92505060406110e886828701610fea565b9150509250925092565b600060ff82169050919050565b611108816110f2565b82525050565b600060208201905061112360008301846110ff565b92915050565b60006020828403121561113f5761113e610f66565b5b600061114d84828501610fb4565b91505092915050565b61115f81610f8b565b82525050565b600060208201905061117a6000830184611156565b92915050565b6000806040838503121561119757611196610f66565b5b60006111a585828601610fb4565b92505060206111b685828601610fb4565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061120757607f821691505b60208210810361121a576112196111c0565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061125a82610fc9565b915061126583610fc9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561129a57611299611220565b5b828201905092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611301602583610eb6565b915061130c826112a5565b604082019050919050565b60006020820190508181036000830152611330816112f4565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611393602683610eb6565b915061139e82611337565b604082019050919050565b600060208201905081810360008301526113c281611386565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611425602483610eb6565b9150611430826113c9565b604082019050919050565b6000602082019050818103600083015261145481611418565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006114b7602283610eb6565b91506114c28261145b565b604082019050919050565b600060208201905081810360008301526114e6816114aa565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611523601d83610eb6565b915061152e826114ed565b602082019050919050565b6000602082019050818103600083015261155281611516565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006115b5602583610eb6565b91506115c082611559565b604082019050919050565b600060208201905081810360008301526115e4816115a8565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611647602383610eb6565b9150611652826115eb565b604082019050919050565b600060208201905081810360008301526116768161163a565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006116d9602683610eb6565b91506116e48261167d565b604082019050919050565b60006020820190508181036000830152611708816116cc565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611745602083610eb6565b91506117508261170f565b602082019050919050565b6000602082019050818103600083015261177481611738565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006117b1601f83610eb6565b91506117bc8261177b565b602082019050919050565b600060208201905081810360008301526117e0816117a4565b905091905056fea26469706673582212203abed7e10276b4bd5352d532b7e2ed51e1564013fe8ee1d928766c5b857d636864736f6c634300080f0033"

// DeployMockUSD deploys a new Klaytn contract, binding an instance of MockUSD to it.
func DeployMockUSD(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(MockUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MockUSDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockUSD{MockUSDCaller: MockUSDCaller{contract: contract}, MockUSDTransactor: MockUSDTransactor{contract: contract}, MockUSDFilterer: MockUSDFilterer{contract: contract}}, nil
}

// MockUSD is an auto generated Go binding around a Klaytn contract.
type MockUSD struct {
	MockUSDCaller     // Read-only binding to the contract
	MockUSDTransactor // Write-only binding to the contract
	MockUSDFilterer   // Log filterer for contract events
}

// MockUSDCaller is an auto generated read-only Go binding around a Klaytn contract.
type MockUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDTransactor is an auto generated write-only Go binding around a Klaytn contract.
type MockUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type MockUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type MockUSDSession struct {
	Contract     *MockUSD          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockUSDCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type MockUSDCallerSession struct {
	Contract *MockUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MockUSDTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type MockUSDTransactorSession struct {
	Contract     *MockUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MockUSDRaw is an auto generated low-level Go binding around a Klaytn contract.
type MockUSDRaw struct {
	Contract *MockUSD // Generic contract binding to access the raw methods on
}

// MockUSDCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type MockUSDCallerRaw struct {
	Contract *MockUSDCaller // Generic read-only contract binding to access the raw methods on
}

// MockUSDTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type MockUSDTransactorRaw struct {
	Contract *MockUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockUSD creates a new instance of MockUSD, bound to a specific deployed contract.
func NewMockUSD(address common.Address, backend bind.ContractBackend) (*MockUSD, error) {
	contract, err := bindMockUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUSD{MockUSDCaller: MockUSDCaller{contract: contract}, MockUSDTransactor: MockUSDTransactor{contract: contract}, MockUSDFilterer: MockUSDFilterer{contract: contract}}, nil
}

// NewMockUSDCaller creates a new read-only instance of MockUSD, bound to a specific deployed contract.
func NewMockUSDCaller(address common.Address, caller bind.ContractCaller) (*MockUSDCaller, error) {
	contract, err := bindMockUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCaller{contract: contract}, nil
}

// NewMockUSDTransactor creates a new write-only instance of MockUSD, bound to a specific deployed contract.
func NewMockUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUSDTransactor, error) {
	contract, err := bindMockUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDTransactor{contract: contract}, nil
}

// NewMockUSDFilterer creates a new log filterer instance of MockUSD, bound to a specific deployed contract.
func NewMockUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUSDFilterer, error) {
	contract, err := bindMockUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUSDFilterer{contract: contract}, nil
}

// bindMockUSD binds a generic wrapper to an already deployed contract.
func bindMockUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUSD *MockUSDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MockUSD.Contract.MockUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUSD *MockUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSD.Contract.MockUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUSD *MockUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSD.Contract.MockUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUSD *MockUSDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MockUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUSD *MockUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUSD *MockUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSD.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSD *MockUSDCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSD *MockUSDSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockUSD.Contract.Allowance(&_MockUSD.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSD *MockUSDCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockUSD.Contract.Allowance(&_MockUSD.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSD *MockUSDCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSD *MockUSDSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUSD.Contract.BalanceOf(&_MockUSD.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSD *MockUSDCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUSD.Contract.BalanceOf(&_MockUSD.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockUSD *MockUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockUSD *MockUSDSession) Decimals() (uint8, error) {
	return _MockUSD.Contract.Decimals(&_MockUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MockUSD *MockUSDCallerSession) Decimals() (uint8, error) {
	return _MockUSD.Contract.Decimals(&_MockUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSD *MockUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSD *MockUSDSession) Name() (string, error) {
	return _MockUSD.Contract.Name(&_MockUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSD *MockUSDCallerSession) Name() (string, error) {
	return _MockUSD.Contract.Name(&_MockUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockUSD *MockUSDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockUSD *MockUSDSession) Owner() (common.Address, error) {
	return _MockUSD.Contract.Owner(&_MockUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockUSD *MockUSDCallerSession) Owner() (common.Address, error) {
	return _MockUSD.Contract.Owner(&_MockUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSD *MockUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSD *MockUSDSession) Symbol() (string, error) {
	return _MockUSD.Contract.Symbol(&_MockUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSD *MockUSDCallerSession) Symbol() (string, error) {
	return _MockUSD.Contract.Symbol(&_MockUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSD *MockUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MockUSD.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSD *MockUSDSession) TotalSupply() (*big.Int, error) {
	return _MockUSD.Contract.TotalSupply(&_MockUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSD *MockUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _MockUSD.Contract.TotalSupply(&_MockUSD.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockUSD *MockUSDSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Approve(&_MockUSD.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Approve(&_MockUSD.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockUSD *MockUSDTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockUSD *MockUSDSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.DecreaseAllowance(&_MockUSD.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MockUSD *MockUSDTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.DecreaseAllowance(&_MockUSD.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockUSD *MockUSDTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockUSD *MockUSDSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.IncreaseAllowance(&_MockUSD.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MockUSD *MockUSDTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.IncreaseAllowance(&_MockUSD.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSD *MockUSDTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSD *MockUSDSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Mint(&_MockUSD.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSD *MockUSDTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Mint(&_MockUSD.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockUSD *MockUSDTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockUSD *MockUSDSession) RenounceOwnership() (*types.Transaction, error) {
	return _MockUSD.Contract.RenounceOwnership(&_MockUSD.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MockUSD *MockUSDTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MockUSD.Contract.RenounceOwnership(&_MockUSD.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Transfer(&_MockUSD.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.Transfer(&_MockUSD.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.TransferFrom(&_MockUSD.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_MockUSD *MockUSDTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSD.Contract.TransferFrom(&_MockUSD.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockUSD *MockUSDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MockUSD.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockUSD *MockUSDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MockUSD.Contract.TransferOwnership(&_MockUSD.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MockUSD *MockUSDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MockUSD.Contract.TransferOwnership(&_MockUSD.TransactOpts, newOwner)
}

// MockUSDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockUSD contract.
type MockUSDApprovalIterator struct {
	Event *MockUSDApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUSDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDApproval)
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
		it.Event = new(MockUSDApproval)
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
func (it *MockUSDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUSDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUSDApproval represents a Approval event raised by the MockUSD contract.
type MockUSDApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSD *MockUSDFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockUSDApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockUSD.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDApprovalIterator{contract: _MockUSD.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSD *MockUSDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockUSDApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockUSD.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUSDApproval)
				if err := _MockUSD.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSD *MockUSDFilterer) ParseApproval(log types.Log) (*MockUSDApproval, error) {
	event := new(MockUSDApproval)
	if err := _MockUSD.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MockUSDOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MockUSD contract.
type MockUSDOwnershipTransferredIterator struct {
	Event *MockUSDOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUSDOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDOwnershipTransferred)
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
		it.Event = new(MockUSDOwnershipTransferred)
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
func (it *MockUSDOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUSDOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUSDOwnershipTransferred represents a OwnershipTransferred event raised by the MockUSD contract.
type MockUSDOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MockUSD *MockUSDFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MockUSDOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MockUSD.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDOwnershipTransferredIterator{contract: _MockUSD.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MockUSD *MockUSDFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockUSDOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MockUSD.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUSDOwnershipTransferred)
				if err := _MockUSD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MockUSD *MockUSDFilterer) ParseOwnershipTransferred(log types.Log) (*MockUSDOwnershipTransferred, error) {
	event := new(MockUSDOwnershipTransferred)
	if err := _MockUSD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MockUSDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockUSD contract.
type MockUSDTransferIterator struct {
	Event *MockUSDTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockUSDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDTransfer)
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
		it.Event = new(MockUSDTransfer)
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
func (it *MockUSDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUSDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUSDTransfer represents a Transfer event raised by the MockUSD contract.
type MockUSDTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSD *MockUSDFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockUSDTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockUSD.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDTransferIterator{contract: _MockUSD.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSD *MockUSDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockUSDTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockUSD.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUSDTransfer)
				if err := _MockUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSD *MockUSDFilterer) ParseTransfer(log types.Log) (*MockUSDTransfer, error) {
	event := new(MockUSDTransfer)
	if err := _MockUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
