package common

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	Endpoint = "http://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:8545"

	BSCValidatorSet       = common.HexToAddress("0x0000000000000000000000000000000000001000")
	SlashIndicator        = common.HexToAddress("0x0000000000000000000000000000000000001001")
	SystemReward          = common.HexToAddress("0x0000000000000000000000000000000000001002")
	TendermintLightClient = common.HexToAddress("0x0000000000000000000000000000000000001003")
	TokenHub              = common.HexToAddress("0x0000000000000000000000000000000000001004")
	RelayerIncentivize    = common.HexToAddress("0x0000000000000000000000000000000000001005")
	RelayerHub            = common.HexToAddress("0x0000000000000000000000000000000000001006")
)

func SendEther(client *ethclient.Client, fromEO ExtAcc, toAddr common.Address, value *big.Int, hugegasPrice bool) (common.Hash, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromEO.Addr)
	if err != nil {
		return common.Hash{}, err
	}
	gasLimit := uint64(3e4)
	var gasPrice *big.Int
	if hugegasPrice {
		gasPrice = big.NewInt(params.Ether / 10000)
	} else {
		gasPrice, err = client.SuggestGasPrice(context.Background())
	}
	if err != nil {
		return common.Hash{}, err
	}
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return common.Hash{}, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), fromEO.Key)
	if err != nil {
		return common.Hash{}, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return common.Hash{}, err
	}
	txhash := signedTx.Hash()
	log.Printf("tx hash, sendEther: %s\n", txhash.Hex())
	return txhash, nil
}

type ExtAcc struct {
	Key  *ecdsa.PrivateKey
	Addr common.Address
}

func FromHexKey(hexkey string) (ExtAcc, error) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return ExtAcc{}, err
	}
	pubKey := key.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("publicKey is not of type *ecdsa.PublicKey")
		return ExtAcc{}, err
	}
	addr := crypto.PubkeyToAddress(*pubKeyECDSA)
	return ExtAcc{key, addr}, nil
}
