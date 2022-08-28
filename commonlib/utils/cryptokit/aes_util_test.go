package cryptokit

import (
	"fmt"
	"testing"
)

var (
	text   = "123456"                   // 你要加密的数据
	AesKey = []byte("0123456789123456") // 秘钥,对称秘钥长度必须是16的倍数
)

func TestAesEncrypt(t *testing.T) {
	cryptoBase64, err := AesEncrypt([]byte(text), AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后:", cryptoBase64) // KaknGVd4nFWtpiXyZ540SA==

	origin, err := AesDecrypt(cryptoBase64, AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后明文:", string(origin)) //123456
}

func TestAesDecrypt(t *testing.T) {

}
