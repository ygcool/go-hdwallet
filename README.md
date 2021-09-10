## go-hdwallet

Golang实现的多币种钱包

鸣谢：[https://github.com/foxnut/go-hdwallet](https://github.com/foxnut/go-hdwallet)

## 支持的币种

- BTC
- LTC
- DOGE
- DASH
- ETH
- ETC
- BCH
- QTUM
- USDT
- IOST
- USDC
- TRX
- BNB(Binance Chain)
- FIL

## 安装

```sh
go get -v -u github.com/ygcool/go-hdwallet
```

## 示例
更多示例请参考example

```go
package main

import (
    "fmt"

    "github.com/tyler-smith/go-bip39"
    "github.com/ygcool/go-hdwallet"
)

func main() {

    // 128: 12 phrases
    // 256: 24 phrases
    mnemonic, _ := hdwallet.NewMnemonic(12, "")
       
    master, err := hdwallet.NewKey(
        hdwallet.Mnemonic(mnemonic),
    )
    if err != nil {
        panic(err)
    }
    fmt.Println("助记词：", mnemonic)

    wallet, _ := master.GetWallet(hdwallet.Purpose(hdwallet.ZeroQuote+44), hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(0))
    address, _ := wallet.GetAddress()  // 1AwEPfoojHnKrhgt1vfuZAhrvPrmz7Rh44
    addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH() // bc1qdnavt2xqvmc58ktff7rhvtn9s62zylp5lh5sry
    addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH() // 39vtu9kWfGigXTKMMyc8tds7q36JBCTEHg 
    // addressP2WPKHInP2SH的特别说明:这个隔离见证的地址，是属于当前wif私钥的（默认bip44）。 
    // 假设你是用生成的助记词导入到imtoken中，对应的隔离见证地址不是这个。
    // 若想和imtoken一致，请在 master.GetWallet 时传入 hdwallet.ZeroQuote+49 （即bip49）得到的隔离见证地址和对应私钥即可
    btcwif, err := wallet.GetKey().PrivateWIF(true)
    if err != nil {
        panic(err)
    }
    fmt.Println("BTC私钥：", btcwif)
    fmt.Println("BTC: ", address, addressP2WPKH, addressP2WPKHInP2SH)

    // BCH: 1CSBT18sjcCwLCpmnnyN5iqLc46Qx7CC91
    wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.BCH))
    address, _ = wallet.GetAddress()
    addressBCH, _ := wallet.GetKey().AddressBCH()
    fmt.Println("BCH: ", address, addressBCH)

    // LTC: LLCaMFT8AKjDTvz1Ju8JoyYXxuug4PZZmS
    wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.LTC))
    address, _ = wallet.GetAddress()
    fmt.Println("LTC: ", address)

    // DOGE: DHLA3rJcCjG2tQwvnmoJzD5Ej7dBTQqhHK
    wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.DOGE))
    address, _ = wallet.GetAddress()
    fmt.Println("DOGE:", address)

    // ETH: 0x37039021cBA199663cBCb8e86bB63576991A28C1
    wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETH))
    address, _ = wallet.GetAddress()
	fmt.Println("ETH私钥：", wallet.GetKey().PrivateHex())
    fmt.Println("ETH: ", address)

    // ETC: 0x480C69E014C7f018dAbF17A98273e90f0b0680cf
    wallet, _ = master.GetWallet(hdwallet.CoinType(hdwallet.ETC))
    address, _ = wallet.GetAddress()
    fmt.Println("ETC: ", address)
}
```
