//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

/*
DES

DES是一种对称加密算法，又称为美国数据加密标准。DES加密时以64位分组对数据进行加密，
加密和解密都使用的是同一个长度为64位的密钥，实际上只用到了其中的56位，密钥中的第8、16…64位用来作奇偶校验。

DES有ECB（电子密码本）和CBC（加密块）等加密模式。

DES算法的安全性很高，目前除了穷举搜索破解外， 尚无更好的的办法来破解。其密钥长度越长，破解难度就越大。
填充和去填充函数。
*/

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}
func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}
func main() {
	//在这里，DES中使用的密钥key只能为8位。
	key := []byte("2fa6c1e9")
	str := "I love this beautiful world!"
	strEncrypted, err := Encrypt(str, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	strDecrypted, err := Decrypt(strEncrypted, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)
}

//Output:
//Encrypted: 5d2333b9fbbe5892379e6bcc25ffd1f3a51b6ffe4dc7af62beb28e1270d5daa1
//Decrypted: I love this beautiful world!
