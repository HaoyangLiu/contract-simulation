package relayerhub

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	cmm "github/HaoyangLiu/contract-simulation/contract/common"
)

var account, _ = cmm.FromHexKey("B7F5FE087D796A24E7D1215DDF809F101B3147CB8DB0B5B869A5CD7AFDEF9B16")
var receiveAccount = common.HexToAddress("0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3")
var relayerhubABI, _ = abi.JSON(strings.NewReader(RelayerhubABI))

func TestSimulateRegister(t *testing.T) {
	client, err := ethclient.Dial(cmm.TestnetEndpoint)
	assert.NoError(t, err)

	nonce, err := client.PendingNonceAt(context.Background(), account.Addr)
	assert.NoError(t, err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)

	auth := bind.NewKeyedTransactor(account.Key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1).Mul(big.NewInt(100), big.NewInt(1000000000000000000)) //100 BNB
	auth.GasPrice = gasPrice
	auth.GasLimit = 4700000

	data, err := relayerhubABI.Pack("register")
	assert.NoError(t, err)

	msg := ethereum.CallMsg{
		From:     auth.From,
		To:       &cmm.RelayerHub,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)

	fmt.Println(err)
	fmt.Println(string(res))
}
