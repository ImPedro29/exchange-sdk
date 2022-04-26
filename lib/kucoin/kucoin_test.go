package kucoin_test

import (
	"os"
	"testing"

	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/lib/kucoin"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/stretchr/testify/require"
)

var exchange interfaces.Exchange

func TestMain(m *testing.M) {
	//FIXME
	exchange = kucoin.NewKucoin(
		"https://api.kucoin.com",
		"",
		"")

	os.Exit(m.Run())
}

func TestKucoin_GetPairs(t *testing.T) {
	pairs, err := exchange.GetPairs()

	require.NoError(t, err)
	require.NotEmpty(t, pairs)
}

func TestBinance_DepositAddress(t *testing.T) {
	address, err := exchange.DepositAddress(models.Asset{Symbol: "USDT"})

	require.NoError(t, err)
	require.NotEmpty(t, address)
}
