package relayerincentivize

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"

	cmm "github/HaoyangLiu/contract-simulation/contract/common"
)

func TestSimulateRelayerincentivize(t *testing.T) {
	client, err := ethclient.Dial(cmm.TestnetEndpoint)
	assert.NoError(t, err)

	callOpt, err := cmm.GetCallOpts()
	assert.NoError(t, err)
	relayerincentivizeInstance, err := NewRelayerincentivize(cmm.RelayerIncentivize, client)
	assert.NoError(t, err)

	roundSequence, err := relayerincentivizeInstance.RoundSequence(callOpt)
	assert.NoError(t, err)
	t.Log(roundSequence)
}
