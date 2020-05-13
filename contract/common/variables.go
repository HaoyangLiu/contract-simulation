package common

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	TestnetEndpoint = "https://data-seed-prebsc-1-s1.binance.org:8545"
	DevEndpoint= "http://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:8545"
)

var (

	BSCValidatorSet       = common.HexToAddress("0x0000000000000000000000000000000000001000")
	SlashIndicator        = common.HexToAddress("0x0000000000000000000000000000000000001001")
	SystemReward          = common.HexToAddress("0x0000000000000000000000000000000000001002")
	TendermintLightClient = common.HexToAddress("0x0000000000000000000000000000000000001003")
	TokenHub              = common.HexToAddress("0x0000000000000000000000000000000000001004")
	RelayerIncentivize    = common.HexToAddress("0x0000000000000000000000000000000000001005")
	RelayerHub            = common.HexToAddress("0x0000000000000000000000000000000000001006")
)
