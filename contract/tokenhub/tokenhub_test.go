package tokenhub

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	cmm "github/HaoyangLiu/contract-simulation/contract/common"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

var account, _ = cmm.FromHexKey("B7F5FE087D796A24E7D1215DDF809F101B3147CB8DB0B5B869A5CD7AFDEF9B16")
var receiveAccount = common.HexToAddress("0xaa25Aa7a19f9c426E07dee59b12f944f4d9f1DD3")
var tokenhubABI, _ = abi.JSON(strings.NewReader(TokenhubABI))

func TestSimulateUpdateValidatorContract(t *testing.T) {
	client, err := ethclient.Dial(cmm.Endpoint)
	assert.NoError(t, err)

	nonce, err := client.PendingNonceAt(context.Background(), account.Addr)
	assert.NoError(t, err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.NoError(t, err)

	auth := bind.NewKeyedTransactor(account.Key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.GasLimit = 4700000

	value, _ := hex.DecodeString("000000000000000000000000000000000000000000000000030885f52d3e2c00000000000000000000000000000000000000000053a9e6a64992ffce817b7e973a8df704b5058f300002")
	proofBytes, _ := hex.DecodeString("0aaa020a066961766c3a76120e00000100020300000000000000001a8f028d020a8a020a290805101918d00f2a20ae4a5ecbadd9e97a4c646e0a5b1de8413625d4a3f25b5b6e09a072e8f019181b0a290804101018d00f2a2015a80263ea03112a825e44215281c3c5ad43fc4de72b5000e050bb0aba51ff170a2808031008180a2a208caace7fcd1e2df0a9c3ed6343b13e625d11684ec0eb6138881ad221225fdb940a280802100418062a2048cddfb0f19dc4b6f8055438d5550c90bb67c4de02c8909bff4985d50f0ce9960a280801100218042a20780567c38faac6e179a1fe6bf034e5f91f949c49cbb69c7fde945570cfa46f931a340a0e00000100020300000000000000001220d2ef17421249f45568301da10bdc3690aeb1ffc75f56239e413c480e811fa59618020af0040a0a6d756c746973746f726512036962631adc04da040ad7040a330a08736c617368696e6712270a2508f611122010674a7418aa7542e7633655d3c1c77028f42627b8d5c45aa583c13bf364e3170a310a06746f6b656e7312270a2508f61112200c952b3e3844a8e89867b18eee57afa1d0d0bfef0f460f789040be6cc4a276c70a120a0974696d655f6c6f636b12050a0308f6110a310a06706172616d7312270a2508f611122007e752eaefa2f29d5ddcc1b2cf64ed500f4d7164cad5b24aefbde32f1eb5ce720a2e0a0369626312270a2508f6111220ba16ac51c2bca14c814ac4023af029af24c8464d5de4dd2bbfd3a56c6610bb350a2d0a02736312270a2508f6111220b90beab212945011a1303ea43165b5fa8e777f6210114e74e542ad79462dc14a0a2e0a0364657812270a2508f6111220b893b27013417d8596647dc78b21e031d15f2dba842ab251e48a61c4278a146d0a0d0a046d61696e12050a0308f6110a0f0a0662726964676512050a0308f6110a2e0a0361636312270a2508f61112204e00f089ff2dded7659028a38bb096fdef10e7e96580c33e24eade5a062529b90a0e0a05706169727312050a0308f6110a0c0a0376616c12050a0308f6110a300a057374616b6512270a2508f61112204e693eccf5e03e46cb82ea6f77782a0956687f3317b0ca0b2ce90b551d6a53440a140a0b61746f6d69635f7377617012050a0308f6110a2e0a03676f7612270a2508f611122047d9b06442e048391331c8ddefc8d655e37b2b57b2f2e4e0034e735d3a1fefe90a310a066f7261636c6512270a2508f6111220f5de430b61c1f312da7292e6426159e220b950caba5c0b656aedcb2ac15ab9d3")

	data, err := tokenhubABI.Pack("handleRefundPackage", value, proofBytes, uint64(2295), uint64(0))
	assert.NoError(t, err)

	msg := ethereum.CallMsg{
		From:     auth.From,
		To:       &cmm.TokenHub,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)

	fmt.Println(err)
	fmt.Println(string(res))
}
