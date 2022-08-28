package cryptokit

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateRSAKey 生成RSA私钥和公钥，保存到文件中,bits 证书大小
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1的DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("./private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("./public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

// RsaEncrypt RSA加密 ,plainText 要加密的数据 ,path 公钥匙文件地址
func RsaEncrypt(plainText []byte, path string) string {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回base编码密文，否则看不懂
	crypteBase64 := base64.StdEncoding.EncodeToString(cipherText) //返回base64编码
	return crypteBase64
}

// RsaDecrypt RSA解密 ,cipherText 需要解密的byte数据 ,path 私钥文件路径
func RsaDecrypt(encryptBase64 string, path string) []byte {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//base64解码
	cipherText, _ := base64.StdEncoding.DecodeString(encryptBase64)
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}

func main() {
	//生成密钥对，保存到文件
	GenerateRSAKey(2048)

	//加密
	data := []byte("123456") // 原数据
	encryptBase64 := RsaEncrypt(data, "./public.pem")
	fmt.Println(encryptBase64)
	//gHudPWYOAPsq7C3RVG/ilykMgKbvqsxQLFhaT9TbdjsM0FDy1odF3AuJkoakLaWMdueigC8UCdQEJhFADh0skPwAonuecqpk3UqljBeDPPgCSPhy2S5nrQMqZuK+EB+fdGnu8ieMmkIQACFQShhnFlqfXbvY5pGi+rjdXTaaXFtxRygtaA+OGHBh7rNWprEainoOC4WeLiHg/v1T6Pk6NiAR3eYjq0L4vadRvkV8MhyjUVifBziW5EXJk00omraoQjAu+1P+PV/wVlFj8BA/4p3ZfMrhzAkTpI/Bewsz5KHJ2xeQg4wEWMrbK8MyHTcH0GRLG+Pk0D/tGawpgW52TA==

	// 解密
	decrypt := RsaDecrypt(encryptBase64, "./private.pem")
	fmt.Println(string(decrypt)) //123456
}
