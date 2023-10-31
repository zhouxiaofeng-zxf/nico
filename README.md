# nico
nico是一种最简单易于记忆的个性化助记词生成方案.用户可以自定义任何语句,支持多语言没有语言限制,映射出对应的助记词,从而代替助记词的记忆.
这是基于BIP-39方案的升级,使用户更加方便有效的获得自己的助记词,不在当心助记词的忘记,甚至不用记下助记词,仅靠脑部记忆自定义的语句就行.
代码原理主要是将自定义语句进行sha256算法处理生成对应的熵,在通过熵获得对应的助记词,并且此方法完全兼容bip39方案.

### 动机
助记词,私钥是用户参与区块链和web3的钥匙.本人是web3的开发技术人员,有时候会经常用到不同的钱包,管理很多的助记词和公私钥.虽然bip39方案通过助记词方式已经大大提高了用户的体验,
让用户越来越方便的管理自己的钱包和私钥地址.但是还是免不了助记词的难以记忆,需要用其他途径把助记词给备份.基于该痛点,我思考了以自定义语句方式生产对应的助记词,只需记住自定义语句即可


## 举例(每组字符之间以空格分开 UTF-8)

### 案例1 诗 将进酒 李白 : 
自定义语句: 君 不 见 ， 黄 河 之 水 天 上 来 ， 奔 流 到 海 不 复 回

对应的熵: 6140274fd3a01382a89c0e230303e051

对应的助记词: genuine abuse stadium point abuse scout pen limb cart blossom way penalty

### 案例2 诗 将进酒 李白(进行不同的组合排列,把不见和天上来合为一组数据):  
自定义语句: 君 不见 ， 黄 河 之 水 天上来 ， 奔 流 到 海 不 复 回

对应的熵: 951837426fe31789151198b72f29b251

对应的助记词: never script spatial thank cousin setup february credit rescue junk hold permit

### 案例3 歌曲 my love westlife (支持多语言,不受语言限制):
自定义语句: an empty street an empty house a hole inside my heart 西域男孩 my love

对应的熵: 3644499e6722c180a678ecedf69c46a9

对应的助记词: curve car guide soft clown scare okay budget unknown reject balance false

#### 说明:用户可以根据自己的喜好,方便记忆的自定义语句来获取到对应的助记词和熵,当然为了保证语句不易被恶意获取,可以对自定义语句进行简单的加工增加一定的复杂度.用户可自由发挥,改变组合,增加一些随机密码等.
例如

原句:  君 不 见 ， 黄 河 之 水 天 上 来 ， 奔 流 到 海 不 复 回             ->对应助记词: genuine abuse stadium point abuse scout pen limb cart blossom way penalty

修饰后: 君 不 见 ， 黄 河 之 水 天 上 来 ， 奔 流 到 海 不 复 回 1qaz@wsx   ->对应助记词: harsh door spin resemble like shed liquid liar climb flame wheel raven




## 安装：

```console
git clone https://github.com/zhouxiaofeng-zxf/nico.git
```

### 可以参考代码中的main.go和bip39_pro_test.go文件

```go
package main

import (
  "fmt"
  "github.com/tyler-smith/go-bip39"
  "github.com/tyler-smith/go-bip32"
)

func main(){
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := NewEntropyPro("nico 床 前 明 月 光 , 疑 是 地 上 霜 .", 128)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey)
}
```


