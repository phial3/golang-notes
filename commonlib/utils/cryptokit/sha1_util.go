package cryptokit

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// 加密算法：
//一、对称加密算法 ：常⽤的算法包括DES、3DES、AES、DESX、Blowfish、RC4、RC5、RC6。推荐⽤AES。
//⼆、⾮对称加密算法：常见的⾮对称加密算法：RSA、DSA（数字签名⽤）、ECC（移动设备⽤）、Diffie-Hellman、El Gamal。推荐⽤ECC（椭圆曲线密码编码学）。
//三、散列算法（Hash算法---单向加密算法）：常见的Hash算法：MD2、MD4、MD5、HAVAL、SHA、SHA-1、HMAC、HMAC-MD5、HMAC-SHA1。推荐MD5、SHA-1。
//
//对称和非对称区别：
//对称加密：加解密密码相同
//非对称加密：加解密密码不相同

// 计算字符串的sha1散列值
func Sha1Str(s string) string {
	data := []byte(s)
	m := sha1.New()
	_, _ = m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

// 计算字符串的 shaX 散列值
// x为1/256/512
func ShaXStr(str string, x uint16) string {
	return string(shaXStr([]byte(str), x))
}

// 计算字符串的 shaX 散列值
// x为1/256/512
func shaXStr(str []byte, x uint16) []byte {
	var h hash.Hash
	switch x {
	case 1:
		h = sha1.New()
	case 256:
		h = sha256.New()
	case 512:
		h = sha512.New()
	default:
		panic("[shaXStr] x must be in [1, 256, 512]")
	}

	h.Write(str)
	hBytes := h.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}
