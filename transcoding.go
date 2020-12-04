package tools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

//MD5方法
func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//Base64编码方法
func Base64Encoder(str string) string {
	src := []byte(str)
	return base64.StdEncoding.EncodeToString(src)
}

//Base64解码方法
func Base64Decode(str string) (string, error) {
	src, err := base64.StdEncoding.DecodeString(str)
	return string(src), err
}