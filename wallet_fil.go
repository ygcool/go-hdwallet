package hdwallet

func init() {
	coins[FIL] = newFIL
}

type fil struct {
	name   string
	symbol string
	key    *Key
}

func newFIL(key *Key) Wallet {
	return &fil{
		name:   "Filecoin",
		symbol: "FIL",
		key:    key,
	}
}

func (c *fil) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *fil) GetName() string {
	return c.name
}

func (c *fil) GetSymbol() string {
	return c.symbol
}

func (c *fil) GetKey() *Key {
	return c.key
}

func (c *fil) GetAddress() (string, error) {
	return c.key.AddressFIL(MAINNET)
}

func (c *fil) GetPrivateKey() (string, error) {
	return c.key.PrivateHex(), nil
}
