package bscvalidatorset

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

var account, _ = cmm.FromHexKey("B7F5FE087D796A24E7D1215DDF809F101B3147CB8DB0B5B869A5CD7AFDEF9B16")
var bscvalidatorsetABI, _ = abi.JSON(strings.NewReader(BscvalidatorsetABI))

func TestSimulateBSCValidatorPackage(t *testing.T) {
	client, err := ethclient.Dial(cmm.DevEndpoint)
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

	value, _ := hex.DecodeString("00")
	proofBytes, _ := hex.DecodeString("0a93040a066961766c3a76120e000001000208000000000000006b1af803f6030af3030a2b080b10840618a4ac172a20986fc24e8fbe8219e257cf30dc79c7362ddd430225b1df62128c4a19e14309540a2b080910ef021888f4102a20d9c729d658372cdf8838d66c102048bb8fbc4fb611627a427b6461ea01ed0e060a2b080810af0118888f082220ef59abeaedb6ede9ac90b7a02da04a8cb0cacbd4e545b770e19489af31d878030a2a0807106a18c8d3062a20c5f6e1a2ec2dd5789744c9d6b4030bab3177a4bb3889671c4745a9c44fce01690a2a0806103a18a8ba0422203e0e889581166e35d923b9b07aea8006abfc02d6cf85efc5938aaac058506af90a2a0805101a18b88b042a20eff4441afc6ddf83ba3d0a5e50081ca107ba1464d858e857beaa29ddc41ae77c0a2a0804100e1890c5032220cfe82506332c63be747e736ac1c51efdc05f0ca3182d0a4e6197fff255def0670a2a0803100618a8bd032a20c322baf969958c47636d52194a019a2d896671f6f84cecd28417dff82be852d00a2a0802100418c0b5032220c0d9d1850008baa94de42b360996d7e4cea052bb5c3ea8bd708e4ca2e128c30b0a2a0801100218ccb1032220983e63b6755057268d28ebf4fc90f5c90f2012bacf4752bdc0090f19c3eea4191a360a0e000001000208000000000000006b12206e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01d18e4a9030a80050a0a6d756c746973746f726512036962631aec04ea040ae7040a0e0a046d61696e12060a0408deaf170a310a057374616b6512280a2608deaf17122051b1c251b2513a61158254db921a5971c8da340e5193584f184f5fcbd19ec9060a320a066f7261636c6512280a2608deaf171220ac7fced02a58ae92b4815a68d695e9647a93e30fdb06a716bbd3e306695e197a0a130a0974696d655f6c6f636b12060a0408deaf170a320a06706172616d7312280a2608deaf17122007e752eaefa2f29d5ddcc1b2cf64ed500f4d7164cad5b24aefbde32f1eb5ce720a2f0a0361636312280a2608deaf1712203c3cb62200f55a72c434fd855fc905ec4b2087528fdd3df0e31e6573a4a0f3470a0f0a05706169727312060a0408deaf170a320a06746f6b656e7312280a2608deaf171220f5158c1c5786083fdd104a253b1767a16e7e7232d6271a9c96d3041034cea4520a100a0662726964676512060a0408deaf170a2f0a0364657812280a2608deaf171220d1ccd2a6baf2421d0a4b4a61d4b911b73245d117bb190ba0dc5f75c442b321a60a150a0b61746f6d69635f7377617012060a0408deaf170a2f0a0369626312280a2608deaf171220b2069a4eafaa724ba37cfbd2775b3339fa084aa344eb2e72d195c170b2ec19e00a2f0a03676f7612280a2608deaf17122047d9b06442e048391331c8ddefc8d655e37b2b57b2f2e4e0034e735d3a1fefe90a340a08736c617368696e6712280a2608deaf17122015fada5d33ab9fc42b454da7f95c55734a06c06d0bb6bd6c00fd231ea9bd58ab0a0d0a0376616c12060a0408deaf170a2e0a02736312280a2608deaf171220b90beab212945011a1303ea43165b5fa8e777f6210114e74e542ad79462dc14a")

	data, err := bscvalidatorsetABI.Pack("handlePackage", value, proofBytes, uint64(382943), uint64(107))
	assert.NoError(t, err)
	msg := ethereum.CallMsg{
		From:     auth.From,
		To:       &cmm.BSCValidatorSet,
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	res, err := client.CallContract(context.Background(), msg, nil)

	fmt.Println(err)
	fmt.Println(string(res))
}

