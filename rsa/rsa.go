package rsa

import (
	"CryptCode/utils"
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"flag"
)

/**
 * 私钥：
 * 公钥：
 */
func CreatePairKeys() (*rsa.PrivateKey, crypto.PublicKey, error) {
	//1、先生成私钥
	var bits int
	flag.IntVar(&bits,"b",2048,"密钥长度")

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	//2、根据私钥生产公钥
	publicKey := privateKey.Public()

	//3、将公私钥返回
	return privateKey, publicKey, nil
}


func RSAEncrypt(key crypto.PublicKey, data []byte) []byte {
	return rsa.EncryptPKCS1v15(rand.Reader,&key, data)
}


/**
 * 使用RSA算法对数据进行解密，女兵返回解密后的明文
 */
func RSADecrypt(Private *rsa.PrivateKey,cipher []byte)([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader,private, cipher)
}

func RSASign (private *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashed := utils.Md5Hash()
	return rsa.SignPKCS1v15(rand.Reader, private, crypo.MD5, hashed)
}

func RSAVerify(pub rsa.PrivateKey, data []byte, signText []byte) (bool, error) {
	hashed := utils.Md5Hash(data)
	err := rsa.VerifyPKCS1v15(&pub, crypto.MD5, hashed, signText)
	return err == nil, err
}