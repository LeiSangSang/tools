package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//加密
func AesEncryptSimple(origData []byte, key string, iv string) ([]byte, error) {
	return aesEncryptPkcs5(origData, []byte(key), []byte(iv))
}

func aesEncryptPkcs5(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return aesEncrypt(origData, key, iv, pKCS5Padding)
}

func aesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//解密
func AesDecryptSimple(crypted []byte, key string, iv string) ([]byte, error) {
	return aesDecryptPkcs5(crypted, []byte(key), []byte(iv))
}

func aesDecryptPkcs5(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return aesDecrypt(crypted, key, iv, pKCS5UnPadding)
}

func aesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPaddingFunc(origData)
	return origData, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}
