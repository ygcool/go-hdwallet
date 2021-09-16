package hdwallet

func init() {
	coins[BNB] = newBNB
}

type bnb struct {
	name   string
	symbol string
	key    *Key

	// bnb token
	contract string
}

func newBNB(key *Key) Wallet {
	return &bnb{
		name:   "Binance Chain",
		symbol: "BNB",
		key:    key,
	}
}

func (c *bnb) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *bnb) GetName() string {
	return c.name
}

func (c *bnb) GetSymbol() string {
	return c.symbol
}

func (c *bnb) GetKey() *Key {
	return c.key
}

func (c *bnb) GetAddress() (string, error) {
	return c.key.AddressBNB(MAINNET)
}

func (c *bnb) GetPrivateKey() (string, error) {
	return c.key.PrivateHex(), nil
}
