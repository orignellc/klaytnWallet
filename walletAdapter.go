package klaytnWallet

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	crypto2 "github.com/klaytn/klaytn/crypto"
	"github.com/orignellc/klaytnWallet/contracts"
	"log"
	"math/big"
)

var NilAddress = common.Address{}

type WalletOptions struct {
	RpcUrl           string
	Address          common.Address
	SignerPrivateKey string
}

type WalletAdapter struct {
	client   *client.Client
	contract *contracts.GlobalP2P
	abi      *abi.ABI
	address  common.Address
	options  *WalletOptions
}

func NewWalletAdapterWithClient(options *WalletOptions) *WalletAdapter {
	connection, err := ConnectKlaytonClient(options.RpcUrl)
	if err != nil {
		panic(err)
	}
	globalP2P, err := contracts.NewGlobalP2P(options.Address, connection)
	if err != nil {
		panic(err)
	}
	gp2pAbi, err := contracts.ParseGlobalP2PABI()
	if err != nil {
		panic(err)
	}

	return NewWalletAdapter(connection, gp2pAbi, globalP2P, options)
}

func NewWalletAdapter(client *client.Client, abi *abi.ABI, gp2p *contracts.GlobalP2P, options *WalletOptions) *WalletAdapter {
	return &WalletAdapter{client: client, abi: abi, contract: gp2p, options: options}
}

func (g *WalletAdapter) CreateWallet(userID string) (string, error) {
	opt, _ := g.transactOpt(big.NewInt(0))

	//var params []interface{}
	//params = append(params, userID)

	// Todo: remove default gas limit
	//opt.GasLimit = g.estimateGasUsage("deployWallet", params)
	opt.GasLimit = 800000
	_, err := g.contract.DeployWallet(opt, userID)
	if err != nil {
		return NilAddress.String(), err
	}

	address, err := g.GetWalletAddressFor(userID)
	if err != nil {
		return NilAddress.String(), err
	}

	return address, nil
}

func (g *WalletAdapter) GetWalletAddressFor(userId string) (string, error) {
	opts := &bind.CallOpts{Context: context.Background()}

	address, err := g.contract.Wallets(opts, userId)
	if err != nil {
		return NilAddress.String(), err
	}

	return address.String(), nil
}

func (g *WalletAdapter) transactOpt(value *big.Int) (*bind.TransactOpts, error) {
	signer, _, publicKeyECDSA, err := g.getSigner()
	if err != nil {
		return nil, err
	}

	fromAddress := crypto2.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := g.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := g.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	//networkID, err := g.client.NetworkID(context.Background())

	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(signer)

	if err != nil {
		return nil, err
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = value             // in wei
	auth.GasLimit = uint64(800000) // set default Gas Limit to be overridden by call to GasEstimate eth_estimate
	auth.GasPrice = gasPrice

	return auth, nil
}

func (g *WalletAdapter) getSigner() (*ecdsa.PrivateKey, crypto.PublicKey, *ecdsa.PublicKey, error) {
	privateKey, err := crypto2.HexToECDSA(g.options.SignerPrivateKey)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		panic(errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey"))
	}

	return privateKey, publicKey, publicKeyECDSA, err
}

func (g *WalletAdapter) estimateGasUsage(methodName string, methodArgs []interface{}) uint64 {

	var interfaceArgs []interface{}
	interfaceArgs = append(interfaceArgs, methodArgs...)

	var (
		data []byte
		err  error
	)

	if len(interfaceArgs) > 0 {
		data, err = g.abi.Pack(methodName, interfaceArgs...)
	} else {
		data, err = g.abi.Pack(methodName)
	}

	if err != nil {
		log.Fatalln("Cannot run abi pack: ", err)
	}

	_, _, publicKeyECDSA, _ := g.getSigner()

	from := crypto2.PubkeyToAddress(*publicKeyECDSA)

	callMsg := klaytn.CallMsg{
		From:     from,
		To:       &g.options.Address,
		Gas:      0,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     data,
	}
	gasLimit, err := g.client.EstimateGas(context.Background(), callMsg)
	fmt.Println("Max Transaction Gas: ", gasLimit)
	if err != nil {
		log.Fatalln("Could not estimate gas: ", err)
	}

	return gasLimit
}

func ConnectKlaytonClient(rpcUrl string) (*client.Client, error) {
	connection, err := client.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
