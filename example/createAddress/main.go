package main

import (
	"fmt"

	"github.com/ygcool/go-hdwallet"
)

var (
	mnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

func main() {
	//mnemonic, _ = hdwallet.NewMnemonic(12, "")

	master, err := hdwallet.NewKey(
		hdwallet.Mnemonic(mnemonic),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("助记词：", mnemonic)
	fmt.Println()

	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.FIL))
	prikey, _ := wallet.GetPrivateKey()
	fmt.Println("FIL私钥：", prikey)
	address, _ := wallet.GetAddress()
	fmt.Println("FIL: ", address)
	addressTestnet, _ := wallet.GetKey().AddressFIL(hdwallet.TESTNET)
	fmt.Println("FIL-testnet: ", addressTestnet)

	//wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BNB))
	//fmt.Println("BNB私钥：", wallet.GetKey().PrivateHex())
	//address, _ := wallet.GetKey().AddressBNB(hdwallet.MAINNET)
	//fmt.Println("BNB: ", address)
	//addressTestnet, _ := wallet.GetKey().AddressBNB(hdwallet.TESTNET)
	//fmt.Println("BNBTestnet: ", addressTestnet)

	//wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.TRX))
	//fmt.Println("TRX私钥：", wallet.GetKey().PrivateHex())
	//address, _ := wallet.GetAddress()
	//fmt.Println("TRX: ", address)

	//wallet, _ = master.GetWallet(hdwallet.Purpose(hdwallet.ZeroQuote+44), hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(0))
	//address, _ = wallet.GetAddress()                                // 1AwEPfoojHnKrhgt1vfuZAhrvPrmz7Rh44
	//addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()             // bc1qdnavt2xqvmc58ktff7rhvtn9s62zylp5lh5sry
	//addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH() // 39vtu9kWfGigXTKMMyc8tds7q36JBCTEHg
	////btcwif, err := wallet.GetKey().PrivateWIF(true)
	//btcwif, err := wallet.GetPrivateKey()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("BTC私钥：", btcwif)
	//fmt.Println("BTC: ", address, addressP2WPKH, addressP2WPKHInP2SH)

	// BCH: 1CSBT18sjcCwLCpmnnyN5iqLc46Qx7CC91
	//wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.BCH))
	//address, _ = wallet.GetAddress()
	//addressBCH, _ := wallet.GetKey().AddressBCH()
	//fmt.Println("BCH: ", address, addressBCH)

	// LTC: LLCaMFT8AKjDTvz1Ju8JoyYXxuug4PZZmS
	//wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.LTC))
	//address, _ = wallet.GetAddress()
	//fmt.Println("LTC: ", address)

	// DOGE: DHLA3rJcCjG2tQwvnmoJzD5Ej7dBTQqhHK
	//wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.DOGE))
	//address, _ = wallet.GetAddress()
	//fmt.Println("DOGE:", address)

	// ETH: 0x37039021cBA199663cBCb8e86bB63576991A28C1
	//wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
	//address, _ = wallet.GetAddress()
	//fmt.Println("ETH私钥：", wallet.GetKey().PrivateHex())
	//fmt.Println("ETH: ", address)

	// ETC: 0x480C69E014C7f018dAbF17A98273e90f0b0680cf
	//wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETC))
	//address, _ = wallet.GetAddress()
	//fmt.Println("ETC: ", address)
}
