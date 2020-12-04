package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func GetSignMd5(params map[string]string)string{
	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		dataParams = dataParams + k + "=" + params[k] + "&"
	}
	dataParams = dataParams[0 : len(dataParams)-1]
	sign := strings.ToUpper(Md5(dataParams))
	return sign
}

func GetSignHex(params map[string]string)string{
	var dataParams string
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		dataParams = dataParams + k + "=" + params[k] + "&"
	}
	dataParams = dataParams[0 : len(dataParams)-1]
	h := sha1.New()
	h.Write([]byte(dataParams))
	bs := h.Sum(nil)
	sign := hex.EncodeToString(bs)
	return sign
}