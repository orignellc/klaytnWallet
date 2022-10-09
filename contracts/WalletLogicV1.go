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

// WalletLogicV1ABI is the input ABI used to generate the binding from.
const WalletLogicV1ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"newFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"quoteId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"customer\",\"type\":\"address\"}],\"name\":\"newPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"newUnFreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"witpendingFee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"newWithdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"quoteId\",\"type\":\"string\"}],\"name\":\"execTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"freeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalP2P\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unfreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WalletLogicV1BinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const WalletLogicV1BinRuntime = ``

// WalletLogicV1Bin is the compiled bytecode used for deploying new contracts.
var WalletLogicV1Bin = "0x608060405234801561001057600080fd5b50611500806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635c2e3a7d116100715780635c2e3a7d146101425780636623fc461461015e57806395e17d841461017a578063a890bd1514610196578063c55dae63146101b4578063d7a78db8146101d2576100a9565b8063224438d1146100ae57806326e6074b146100cc5780632a9a872a146100ea578063522f68151461010857806354fd4d5014610124575b600080fd5b6100b66101ee565b6040516100c39190610c97565b60405180910390f35b6100d46101f4565b6040516100e19190610c97565b60405180910390f35b6100f26101fa565b6040516100ff9190610c97565b60405180910390f35b610122600480360381019061011d9190610d50565b610209565b005b61012c610350565b6040516101399190610e29565b60405180910390f35b61015c60048036038101906101579190610f80565b61038d565b005b61017860048036038101906101739190611003565b6105c3565b005b610194600480360381019061018f9190611030565b61071d565b005b61019e610a0e565b6040516101ab9190611092565b60405180910390f35b6101bc610a32565b6040516101c991906110ce565b60405180910390f35b6101ec60048036038101906101e79190611003565b610a58565b005b60335481565b60345481565b6000610204610bb2565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610274573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102989190611115565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610305576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fc9061118e565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561034b573d6000803e3d6000fd5b505050565b60606040518060400160405280600681526020017f76312e312e300000000000000000000000000000000000000000000000000000815250905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041c9190611115565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610489576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104809061118e565b60405180910390fd5b828461049591906111dd565b61049d610bb2565b116104dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d49061127f565b60405180910390fd5b82603360008282546104ef91906111dd565b925050819055506104fe610c54565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83866040518363ffffffff1660e01b81526004016105389291906112fe565b6020604051808303816000875af1158015610557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057b919061135f565b507f404fae266883113fdb648f91507a22fd49f177aaed47f51c97349b02dacbf489818560335486866040516105b595949392919061138c565b60405180910390a150505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561062e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106529190611115565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b69061118e565b60405180910390fd5b80603460008282546106d191906113e6565b925050819055507f49f0e6e6a9465a79ab623e1165dc01f7015b14103f87f8a80b565a2cc8d79eb781610702610bb2565b306040516107129392919061141a565b60405180910390a150565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610788573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ac9190611115565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610819576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108109061118e565b60405180910390fd5b6000828461082791906111dd565b61082f610bb2565b1015905080610873576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086a9061127f565b60405180910390fd5b6000603381905550610883610c54565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856033546108cd91906111dd565b6040518363ffffffff1660e01b81526004016108ea9291906112fe565b6020604051808303816000875af1158015610909573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092d919061135f565b50610936610c54565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83866040518363ffffffff1660e01b81526004016109709291906112fe565b6020604051808303816000875af115801561098f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b3919061135f565b508173ffffffffffffffffffffffffffffffffffffffff167f1be5c93cb35ef6c42bc250de0f3835c641d0ab3b50abed71892f8890156526c88560335486604051610a0093929190611451565b60405180910390a250505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ac3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae79190611115565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4b9061118e565b60405180910390fd5b8060346000828254610b6691906111dd565b925050819055507fd7c98bdf591568a7d44ba2dd863ec060c0e47aa9f636f8725a6865f3c8de24b181610b97610bb2565b30604051610ba79392919061141a565b60405180910390a150565b6000603354603454610bc491906111dd565b610bcc610c54565b73ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610c0491906110ce565b602060405180830381865afa158015610c21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c45919061149d565b610c4f91906113e6565b905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000819050919050565b610c9181610c7e565b82525050565b6000602082019050610cac6000830184610c88565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610cf182610cc6565b9050919050565b610d0181610ce6565b8114610d0c57600080fd5b50565b600081359050610d1e81610cf8565b92915050565b610d2d81610c7e565b8114610d3857600080fd5b50565b600081359050610d4a81610d24565b92915050565b60008060408385031215610d6757610d66610cbc565b5b6000610d7585828601610d0f565b9250506020610d8685828601610d3b565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610dca578082015181840152602081019050610daf565b83811115610dd9576000848401525b50505050565b6000601f19601f8301169050919050565b6000610dfb82610d90565b610e058185610d9b565b9350610e15818560208601610dac565b610e1e81610ddf565b840191505092915050565b60006020820190508181036000830152610e438184610df0565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e8d82610ddf565b810181811067ffffffffffffffff82111715610eac57610eab610e55565b5b80604052505050565b6000610ebf610cb2565b9050610ecb8282610e84565b919050565b600067ffffffffffffffff821115610eeb57610eea610e55565b5b610ef482610ddf565b9050602081019050919050565b82818337600083830152505050565b6000610f23610f1e84610ed0565b610eb5565b905082815260208101848484011115610f3f57610f3e610e50565b5b610f4a848285610f01565b509392505050565b600082601f830112610f6757610f66610e4b565b5b8135610f77848260208601610f10565b91505092915050565b60008060008060808587031215610f9a57610f99610cbc565b5b6000610fa887828801610d3b565b9450506020610fb987828801610d3b565b9350506040610fca87828801610d0f565b925050606085013567ffffffffffffffff811115610feb57610fea610cc1565b5b610ff787828801610f52565b91505092959194509250565b60006020828403121561101957611018610cbc565b5b600061102784828501610d3b565b91505092915050565b60008060006060848603121561104957611048610cbc565b5b600061105786828701610d3b565b935050602061106886828701610d3b565b925050604061107986828701610d0f565b9150509250925092565b61108c81610ce6565b82525050565b60006020820190506110a76000830184611083565b92915050565b60006110b882610cc6565b9050919050565b6110c8816110ad565b82525050565b60006020820190506110e360008301846110bf565b92915050565b6110f2816110ad565b81146110fd57600080fd5b50565b60008151905061110f816110e9565b92915050565b60006020828403121561112b5761112a610cbc565b5b600061113984828501611100565b91505092915050565b7f57503a20556e617574686f72697a656400000000000000000000000000000000600082015250565b6000611178601083610d9b565b915061118382611142565b602082019050919050565b600060208201905081810360008301526111a78161116b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006111e882610c7e565b91506111f383610c7e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611228576112276111ae565b5b828201905092915050565b7f574c3a20496e73756666696369656e742062616c616e63650000000000000000600082015250565b6000611269601883610d9b565b915061127482611233565b602082019050919050565b600060208201905081810360008301526112988161125c565b9050919050565b6000819050919050565b60006112c46112bf6112ba84610cc6565b61129f565b610cc6565b9050919050565b60006112d6826112a9565b9050919050565b60006112e8826112cb565b9050919050565b6112f8816112dd565b82525050565b600060408201905061131360008301856112ef565b6113206020830184610c88565b9392505050565b60008115159050919050565b61133c81611327565b811461134757600080fd5b50565b60008151905061135981611333565b92915050565b60006020828403121561137557611374610cbc565b5b60006113838482850161134a565b91505092915050565b600060a08201905081810360008301526113a68188610df0565b90506113b56020830187610c88565b6113c26040830186610c88565b6113cf6060830185610c88565b6113dc60808301846112ef565b9695505050505050565b60006113f182610c7e565b91506113fc83610c7e565b92508282101561140f5761140e6111ae565b5b828203905092915050565b600060608201905061142f6000830186610c88565b61143c6020830185610c88565b61144960408301846110bf565b949350505050565b60006060820190506114666000830186610c88565b6114736020830185610c88565b6114806040830184610c88565b949350505050565b60008151905061149781610d24565b92915050565b6000602082840312156114b3576114b2610cbc565b5b60006114c184828501611488565b9150509291505056fea26469706673582212201505d98a13f5ed5bff2128466095c1fa6fcb5d77b3697026613ad6195cadcbe964736f6c634300080f0033"

// DeployWalletLogicV1 deploys a new Klaytn contract, binding an instance of WalletLogicV1 to it.
func DeployWalletLogicV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletLogicV1, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletLogicV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletLogicV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletLogicV1{WalletLogicV1Caller: WalletLogicV1Caller{contract: contract}, WalletLogicV1Transactor: WalletLogicV1Transactor{contract: contract}, WalletLogicV1Filterer: WalletLogicV1Filterer{contract: contract}}, nil
}

// WalletLogicV1 is an auto generated Go binding around a Klaytn contract.
type WalletLogicV1 struct {
	WalletLogicV1Caller     // Read-only binding to the contract
	WalletLogicV1Transactor // Write-only binding to the contract
	WalletLogicV1Filterer   // Log filterer for contract events
}

// WalletLogicV1Caller is an auto generated read-only Go binding around a Klaytn contract.
type WalletLogicV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLogicV1Transactor is an auto generated write-only Go binding around a Klaytn contract.
type WalletLogicV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLogicV1Filterer is an auto generated log filtering Go binding around a Klaytn contract events.
type WalletLogicV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletLogicV1Session is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type WalletLogicV1Session struct {
	Contract     *WalletLogicV1    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletLogicV1CallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type WalletLogicV1CallerSession struct {
	Contract *WalletLogicV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletLogicV1TransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type WalletLogicV1TransactorSession struct {
	Contract     *WalletLogicV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletLogicV1Raw is an auto generated low-level Go binding around a Klaytn contract.
type WalletLogicV1Raw struct {
	Contract *WalletLogicV1 // Generic contract binding to access the raw methods on
}

// WalletLogicV1CallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type WalletLogicV1CallerRaw struct {
	Contract *WalletLogicV1Caller // Generic read-only contract binding to access the raw methods on
}

// WalletLogicV1TransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type WalletLogicV1TransactorRaw struct {
	Contract *WalletLogicV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletLogicV1 creates a new instance of WalletLogicV1, bound to a specific deployed contract.
func NewWalletLogicV1(address common.Address, backend bind.ContractBackend) (*WalletLogicV1, error) {
	contract, err := bindWalletLogicV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1{WalletLogicV1Caller: WalletLogicV1Caller{contract: contract}, WalletLogicV1Transactor: WalletLogicV1Transactor{contract: contract}, WalletLogicV1Filterer: WalletLogicV1Filterer{contract: contract}}, nil
}

// NewWalletLogicV1Caller creates a new read-only instance of WalletLogicV1, bound to a specific deployed contract.
func NewWalletLogicV1Caller(address common.Address, caller bind.ContractCaller) (*WalletLogicV1Caller, error) {
	contract, err := bindWalletLogicV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1Caller{contract: contract}, nil
}

// NewWalletLogicV1Transactor creates a new write-only instance of WalletLogicV1, bound to a specific deployed contract.
func NewWalletLogicV1Transactor(address common.Address, transactor bind.ContractTransactor) (*WalletLogicV1Transactor, error) {
	contract, err := bindWalletLogicV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1Transactor{contract: contract}, nil
}

// NewWalletLogicV1Filterer creates a new log filterer instance of WalletLogicV1, bound to a specific deployed contract.
func NewWalletLogicV1Filterer(address common.Address, filterer bind.ContractFilterer) (*WalletLogicV1Filterer, error) {
	contract, err := bindWalletLogicV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1Filterer{contract: contract}, nil
}

// bindWalletLogicV1 binds a generic wrapper to an already deployed contract.
func bindWalletLogicV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletLogicV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletLogicV1 *WalletLogicV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletLogicV1.Contract.WalletLogicV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletLogicV1 *WalletLogicV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WalletLogicV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletLogicV1 *WalletLogicV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WalletLogicV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletLogicV1 *WalletLogicV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletLogicV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletLogicV1 *WalletLogicV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletLogicV1 *WalletLogicV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.contract.Transact(opts, method, params...)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_WalletLogicV1 *WalletLogicV1Caller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletLogicV1.contract.Call(opts, out, "baseToken")
	return *ret0, err
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_WalletLogicV1 *WalletLogicV1Session) BaseToken() (common.Address, error) {
	return _WalletLogicV1.Contract.BaseToken(&_WalletLogicV1.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_WalletLogicV1 *WalletLogicV1CallerSession) BaseToken() (common.Address, error) {
	return _WalletLogicV1.Contract.BaseToken(&_WalletLogicV1.CallOpts)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x26e6074b.
//
// Solidity: function frozenBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Caller) FrozenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletLogicV1.contract.Call(opts, out, "frozenBalance")
	return *ret0, err
}

// FrozenBalance is a free data retrieval call binding the contract method 0x26e6074b.
//
// Solidity: function frozenBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Session) FrozenBalance() (*big.Int, error) {
	return _WalletLogicV1.Contract.FrozenBalance(&_WalletLogicV1.CallOpts)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x26e6074b.
//
// Solidity: function frozenBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1CallerSession) FrozenBalance() (*big.Int, error) {
	return _WalletLogicV1.Contract.FrozenBalance(&_WalletLogicV1.CallOpts)
}

// GlobalP2P is a free data retrieval call binding the contract method 0xa890bd15.
//
// Solidity: function globalP2P() view returns(address)
func (_WalletLogicV1 *WalletLogicV1Caller) GlobalP2P(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WalletLogicV1.contract.Call(opts, out, "globalP2P")
	return *ret0, err
}

// GlobalP2P is a free data retrieval call binding the contract method 0xa890bd15.
//
// Solidity: function globalP2P() view returns(address)
func (_WalletLogicV1 *WalletLogicV1Session) GlobalP2P() (common.Address, error) {
	return _WalletLogicV1.Contract.GlobalP2P(&_WalletLogicV1.CallOpts)
}

// GlobalP2P is a free data retrieval call binding the contract method 0xa890bd15.
//
// Solidity: function globalP2P() view returns(address)
func (_WalletLogicV1 *WalletLogicV1CallerSession) GlobalP2P() (common.Address, error) {
	return _WalletLogicV1.Contract.GlobalP2P(&_WalletLogicV1.CallOpts)
}

// PendingFees is a free data retrieval call binding the contract method 0x224438d1.
//
// Solidity: function pendingFees() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Caller) PendingFees(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletLogicV1.contract.Call(opts, out, "pendingFees")
	return *ret0, err
}

// PendingFees is a free data retrieval call binding the contract method 0x224438d1.
//
// Solidity: function pendingFees() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Session) PendingFees() (*big.Int, error) {
	return _WalletLogicV1.Contract.PendingFees(&_WalletLogicV1.CallOpts)
}

// PendingFees is a free data retrieval call binding the contract method 0x224438d1.
//
// Solidity: function pendingFees() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1CallerSession) PendingFees() (*big.Int, error) {
	return _WalletLogicV1.Contract.PendingFees(&_WalletLogicV1.CallOpts)
}

// SpendableBalance is a free data retrieval call binding the contract method 0x2a9a872a.
//
// Solidity: function spendableBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Caller) SpendableBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WalletLogicV1.contract.Call(opts, out, "spendableBalance")
	return *ret0, err
}

// SpendableBalance is a free data retrieval call binding the contract method 0x2a9a872a.
//
// Solidity: function spendableBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1Session) SpendableBalance() (*big.Int, error) {
	return _WalletLogicV1.Contract.SpendableBalance(&_WalletLogicV1.CallOpts)
}

// SpendableBalance is a free data retrieval call binding the contract method 0x2a9a872a.
//
// Solidity: function spendableBalance() view returns(uint256)
func (_WalletLogicV1 *WalletLogicV1CallerSession) SpendableBalance() (*big.Int, error) {
	return _WalletLogicV1.Contract.SpendableBalance(&_WalletLogicV1.CallOpts)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x5c2e3a7d.
//
// Solidity: function execTransaction(uint256 amount, uint256 fee, address customer, string quoteId) returns()
func (_WalletLogicV1 *WalletLogicV1Transactor) ExecTransaction(opts *bind.TransactOpts, amount *big.Int, fee *big.Int, customer common.Address, quoteId string) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "execTransaction", amount, fee, customer, quoteId)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x5c2e3a7d.
//
// Solidity: function execTransaction(uint256 amount, uint256 fee, address customer, string quoteId) returns()
func (_WalletLogicV1 *WalletLogicV1Session) ExecTransaction(amount *big.Int, fee *big.Int, customer common.Address, quoteId string) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.ExecTransaction(&_WalletLogicV1.TransactOpts, amount, fee, customer, quoteId)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x5c2e3a7d.
//
// Solidity: function execTransaction(uint256 amount, uint256 fee, address customer, string quoteId) returns()
func (_WalletLogicV1 *WalletLogicV1TransactorSession) ExecTransaction(amount *big.Int, fee *big.Int, customer common.Address, quoteId string) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.ExecTransaction(&_WalletLogicV1.TransactOpts, amount, fee, customer, quoteId)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Transactor) Freeze(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "freeze", amount)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Session) Freeze(amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Freeze(&_WalletLogicV1.TransactOpts, amount)
}

// Freeze is a paid mutator transaction binding the contract method 0xd7a78db8.
//
// Solidity: function freeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1TransactorSession) Freeze(amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Freeze(&_WalletLogicV1.TransactOpts, amount)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Transactor) Unfreeze(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "unfreeze", amount)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Session) Unfreeze(amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Unfreeze(&_WalletLogicV1.TransactOpts, amount)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x6623fc46.
//
// Solidity: function unfreeze(uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1TransactorSession) Unfreeze(amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Unfreeze(&_WalletLogicV1.TransactOpts, amount)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_WalletLogicV1 *WalletLogicV1Transactor) Version(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "version")
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_WalletLogicV1 *WalletLogicV1Session) Version() (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Version(&_WalletLogicV1.TransactOpts)
}

// Version is a paid mutator transaction binding the contract method 0x54fd4d50.
//
// Solidity: function version() returns(string)
func (_WalletLogicV1 *WalletLogicV1TransactorSession) Version() (*types.Transaction, error) {
	return _WalletLogicV1.Contract.Version(&_WalletLogicV1.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address to, uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Transactor) WithdrawEther(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "withdrawEther", to, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address to, uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1Session) WithdrawEther(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WithdrawEther(&_WalletLogicV1.TransactOpts, to, amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x522f6815.
//
// Solidity: function withdrawEther(address to, uint256 amount) returns()
func (_WalletLogicV1 *WalletLogicV1TransactorSession) WithdrawEther(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WithdrawEther(&_WalletLogicV1.TransactOpts, to, amount)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x95e17d84.
//
// Solidity: function withdrawLiquidity(uint256 amount, uint256 withdrawalFee, address to) returns()
func (_WalletLogicV1 *WalletLogicV1Transactor) WithdrawLiquidity(opts *bind.TransactOpts, amount *big.Int, withdrawalFee *big.Int, to common.Address) (*types.Transaction, error) {
	return _WalletLogicV1.contract.Transact(opts, "withdrawLiquidity", amount, withdrawalFee, to)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x95e17d84.
//
// Solidity: function withdrawLiquidity(uint256 amount, uint256 withdrawalFee, address to) returns()
func (_WalletLogicV1 *WalletLogicV1Session) WithdrawLiquidity(amount *big.Int, withdrawalFee *big.Int, to common.Address) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WithdrawLiquidity(&_WalletLogicV1.TransactOpts, amount, withdrawalFee, to)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x95e17d84.
//
// Solidity: function withdrawLiquidity(uint256 amount, uint256 withdrawalFee, address to) returns()
func (_WalletLogicV1 *WalletLogicV1TransactorSession) WithdrawLiquidity(amount *big.Int, withdrawalFee *big.Int, to common.Address) (*types.Transaction, error) {
	return _WalletLogicV1.Contract.WithdrawLiquidity(&_WalletLogicV1.TransactOpts, amount, withdrawalFee, to)
}

// WalletLogicV1NewFreezeIterator is returned from FilterNewFreeze and is used to iterate over the raw logs and unpacked data for NewFreeze events raised by the WalletLogicV1 contract.
type WalletLogicV1NewFreezeIterator struct {
	Event *WalletLogicV1NewFreeze // Event containing the contract specifics and raw log

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
func (it *WalletLogicV1NewFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLogicV1NewFreeze)
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
		it.Event = new(WalletLogicV1NewFreeze)
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
func (it *WalletLogicV1NewFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLogicV1NewFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLogicV1NewFreeze represents a NewFreeze event raised by the WalletLogicV1 contract.
type WalletLogicV1NewFreeze struct {
	Amount     *big.Int
	NewBalance *big.Int
	Wallet     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewFreeze is a free log retrieval operation binding the contract event 0xd7c98bdf591568a7d44ba2dd863ec060c0e47aa9f636f8725a6865f3c8de24b1.
//
// Solidity: event newFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) FilterNewFreeze(opts *bind.FilterOpts) (*WalletLogicV1NewFreezeIterator, error) {

	logs, sub, err := _WalletLogicV1.contract.FilterLogs(opts, "newFreeze")
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1NewFreezeIterator{contract: _WalletLogicV1.contract, event: "newFreeze", logs: logs, sub: sub}, nil
}

// WatchNewFreeze is a free log subscription operation binding the contract event 0xd7c98bdf591568a7d44ba2dd863ec060c0e47aa9f636f8725a6865f3c8de24b1.
//
// Solidity: event newFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) WatchNewFreeze(opts *bind.WatchOpts, sink chan<- *WalletLogicV1NewFreeze) (event.Subscription, error) {

	logs, sub, err := _WalletLogicV1.contract.WatchLogs(opts, "newFreeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLogicV1NewFreeze)
				if err := _WalletLogicV1.contract.UnpackLog(event, "newFreeze", log); err != nil {
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

// ParseNewFreeze is a log parse operation binding the contract event 0xd7c98bdf591568a7d44ba2dd863ec060c0e47aa9f636f8725a6865f3c8de24b1.
//
// Solidity: event newFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) ParseNewFreeze(log types.Log) (*WalletLogicV1NewFreeze, error) {
	event := new(WalletLogicV1NewFreeze)
	if err := _WalletLogicV1.contract.UnpackLog(event, "newFreeze", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletLogicV1NewPurchaseIterator is returned from FilterNewPurchase and is used to iterate over the raw logs and unpacked data for NewPurchase events raised by the WalletLogicV1 contract.
type WalletLogicV1NewPurchaseIterator struct {
	Event *WalletLogicV1NewPurchase // Event containing the contract specifics and raw log

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
func (it *WalletLogicV1NewPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLogicV1NewPurchase)
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
		it.Event = new(WalletLogicV1NewPurchase)
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
func (it *WalletLogicV1NewPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLogicV1NewPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLogicV1NewPurchase represents a NewPurchase event raised by the WalletLogicV1 contract.
type WalletLogicV1NewPurchase struct {
	QuoteId     string
	Amount      *big.Int
	PendingFees *big.Int
	Fee         *big.Int
	Customer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewPurchase is a free log retrieval operation binding the contract event 0x404fae266883113fdb648f91507a22fd49f177aaed47f51c97349b02dacbf489.
//
// Solidity: event newPurchase(string quoteId, uint256 amount, uint256 pendingFees, uint256 fee, address customer)
func (_WalletLogicV1 *WalletLogicV1Filterer) FilterNewPurchase(opts *bind.FilterOpts) (*WalletLogicV1NewPurchaseIterator, error) {

	logs, sub, err := _WalletLogicV1.contract.FilterLogs(opts, "newPurchase")
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1NewPurchaseIterator{contract: _WalletLogicV1.contract, event: "newPurchase", logs: logs, sub: sub}, nil
}

// WatchNewPurchase is a free log subscription operation binding the contract event 0x404fae266883113fdb648f91507a22fd49f177aaed47f51c97349b02dacbf489.
//
// Solidity: event newPurchase(string quoteId, uint256 amount, uint256 pendingFees, uint256 fee, address customer)
func (_WalletLogicV1 *WalletLogicV1Filterer) WatchNewPurchase(opts *bind.WatchOpts, sink chan<- *WalletLogicV1NewPurchase) (event.Subscription, error) {

	logs, sub, err := _WalletLogicV1.contract.WatchLogs(opts, "newPurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLogicV1NewPurchase)
				if err := _WalletLogicV1.contract.UnpackLog(event, "newPurchase", log); err != nil {
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

// ParseNewPurchase is a log parse operation binding the contract event 0x404fae266883113fdb648f91507a22fd49f177aaed47f51c97349b02dacbf489.
//
// Solidity: event newPurchase(string quoteId, uint256 amount, uint256 pendingFees, uint256 fee, address customer)
func (_WalletLogicV1 *WalletLogicV1Filterer) ParseNewPurchase(log types.Log) (*WalletLogicV1NewPurchase, error) {
	event := new(WalletLogicV1NewPurchase)
	if err := _WalletLogicV1.contract.UnpackLog(event, "newPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletLogicV1NewUnFreezeIterator is returned from FilterNewUnFreeze and is used to iterate over the raw logs and unpacked data for NewUnFreeze events raised by the WalletLogicV1 contract.
type WalletLogicV1NewUnFreezeIterator struct {
	Event *WalletLogicV1NewUnFreeze // Event containing the contract specifics and raw log

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
func (it *WalletLogicV1NewUnFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLogicV1NewUnFreeze)
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
		it.Event = new(WalletLogicV1NewUnFreeze)
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
func (it *WalletLogicV1NewUnFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLogicV1NewUnFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLogicV1NewUnFreeze represents a NewUnFreeze event raised by the WalletLogicV1 contract.
type WalletLogicV1NewUnFreeze struct {
	Amount     *big.Int
	NewBalance *big.Int
	Wallet     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUnFreeze is a free log retrieval operation binding the contract event 0x49f0e6e6a9465a79ab623e1165dc01f7015b14103f87f8a80b565a2cc8d79eb7.
//
// Solidity: event newUnFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) FilterNewUnFreeze(opts *bind.FilterOpts) (*WalletLogicV1NewUnFreezeIterator, error) {

	logs, sub, err := _WalletLogicV1.contract.FilterLogs(opts, "newUnFreeze")
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1NewUnFreezeIterator{contract: _WalletLogicV1.contract, event: "newUnFreeze", logs: logs, sub: sub}, nil
}

// WatchNewUnFreeze is a free log subscription operation binding the contract event 0x49f0e6e6a9465a79ab623e1165dc01f7015b14103f87f8a80b565a2cc8d79eb7.
//
// Solidity: event newUnFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) WatchNewUnFreeze(opts *bind.WatchOpts, sink chan<- *WalletLogicV1NewUnFreeze) (event.Subscription, error) {

	logs, sub, err := _WalletLogicV1.contract.WatchLogs(opts, "newUnFreeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLogicV1NewUnFreeze)
				if err := _WalletLogicV1.contract.UnpackLog(event, "newUnFreeze", log); err != nil {
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

// ParseNewUnFreeze is a log parse operation binding the contract event 0x49f0e6e6a9465a79ab623e1165dc01f7015b14103f87f8a80b565a2cc8d79eb7.
//
// Solidity: event newUnFreeze(uint256 amount, uint256 newBalance, address wallet)
func (_WalletLogicV1 *WalletLogicV1Filterer) ParseNewUnFreeze(log types.Log) (*WalletLogicV1NewUnFreeze, error) {
	event := new(WalletLogicV1NewUnFreeze)
	if err := _WalletLogicV1.contract.UnpackLog(event, "newUnFreeze", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WalletLogicV1NewWithdrawalIterator is returned from FilterNewWithdrawal and is used to iterate over the raw logs and unpacked data for NewWithdrawal events raised by the WalletLogicV1 contract.
type WalletLogicV1NewWithdrawalIterator struct {
	Event *WalletLogicV1NewWithdrawal // Event containing the contract specifics and raw log

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
func (it *WalletLogicV1NewWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletLogicV1NewWithdrawal)
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
		it.Event = new(WalletLogicV1NewWithdrawal)
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
func (it *WalletLogicV1NewWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletLogicV1NewWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletLogicV1NewWithdrawal represents a NewWithdrawal event raised by the WalletLogicV1 contract.
type WalletLogicV1NewWithdrawal struct {
	Amount        *big.Int
	PendingFees   *big.Int
	WitpendingFee *big.Int
	To            common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawal is a free log retrieval operation binding the contract event 0x1be5c93cb35ef6c42bc250de0f3835c641d0ab3b50abed71892f8890156526c8.
//
// Solidity: event newWithdrawal(uint256 amount, uint256 pendingFees, uint256 witpendingFee, address indexed to)
func (_WalletLogicV1 *WalletLogicV1Filterer) FilterNewWithdrawal(opts *bind.FilterOpts, to []common.Address) (*WalletLogicV1NewWithdrawalIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WalletLogicV1.contract.FilterLogs(opts, "newWithdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return &WalletLogicV1NewWithdrawalIterator{contract: _WalletLogicV1.contract, event: "newWithdrawal", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawal is a free log subscription operation binding the contract event 0x1be5c93cb35ef6c42bc250de0f3835c641d0ab3b50abed71892f8890156526c8.
//
// Solidity: event newWithdrawal(uint256 amount, uint256 pendingFees, uint256 witpendingFee, address indexed to)
func (_WalletLogicV1 *WalletLogicV1Filterer) WatchNewWithdrawal(opts *bind.WatchOpts, sink chan<- *WalletLogicV1NewWithdrawal, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WalletLogicV1.contract.WatchLogs(opts, "newWithdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletLogicV1NewWithdrawal)
				if err := _WalletLogicV1.contract.UnpackLog(event, "newWithdrawal", log); err != nil {
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

// ParseNewWithdrawal is a log parse operation binding the contract event 0x1be5c93cb35ef6c42bc250de0f3835c641d0ab3b50abed71892f8890156526c8.
//
// Solidity: event newWithdrawal(uint256 amount, uint256 pendingFees, uint256 witpendingFee, address indexed to)
func (_WalletLogicV1 *WalletLogicV1Filterer) ParseNewWithdrawal(log types.Log) (*WalletLogicV1NewWithdrawal, error) {
	event := new(WalletLogicV1NewWithdrawal)
	if err := _WalletLogicV1.contract.UnpackLog(event, "newWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
