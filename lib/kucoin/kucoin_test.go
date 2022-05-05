package kucoin_test

import (
	"log"
	"os"
	"testing"

	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/lib/kucoin"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var exchange interfaces.Exchange

func TestMain(m *testing.M) {
	//FIXME
	exchange = kucoin.NewKucoin("", "")

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	zap.ReplaceGlobals(logger)

	os.Exit(m.Run())
}

func TestKucoin_GetPairs(t *testing.T) {
	pairs, err := exchange.GetPairs()

	require.NoError(t, err)
	require.NotEmpty(t, pairs)
}

func TestKucoin_DepositAddress(t *testing.T) {
	address, err := exchange.DepositAddress(models.Asset{Symbol: "USDT"})

	require.NoError(t, err)
	require.NotEmpty(t, address)
}

func TestKucoin_GetMarket(t *testing.T) {
	market, err := exchange.GetMarket()

	require.NoError(t, err)
	require.NotEmpty(t, market)
}
