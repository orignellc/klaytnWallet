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

// GlobalP2PABI is the input ABI used to generate the binding from.
const GlobalP2PABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"newWallet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"deployWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletLogic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usd\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletLogic\",\"type\":\"address\"}],\"name\":\"updateWalletLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCUSDBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GlobalP2PBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const GlobalP2PBinRuntime = ``

// GlobalP2PBin is the compiled bytecode used for deploying new contracts.
var GlobalP2PBin = "0x608060405234801561001057600080fd5b50611f94806100206000396000f3fe608060405260043610620001035760003560e01c806372f179f81162000097578063beacf2cb1162000061578063beacf2cb1462000314578063da76d5cd1462000344578063f2fde38b1462000372578063f897a22b14620003a0576200010b565b806372f179f8146200025857806382787c7f146200029c5780638456cb5914620002ca5780638da5cb5b14620002e4576200010b565b80633f4ba83a11620000d95780633f4ba83a14620001c6578063485cc95514620001e05780635c975abb146200020e578063715018a6146200023e576200010b565b8063038defd714620001105780632ab66623146200015457806337ba532c1462000182576200010b565b366200010b57005b600080fd5b3480156200011d57600080fd5b506200013c6004803603810190620001369190620010a6565b620003d0565b6040516200014b9190620010f5565b60405180910390f35b3480156200016157600080fd5b506200018060048036038101906200017a91906200114d565b620003f0565b005b3480156200018f57600080fd5b50620001ae6004803603810190620001a89190620012e1565b620004ac565b604051620001bd919062001343565b60405180910390f35b348015620001d357600080fd5b50620001de6200072d565b005b348015620001ed57600080fd5b506200020c600480360381019062000206919062001360565b62000743565b005b3480156200021b57600080fd5b50620002266200091d565b604051620002359190620010f5565b60405180910390f35b3480156200024b57600080fd5b506200025662000934565b005b3480156200026557600080fd5b506200028460048036038101906200027e9190620012e1565b6200094c565b60405162000293919062001343565b60405180910390f35b348015620002a957600080fd5b50620002c86004803603810190620002c29190620010a6565b62000995565b005b348015620002d757600080fd5b50620002e2620009ed565b005b348015620002f157600080fd5b50620002fc62000a03565b6040516200030b919062001343565b60405180910390f35b3480156200032157600080fd5b506200032c62000a2d565b6040516200033b919062001343565b60405180910390f35b3480156200035157600080fd5b506200037060048036038101906200036a91906200114d565b62000a53565b005b3480156200037f57600080fd5b506200039e6004803603810190620003989190620010a6565b62000af7565b005b348015620003ad57600080fd5b50620003b862000b81565b604051620003c7919062001343565b60405180910390f35b609a6020528060005260406000206000915054906101000a900460ff1681565b620003fa62000ba7565b609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6200044262000a03565b836040518363ffffffff1660e01b815260040162000462929190620013b8565b6020604051808303816000875af115801562000482573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620004a8919062001416565b5050565b6000620004b862000ba7565b620004c262000c2c565b6020825110156200050a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200050190620014a9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166099836040516200053491906200154e565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614620005bc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620005b390620015b7565b60405180910390fd5b600030609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620005f0906200101f565b620005fd929190620015d9565b604051809103906000f0801580156200061a573d6000803e3d6000fd5b509050806099846040516200063091906200154e565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001609a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b846040516200071c919062001647565b60405180910390a280915050919050565b6200073762000ba7565b6200074162000c7b565b565b60008060019054906101000a900460ff16159050808015620007755750600160008054906101000a900460ff1660ff16105b80620007a65750620007873062000ce4565b158015620007a55750600160008054906101000a900460ff1660ff16145b5b620007e8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007df90620016e1565b60405180910390fd5b60016000806101000a81548160ff021916908360ff160217905550801562000826576001600060016101000a81548160ff0219169083151502179055505b6200083062000d07565b6200083a62000d65565b82609760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081609860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015620009185760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516200090f91906200175d565b60405180910390a15b505050565b6000606560009054906101000a900460ff16905090565b6200093e62000ba7565b6200094a600062000dc3565b565b6099818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6200099f62000ba7565b620009a962000e89565b80609760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b620009f762000ba7565b62000a0162000ed7565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b609760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b62000a5d62000ba7565b4781111562000aa3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a9a90620017ca565b60405180910390fd5b62000aad62000a03565b73ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801562000af3573d6000803e3d6000fd5b5050565b62000b0162000ba7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160362000b73576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000b6a9062001862565b60405180910390fd5b62000b7e8162000dc3565b50565b609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b62000bb162000f40565b73ffffffffffffffffffffffffffffffffffffffff1662000bd162000a03565b73ffffffffffffffffffffffffffffffffffffffff161462000c2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000c2190620018d4565b60405180910390fd5b565b62000c366200091d565b1562000c79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000c709062001946565b60405180910390fd5b565b62000c8562000e89565b6000606560006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa62000ccb62000f40565b60405162000cda919062001343565b60405180910390a1565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1662000d59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000d5090620019de565b60405180910390fd5b62000d6362000f48565b565b600060019054906101000a900460ff1662000db7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000dae90620019de565b60405180910390fd5b62000dc162000fb0565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b62000e936200091d565b62000ed5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000ecc9062001a50565b60405180910390fd5b565b62000ee162000c2c565b6001606560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25862000f2762000f40565b60405162000f36919062001343565b60405180910390a1565b600033905090565b600060019054906101000a900460ff1662000f9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000f9190620019de565b60405180910390fd5b62000fae62000fa862000f40565b62000dc3565b565b600060019054906101000a900460ff1662001002576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000ff990620019de565b60405180910390fd5b6000606560006101000a81548160ff021916908315150217905550565b6104ec8062001a7383390190565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200106e8262001041565b9050919050565b620010808162001061565b81146200108c57600080fd5b50565b600081359050620010a08162001075565b92915050565b600060208284031215620010bf57620010be62001037565b5b6000620010cf848285016200108f565b91505092915050565b60008115159050919050565b620010ef81620010d8565b82525050565b60006020820190506200110c6000830184620010e4565b92915050565b6000819050919050565b620011278162001112565b81146200113357600080fd5b50565b60008135905062001147816200111c565b92915050565b60006020828403121562001166576200116562001037565b5b6000620011768482850162001136565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620011d48262001189565b810181811067ffffffffffffffff82111715620011f657620011f56200119a565b5b80604052505050565b60006200120b6200102d565b9050620012198282620011c9565b919050565b600067ffffffffffffffff8211156200123c576200123b6200119a565b5b620012478262001189565b9050602081019050919050565b82818337600083830152505050565b60006200127a62001274846200121e565b620011ff565b90508281526020810184848401111562001299576200129862001184565b5b620012a684828562001254565b509392505050565b600082601f830112620012c657620012c56200117f565b5b8135620012d884826020860162001263565b91505092915050565b600060208284031215620012fa57620012f962001037565b5b600082013567ffffffffffffffff8111156200131b576200131a6200103c565b5b6200132984828501620012ae565b91505092915050565b6200133d8162001061565b82525050565b60006020820190506200135a600083018462001332565b92915050565b600080604083850312156200137a576200137962001037565b5b60006200138a858286016200108f565b92505060206200139d858286016200108f565b9150509250929050565b620013b28162001112565b82525050565b6000604082019050620013cf600083018562001332565b620013de6020830184620013a7565b9392505050565b620013f081620010d8565b8114620013fc57600080fd5b50565b6000815190506200141081620013e5565b92915050565b6000602082840312156200142f576200142e62001037565b5b60006200143f84828501620013ff565b91505092915050565b600082825260208201905092915050565b7f47503a496e76616c696420555549440000000000000000000000000000000000600082015250565b600062001491600f8362001448565b91506200149e8262001459565b602082019050919050565b60006020820190508181036000830152620014c48162001482565b9050919050565b600081519050919050565b600081905092915050565b60005b8381101562001501578082015181840152602081019050620014e4565b8381111562001511576000848401525b50505050565b60006200152482620014cb565b620015308185620014d6565b935062001542818560208601620014e1565b80840191505092915050565b60006200155c828462001517565b915081905092915050565b7f47503a5573657220657869737400000000000000000000000000000000000000600082015250565b60006200159f600d8362001448565b9150620015ac8262001567565b602082019050919050565b60006020820190508181036000830152620015d28162001590565b9050919050565b6000604082019050620015f0600083018562001332565b620015ff602083018462001332565b9392505050565b60006200161382620014cb565b6200161f818562001448565b935062001631818560208601620014e1565b6200163c8162001189565b840191505092915050565b6000602082019050818103600083015262001663818462001606565b905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000620016c9602e8362001448565b9150620016d6826200166b565b604082019050919050565b60006020820190508181036000830152620016fc81620016ba565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620017456200173f620017398462001703565b6200171a565b6200170d565b9050919050565b620017578162001724565b82525050565b60006020820190506200177460008301846200174c565b92915050565b7f47503a20496e73756666696369656e742062616c616e63650000000000000000600082015250565b6000620017b260188362001448565b9150620017bf826200177a565b602082019050919050565b60006020820190508181036000830152620017e581620017a3565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006200184a60268362001448565b91506200185782620017ec565b604082019050919050565b600060208201905081810360008301526200187d816200183b565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000620018bc60208362001448565b9150620018c98262001884565b602082019050919050565b60006020820190508181036000830152620018ef81620018ad565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b60006200192e60108362001448565b91506200193b82620018f6565b602082019050919050565b6000602082019050818103600083015262001961816200191f565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6000620019c6602b8362001448565b9150620019d38262001968565b604082019050919050565b60006020820190508181036000830152620019f981620019b7565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b600062001a3860148362001448565b915062001a458262001a00565b602082019050919050565b6000602082019050818103600083015262001a6b8162001a29565b905091905056fe608060405234801561001057600080fd5b506040516104ec3803806104ec8339818101604052810190610032919061011d565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ea826100bf565b9050919050565b6100fa816100df565b811461010557600080fd5b50565b600081519050610117816100f1565b92915050565b60008060408385031215610134576101336100ba565b5b600061014285828601610108565b925050602061015385828601610108565b9150509250929050565b6103808061016c6000396000f3fe60806040526004361061002d5760003560e01c8063a890bd15146100a3578063c55dae63146100ce57610099565b36610099573073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f476c794848a86924c99aa79862df9075793bc6e48fa9e77f85f196f12731b0bd3460405161008f9190610239565b60405180910390a3005b6100a16100f9565b005b3480156100af57600080fd5b506100b8610113565b6040516100c59190610295565b60405180910390f35b3480156100da57600080fd5b506100e3610137565b6040516100f091906102d1565b60405180910390f35b61010161015d565b61011161010c61015f565b6101fa565b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663beacf2cb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101f4919061031d565b91505090565b3660008037600080366000845af43d6000803e806000811461021b573d6000f35b3d6000fd5b6000819050919050565b61023381610220565b82525050565b600060208201905061024e600083018461022a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061027f82610254565b9050919050565b61028f81610274565b82525050565b60006020820190506102aa6000830184610286565b92915050565b60006102bb82610254565b9050919050565b6102cb816102b0565b82525050565b60006020820190506102e660008301846102c2565b92915050565b600080fd5b6102fa816102b0565b811461030557600080fd5b50565b600081519050610317816102f1565b92915050565b600060208284031215610333576103326102ec565b5b600061034184828501610308565b9150509291505056fea26469706673582212203c45c434c0b02c2b223746b228816e599c63d4b1c6e5519a610c1bbb23f83ac064736f6c634300080f0033a2646970667358221220c2f22d222fa14e331dc4cf839d1563ccdd9a1e158a977ce4d7bdf9d9ffbbb5b664736f6c634300080f0033"

// DeployGlobalP2P deploys a new Klaytn contract, binding an instance of GlobalP2P to it.
func DeployGlobalP2P(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalP2P, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalP2PABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalP2PBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalP2P{GlobalP2PCaller: GlobalP2PCaller{contract: contract}, GlobalP2PTransactor: GlobalP2PTransactor{contract: contract}, GlobalP2PFilterer: GlobalP2PFilterer{contract: contract}}, nil
}

// GlobalP2P is an auto generated Go binding around a Klaytn contract.
type GlobalP2P struct {
	GlobalP2PCaller     // Read-only binding to the contract
	GlobalP2PTransactor // Write-only binding to the contract
	GlobalP2PFilterer   // Log filterer for contract events
}

// GlobalP2PCaller is an auto generated read-only Go binding around a Klaytn contract.
type GlobalP2PCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PTransactor is an auto generated write-only Go binding around a Klaytn contract.
type GlobalP2PTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type GlobalP2PFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type GlobalP2PSession struct {
	Contract     *GlobalP2P        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalP2PCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type GlobalP2PCallerSession struct {
	Contract *GlobalP2PCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GlobalP2PTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type GlobalP2PTransactorSession struct {
	Contract     *GlobalP2PTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GlobalP2PRaw is an auto generated low-level Go binding around a Klaytn contract.
type GlobalP2PRaw struct {
	Contract *GlobalP2P // Generic contract binding to access the raw methods on
}

// GlobalP2PCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type GlobalP2PCallerRaw struct {
	Contract *GlobalP2PCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalP2PTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type GlobalP2PTransactorRaw struct {
	Contract *GlobalP2PTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalP2P creates a new instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2P(address common.Address, backend bind.ContractBackend) (*GlobalP2P, error) {
	contract, err := bindGlobalP2P(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalP2P{GlobalP2PCaller: GlobalP2PCaller{contract: contract}, GlobalP2PTransactor: GlobalP2PTransactor{contract: contract}, GlobalP2PFilterer: GlobalP2PFilterer{contract: contract}}, nil
}

// NewGlobalP2PCaller creates a new read-only instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PCaller(address common.Address, caller bind.ContractCaller) (*GlobalP2PCaller, error) {
	contract, err := bindGlobalP2P(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PCaller{contract: contract}, nil
}

// NewGlobalP2PTransactor creates a new write-only instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalP2PTransactor, error) {
	contract, err := bindGlobalP2P(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PTransactor{contract: contract}, nil
}

// NewGlobalP2PFilterer creates a new log filterer instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalP2PFilterer, error) {
	contract, err := bindGlobalP2P(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PFilterer{contract: contract}, nil
}

// bindGlobalP2P binds a generic wrapper to an already deployed contract.
func bindGlobalP2P(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalP2PABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGlobalP2PABI parses the ABI
func ParseGlobalP2PABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalP2PABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalP2P *GlobalP2PRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalP2P.Contract.GlobalP2PCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalP2P *GlobalP2PRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.Contract.GlobalP2PTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalP2P *GlobalP2PRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalP2P.Contract.GlobalP2PTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalP2P *GlobalP2PCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalP2P.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalP2P *GlobalP2PTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalP2P *GlobalP2PTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalP2P.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PSession) Owner() (common.Address, error) {
	return _GlobalP2P.Contract.Owner(&_GlobalP2P.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) Owner() (common.Address, error) {
	return _GlobalP2P.Contract.Owner(&_GlobalP2P.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PSession) Paused() (bool, error) {
	return _GlobalP2P.Contract.Paused(&_GlobalP2P.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PCallerSession) Paused() (bool, error) {
	return _GlobalP2P.Contract.Paused(&_GlobalP2P.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PCaller) Registry(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "registry", arg0)
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PSession) Registry(arg0 common.Address) (bool, error) {
	return _GlobalP2P.Contract.Registry(&_GlobalP2P.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PCallerSession) Registry(arg0 common.Address) (bool, error) {
	return _GlobalP2P.Contract.Registry(&_GlobalP2P.CallOpts, arg0)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) UsdToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "usdToken")
	return *ret0, err
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PSession) UsdToken() (common.Address, error) {
	return _GlobalP2P.Contract.UsdToken(&_GlobalP2P.CallOpts)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) UsdToken() (common.Address, error) {
	return _GlobalP2P.Contract.UsdToken(&_GlobalP2P.CallOpts)
}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) WalletLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "walletLogic")
	return *ret0, err
}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PSession) WalletLogic() (common.Address, error) {
	return _GlobalP2P.Contract.WalletLogic(&_GlobalP2P.CallOpts)
}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) WalletLogic() (common.Address, error) {
	return _GlobalP2P.Contract.WalletLogic(&_GlobalP2P.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PCaller) Wallets(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GlobalP2P.contract.Call(opts, out, "wallets", arg0)
	return *ret0, err
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PSession) Wallets(arg0 string) (common.Address, error) {
	return _GlobalP2P.Contract.Wallets(&_GlobalP2P.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) Wallets(arg0 string) (common.Address, error) {
	return _GlobalP2P.Contract.Wallets(&_GlobalP2P.CallOpts, arg0)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PTransactor) DeployWallet(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "deployWallet", uuid)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PSession) DeployWallet(uuid string) (*types.Transaction, error) {
	return _GlobalP2P.Contract.DeployWallet(&_GlobalP2P.TransactOpts, uuid)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PTransactorSession) DeployWallet(uuid string) (*types.Transaction, error) {
	return _GlobalP2P.Contract.DeployWallet(&_GlobalP2P.TransactOpts, uuid)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PTransactor) Initialize(opts *bind.TransactOpts, _walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "initialize", _walletLogic, _usd)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PSession) Initialize(_walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.Initialize(&_GlobalP2P.TransactOpts, _walletLogic, _usd)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Initialize(_walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.Initialize(&_GlobalP2P.TransactOpts, _walletLogic, _usd)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GlobalP2P *GlobalP2PTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GlobalP2P *GlobalP2PSession) Pause() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Pause(&_GlobalP2P.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Pause() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Pause(&_GlobalP2P.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PSession) RenounceOwnership() (*types.Transaction, error) {
	return _GlobalP2P.Contract.RenounceOwnership(&_GlobalP2P.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GlobalP2P.Contract.RenounceOwnership(&_GlobalP2P.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.TransferOwnership(&_GlobalP2P.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.TransferOwnership(&_GlobalP2P.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GlobalP2P *GlobalP2PTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GlobalP2P *GlobalP2PSession) Unpause() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Unpause(&_GlobalP2P.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Unpause() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Unpause(&_GlobalP2P.TransactOpts)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PTransactor) UpdateWalletLogic(opts *bind.TransactOpts, _walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "updateWalletLogic", _walletLogic)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PSession) UpdateWalletLogic(_walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.UpdateWalletLogic(&_GlobalP2P.TransactOpts, _walletLogic)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) UpdateWalletLogic(_walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.UpdateWalletLogic(&_GlobalP2P.TransactOpts, _walletLogic)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactor) WithdrawBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "withdrawBalance", amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawBalance(&_GlobalP2P.TransactOpts, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawBalance(&_GlobalP2P.TransactOpts, amount)
}

// WithdrawCUSDBalance is a paid mutator transaction binding the contract method 0x2ab66623.
//
// Solidity: function withdrawCUSDBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactor) WithdrawCUSDBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "withdrawCUSDBalance", amount)
}

// WithdrawCUSDBalance is a paid mutator transaction binding the contract method 0x2ab66623.
//
// Solidity: function withdrawCUSDBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PSession) WithdrawCUSDBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawCUSDBalance(&_GlobalP2P.TransactOpts, amount)
}

// WithdrawCUSDBalance is a paid mutator transaction binding the contract method 0x2ab66623.
//
// Solidity: function withdrawCUSDBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) WithdrawCUSDBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawCUSDBalance(&_GlobalP2P.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PSession) Receive() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Receive(&_GlobalP2P.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Receive() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Receive(&_GlobalP2P.TransactOpts)
}

// GlobalP2PInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GlobalP2P contract.
type GlobalP2PInitializedIterator struct {
	Event *GlobalP2PInitialized // Event containing the contract specifics and raw log

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
func (it *GlobalP2PInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PInitialized)
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
		it.Event = new(GlobalP2PInitialized)
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
func (it *GlobalP2PInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PInitialized represents a Initialized event raised by the GlobalP2P contract.
type GlobalP2PInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) FilterInitialized(opts *bind.FilterOpts) (*GlobalP2PInitializedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PInitializedIterator{contract: _GlobalP2P.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GlobalP2PInitialized) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PInitialized)
				if err := _GlobalP2P.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) ParseInitialized(log types.Log) (*GlobalP2PInitialized, error) {
	event := new(GlobalP2PInitialized)
	if err := _GlobalP2P.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalP2POwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GlobalP2P contract.
type GlobalP2POwnershipTransferredIterator struct {
	Event *GlobalP2POwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GlobalP2POwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2POwnershipTransferred)
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
		it.Event = new(GlobalP2POwnershipTransferred)
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
func (it *GlobalP2POwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2POwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2POwnershipTransferred represents a OwnershipTransferred event raised by the GlobalP2P contract.
type GlobalP2POwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GlobalP2P *GlobalP2PFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GlobalP2POwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GlobalP2POwnershipTransferredIterator{contract: _GlobalP2P.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GlobalP2P *GlobalP2PFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GlobalP2POwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2POwnershipTransferred)
				if err := _GlobalP2P.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GlobalP2P *GlobalP2PFilterer) ParseOwnershipTransferred(log types.Log) (*GlobalP2POwnershipTransferred, error) {
	event := new(GlobalP2POwnershipTransferred)
	if err := _GlobalP2P.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalP2PPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GlobalP2P contract.
type GlobalP2PPausedIterator struct {
	Event *GlobalP2PPaused // Event containing the contract specifics and raw log

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
func (it *GlobalP2PPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PPaused)
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
		it.Event = new(GlobalP2PPaused)
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
func (it *GlobalP2PPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PPaused represents a Paused event raised by the GlobalP2P contract.
type GlobalP2PPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) FilterPaused(opts *bind.FilterOpts) (*GlobalP2PPausedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PPausedIterator{contract: _GlobalP2P.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GlobalP2PPaused) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PPaused)
				if err := _GlobalP2P.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) ParsePaused(log types.Log) (*GlobalP2PPaused, error) {
	event := new(GlobalP2PPaused)
	if err := _GlobalP2P.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalP2PUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GlobalP2P contract.
type GlobalP2PUnpausedIterator struct {
	Event *GlobalP2PUnpaused // Event containing the contract specifics and raw log

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
func (it *GlobalP2PUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PUnpaused)
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
		it.Event = new(GlobalP2PUnpaused)
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
func (it *GlobalP2PUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PUnpaused represents a Unpaused event raised by the GlobalP2P contract.
type GlobalP2PUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GlobalP2PUnpausedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PUnpausedIterator{contract: _GlobalP2P.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GlobalP2PUnpaused) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PUnpaused)
				if err := _GlobalP2P.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) ParseUnpaused(log types.Log) (*GlobalP2PUnpaused, error) {
	event := new(GlobalP2PUnpaused)
	if err := _GlobalP2P.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalP2PNewWalletIterator is returned from FilterNewWallet and is used to iterate over the raw logs and unpacked data for NewWallet events raised by the GlobalP2P contract.
type GlobalP2PNewWalletIterator struct {
	Event *GlobalP2PNewWallet // Event containing the contract specifics and raw log

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
func (it *GlobalP2PNewWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PNewWallet)
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
		it.Event = new(GlobalP2PNewWallet)
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
func (it *GlobalP2PNewWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PNewWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PNewWallet represents a NewWallet event raised by the GlobalP2P contract.
type GlobalP2PNewWallet struct {
	Uuid   string
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewWallet is a free log retrieval operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) FilterNewWallet(opts *bind.FilterOpts, wallet []common.Address) (*GlobalP2PNewWalletIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "newWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PNewWalletIterator{contract: _GlobalP2P.contract, event: "newWallet", logs: logs, sub: sub}, nil
}

// WatchNewWallet is a free log subscription operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) WatchNewWallet(opts *bind.WatchOpts, sink chan<- *GlobalP2PNewWallet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "newWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PNewWallet)
				if err := _GlobalP2P.contract.UnpackLog(event, "newWallet", log); err != nil {
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

// ParseNewWallet is a log parse operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) ParseNewWallet(log types.Log) (*GlobalP2PNewWallet, error) {
	event := new(GlobalP2PNewWallet)
	if err := _GlobalP2P.contract.UnpackLog(event, "newWallet", log); err != nil {
		return nil, err
	}
	return event, nil
}
