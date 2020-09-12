package tokenmanager

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

var account, _ = cmm.FromHexKey("8C25531D1634908C8D45B5A8A18753E72717CA95D17EDA10EFF0FD095D4C4FCB")
//var account, _ = cmm.FromHexKey("B7F5FE087D796A24E7D1215DDF809F101B3147CB8DB0B5B869A5CD7AFDEF9B16")
var receiveAccount = common.HexToAddress("0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3")
var tokenmanagerABI, _ = abi.JSON(strings.NewReader(TokenmanagerABI))


func TestSimulateRejectBind(t *testing.T) {
	client, err := ethclient.Dial(cmm.MainnetEndpoint)
	assert.NoError(t, err)

	nonce, err := client.PendingNonceAt(context.Background(), account.Addr)
	assert.NoError(t, err)

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//assert.NoError(t, err)

	auth := bind.NewKeyedTransactor(account.Key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(1e16)
	auth.GasPrice = big.NewInt(20*10^9)
	auth.GasLimit = 4700000

	contractAddr := common.HexToAddress("0xcD7C5025753a49f1881B31C48caA7C517Bb46308")
	bep2Symbol := "RAVEN-F66"
	data, err := tokenmanagerABI.Pack("approveBind", contractAddr, bep2Symbol)
	//data, err := hex.DecodeString("6b3f1307000000000000000000000000c92595b8a12290b50b421c7eb414c6ed61068b19000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000075457542d38433200000000000000000000000000000000000000000000000000")
	//assert.NoError(t, err)

	//data, _ := hex.DecodeString("095ea7b30000000000000000000000000000000000000000000000000000000000001004000000000000000000000000000000000000000000295be96e64066972000000")
	msg := ethereum.CallMsg{
		From:     common.HexToAddress("0xae80597fa2c3296378a4472e364d104ca0f47617"),
		To:       &cmm.TokenManager,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)

	fmt.Println(err)
	fmt.Println(string(res))
}

func TestSimulateGasPriced(t *testing.T) {
	client, err := ethclient.Dial(cmm.MainnetEndpoint)
	assert.NoError(t, err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)

	t.Log(gasPrice)
}