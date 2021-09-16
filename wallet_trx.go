package hdwallet

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func init() {
	coins[TRX] = newTRX
}

type trx struct {
	name   string
	symbol string
	key    *Key

	// trc20 token
	contract string
}

func newTRX(key *Key) Wallet {
	return &trx{
		name:   "Tron",
		symbol: "TRX",
		key:    key,
	}
}

func (c *trx) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *trx) GetName() string {
	return c.name
}

func (c *trx) GetSymbol() string {
	return c.symbol
}

func (c *trx) GetKey() *Key {
	return c.key
}

func (c *trx) GetAddress() (string, error) {
	a := address.PubkeyToAddress(*c.key.PublicECDSA)

	return a.String(), nil
}

func (c *trx) GetPrivateKey() (string, error) {
	return c.key.PrivateHex(), nil
}
