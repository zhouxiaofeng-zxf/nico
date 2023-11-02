package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"github.com/zhouxiaofeng-zxf/nico/util"
)

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := util.NewEntropyPro("nico 床 前 明 月 光 , 疑 是 地 上 霜 .", 128)
	fmt.Println("entropy: ", entropy)

	mnemonic, _ := bip39.NewMnemonic(entropy)
	fmt.Println(" Mnemonic :", mnemonic)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}
