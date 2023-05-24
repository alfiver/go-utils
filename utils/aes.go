package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	// 判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	// 补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	unPadding := int(data[length-1]) // 获取填充的个数
	return data[:(length - unPadding)], nil
}
func aesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()                // 判断加密快的大小
	encryptBytes := pkcs7Padding(data, blockSize) // 填充
	crypted := make([]byte, len(encryptBytes))    // 初始化加密数据接收切片
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes) // 执行加密
	return crypted, nil
}
func aesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取块的大小
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)   // 执行解密
	crypted, err = pkcs7UnPadding(crypted) // 去除填充
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

// EncryptByAes Aes加密 后 base64
func EncryptByAes(data []byte, key []byte) (string, error) {
	res, err := aesEncrypt(data, key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes base64解码后 Aes 解密
func DecryptByAes(data string, key []byte) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return aesDecrypt(dataByte, key)
}
