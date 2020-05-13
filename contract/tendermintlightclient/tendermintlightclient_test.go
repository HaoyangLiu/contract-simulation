package tendermintlightclient

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	cmm "github/HaoyangLiu/contract-simulation/contract/common"
)

var account, _ = cmm.FromHexKey("E10974B275F6AE7CA64D3469438B3AF31E4EAC2BDDD6B2476438F7E287A2E375")
var tendermintlightclientABI, _ = abi.JSON(strings.NewReader(TendermintlightclientABI))

func TestSimulateSyncTendermintHeader(t *testing.T) {
	client, err := ethclient.Dial(cmm.TestnetEndpoint)
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

	headerBytes, _ := hex.DecodeString("85080a84030ab5020a02080a121342696e616e63652d436861696e2d4b6f6e676f18f7d80f220c08ad82eaf50510fff9dbe2022a480a200a2548ec462885cc12ce872076c675b5661754a9d49d4c3735b8c7c80c0103b01224080112209468153baca743d8473e34d00419c2e1d8ecd055b0cae304c7ea352142a585ac32209302e1110865dffd1f7c95aa6ebff662e45b3e7d105bb97b3406aae6dfce0fdc4220aea1ac326886b992a991d21a6eb155f41b77867cbf659e78f31d89d8205122a84a20aea1ac326886b992a991d21a6eb155f41b77867cbf659e78f31d89d8205122a85220294d8fbd0b94b767a7eba9840f299a3586da7fe6b5dead3b7eecba193c400f935a20267e70700ed59daa8b352ead6ffa788e1b21ae3b7341f114123e5fbe4414a96c7214f42f1d05ac568d12e26b9655395e7fbbd46bc5bb124a1a480a20f6b7938000ff78665c07521c0159992f3205faa5056fb6f537203271e317a9d51224080112206d3d6b94ed0099cb4bfc864d292eae0b6c428889f032db275bf0816aa689f60712bc020a4f0a149ccddd479c0ad8dcd01d754dd95fe15384e8bbdc12251624de64204d1be64f0e9a466c2e66a53433928192783e29f8fa21beb2133499b5ef770f601880a094a58d1d2080c0d7b5e5c5ffffff010a4b0a14d4cecef238e778c7063552c3a7ab95da35c3fb4712251624de642099308aa365c40554bc89982af505d85da95251445d5dd4a9bb37dd2584fd92d31880a094a58d1d2080a094a58d1d0a4b0a14f42f1d05ac568d12e26b9655395e7fbbd46bc5bb12251624de642001776920ff0b0f38d78cf95c033c21adf7045785114e392a7544179652e0a6121880a094a58d1d2080a094a58d1d124f0a149ccddd479c0ad8dcd01d754dd95fe15384e8bbdc12251624de64204d1be64f0e9a466c2e66a53433928192783e29f8fa21beb2133499b5ef770f601880a094a58d1d2080c0d7b5e5c5ffffff011abc020a4f0a149ccddd479c0ad8dcd01d754dd95fe15384e8bbdc12251624de64204d1be64f0e9a466c2e66a53433928192783e29f8fa21beb2133499b5ef770f601880a094a58d1d2080c0d7b5e5c5ffffff010a4b0a14d4cecef238e778c7063552c3a7ab95da35c3fb4712251624de642099308aa365c40554bc89982af505d85da95251445d5dd4a9bb37dd2584fd92d31880a094a58d1d2080a094a58d1d0a4b0a14f42f1d05ac568d12e26b9655395e7fbbd46bc5bb12251624de642001776920ff0b0f38d78cf95c033c21adf7045785114e392a7544179652e0a6121880a094a58d1d2080a094a58d1d124f0a149ccddd479c0ad8dcd01d754dd95fe15384e8bbdc12251624de64204d1be64f0e9a466c2e66a53433928192783e29f8fa21beb2133499b5ef770f601880a094a58d1d2080c0d7b5e5c5ffffff01,")

	data, err := tendermintlightclientABI.Pack("syncTendermintHeader", headerBytes, uint64(2))
	assert.NoError(t, err)

	msg := ethereum.CallMsg{
		From:     auth.From,
		To:       &cmm.TendermintLightClient,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)

	fmt.Println(err)
	fmt.Println(string(res))
}
