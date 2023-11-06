package util

import (
	"encoding/hex"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
	"testing"
)

func TestEntropyPro(t *testing.T) {
	entryPro, err := NewEntropyFromWords("an empty street an empty house a hole inside my heart 西域男孩 my love", 128)
	entryProString := hex.EncodeToString(entryPro)
	fmt.Println("entryProString: ", entryProString)
	m, err := bip39.NewMnemonic(entryPro)
	fmt.Println("entryPro Mnemonic :", m)
	ent, err := bip39.EntropyFromMnemonic(m)
	//ent, err := EntropyFromMnemonic(m)
	entString := hex.EncodeToString(ent)
	fmt.Println("entString: ", entString)

	seed := bip39.NewSeed(m, "")
	fmt.Println("seed:", hex.EncodeToString(seed))

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	address := account.Address.Hex()
	privateKeyEth, _ := wallet.PrivateKeyHex(account)
	var publicKeyEth, _ = wallet.PublicKeyHex(account)

	fmt.Println("address0:", address)         // id为0的钱包地址
	fmt.Println("privateKey:", privateKeyEth) // 私钥
	fmt.Println("publicKey:", publicKeyEth)   // 公钥

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	account2, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	address2 := account2.Address.Hex()
	privateKeyEth2, _ := wallet.PrivateKeyHex(account2)
	var publicKeyEth2, _ = wallet.PublicKeyHex(account2)

	fmt.Println("address2:", address2)             // id为1的钱包地址
	fmt.Println("privateKeyEth2:", privateKeyEth2) // 私钥
	fmt.Println("publicKeyEth2:", publicKeyEth2)   // 公钥
}
