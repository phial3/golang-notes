package cryptokit

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mr "math/rand"
	"time"
)

// 特征：对称加密，加密解密用的是同样的密钥。
// 对称加密：对称加密是最快速、最简单的一种加密方式, 适合经常发送数据的场合
// 非对称加密：加密和解密用的密钥是不同的，通常加密解密的速度比较慢，适合偶尔发送数据的场合。优点是密钥传输方便。

// AES的三要素:
// 密钥：128、192、256
// 填充：
//
//	NoPadding
//	PKCS7Padding
//	ZeroPadding
//	AnsiX923
//	lso10126
//	lso97971
//
// 工作模式：
//
//	CBC、ECB、CTR、CFB、OFB
//
// 工作模式区别:
// ECB模式：
//
//	1、简单
//	2、有利于计算
//	3、相同的明文块经过加密会变成相同的密文块，因此安全性较差
//
// CBC模式：
//	1、无法并行计算，性能上不如ECB
//	2、引入初始化向量IV,增加复杂度。
//	3、安全性高
//
// AES的加密流程:
// 1、把明文按照128bit拆分成若干个明文块
// 2、按照选择的填充模式来填充最后一个明文块
// 3、每个明文块利用AES加密器和密钥，加密成密文块
//
// AES的特点、特征:
// 1、有iv的是特征的是CBC工作模式
// 2、mode和padding标示的加密模式、填充方式
// iv:初始向量
// mode:工作模式
// padding:填充方式

func pkCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt AES加密,CBC
func AesEncrypt(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = pkCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // CBC加密
	//blockMode := cipher.NewCFBEncrypter(block,key[:blockSize]) //CFB加密
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData) //CBC加密
	//blockMode.XORKeyStream(encrypted,origData) //CFB加密
	cryptoBase64 := base64.StdEncoding.EncodeToString(encrypted) //返回base64编码
	return cryptoBase64, nil
}

// AesDecrypt AES解密
func AesDecrypt(cryptoBase64 string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //CBC加密
	//blockMode := cipher.NewCFBDecrypter(block, key[:blockSize]) //CFB加密
	origData := make([]byte, len(cryptoBase64))

	crypted, _ := base64.StdEncoding.DecodeString(cryptoBase64)

	blockMode.CryptBlocks(origData, crypted) //CBC加密
	//blockMode.XORKeyStream(origData, crypted) //CFB加密
	origData = pkCS7UnPadding(origData)
	return origData, nil
}

// DecryptAESPayload 第2种实现方式
// gcm
func DecryptAESPayload(aesKey string, payload string) (string, error) {
	rawPayload, _ := hex.DecodeString(payload)

	cipherData, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(cipherData)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(rawPayload) < nonceSize {
		return "", err
	}

	nonce, ciphertext := rawPayload[:nonceSize], rawPayload[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func EncryptAESPayload(aesKey string, payload string) (string, error) {
	cipherData, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(cipherData)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	sealedData := gcm.Seal(nonce, nonce, []byte(payload), nil)

	return hex.EncodeToString(sealedData), nil
}

func GenerateAESKey() string {
	var symKey = passPhrase(24)
	fmt.Println("symKey raw: ", symKey)
	symKey = base64.StdEncoding.EncodeToString([]byte(symKey))
	fmt.Println(len(symKey))
	fmt.Println("symKey encoded: ", symKey)
	return symKey
}

func passPhrase(n int) string {
	mr.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	buff := make([]rune, n)
	for i := range buff {
		buff[i] = letters[mr.Intn(len(letters))]
	}
	return string(buff)
}
