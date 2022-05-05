package mexc_test

import (
	"os"
	"testing"

	"github.com/ImPedro29/exchange-sdk/interfaces"
	"github.com/ImPedro29/exchange-sdk/lib/mexc"
	"github.com/ImPedro29/exchange-sdk/models"
	"github.com/stretchr/testify/require"
)

var exchange interfaces.Exchange

func TestMain(m *testing.M) {
	exchange = mexc.NewMexc("", "")

	os.Exit(m.Run())
}

func TestMexc_GetPairs(t *testing.T) {
	pairs, err := exchange.GetPairs()

	require.NoError(t, err)
	require.NotEmpty(t, pairs)
}

func TestMexc_DepositAddress(t *testing.T) {
	depositAddress, err := exchange.DepositAddress(models.Asset{Symbol: "BTC"})

	require.Error(t, err)
	require.Empty(t, depositAddress)
}
