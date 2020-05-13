package common

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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


func GetCallOpts() (*bind.CallOpts, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	return callOpts, nil
}
