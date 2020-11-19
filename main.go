package main

import (
	"CryptCode/des"
	"bytes"
	"fmt"
)

func main() {
	//des: 块加密
	//des: key、data、mode
	/**
	key:密钥
	data:明文，即将加密的明文
	mode:两种模式，加密和解密
	 */
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家，遇阿姨成家立业"

//	//加密：crypto
//	block, err := des.NewCipher(key)
//	if err != nil {
//		panic("初始化密钥错误，请重试")
//	}
//
//	//dst,src
//	dst := make([]byte, len([]byte(data)))
//
//	//加密过程
//	block.Encrypt(dst, []byte(data))
//
//	fmt.Println("加密后的内容：", dst)
//
//	//解密
//	deBlock, err := des.NewCipher(key)
//	if err != nil {
//		panic("初始化密钥错误，请重试")
//	}
//	deData := make([]byte, len(dst))
//	deBlock.Decrypt(deData,dst)
//
//	fmt.Println(string(deData))
	//一、对数据进行加密  DES的密钥为8字节、3DES密钥长度为24字节
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家"
	//1、得到cipher
	//block, err := des.NewCipher(key)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	////2、对数据明文进行结尾块填充
	//paddingData := PKCS5Padding([]byte(data),block.BlockSize())
	////3、选择模式
	//mode := cipher.NewCBCEncrypter(block, key)
	////4、加密
	//dstData := make([]byte, len(paddingData))
	//mode.CryptBlocks(dstData,paddingData)
	//
	//fmt.Println("加密后的内容", string(dstData))

	//二、对数据进行解密
	//1、DES三元素：key,data,mode
//	key1 := []byte("c1906041")
//	data1 := dstData//待解密的数据
//	block1, err := des.NewCipher(key1)
//	if err != nil {
//		panic(err.Error())
//	}
//	deMode := cipher.NewCBCDecrypter(block1, key1)
//	originalData := make([]byte, len(data1))
//	deMode.CryptBlocks(originalData,data1)
//	originalData = utils.ClearPKCS5Padding(originalData,block1.BlockSize())
//	fmt.Println("解密后的内容:",string(originalData))


	//一、使用DES进行加解密
	key := []byte("20201112")
	data := "窗含西岭千秋雪，门泊东吴万里船"
	cipherText,err := des.DESDeCrypt([]byte(data),key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText, err := des.DESDeCrypt(cipherText,key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("解密后：",string(originText))

	//3DES加解密
	//key1 := []bytes("")//3DES的密钥长度是32字节
}

/**
 *明文数据尾部填充
 */
func PKCS5Padding (text []byte, blockSize int) []byte{
	//计算要填充的块大小
	paddingSize := blockSize - len(text)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	fmt.Println("明文尾部内容：",paddingText)
	return append(text, paddingText...)
}