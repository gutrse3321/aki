package common

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"log"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 15:18
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	strPtr *string
	sum *[]byte
)

func Encrypt(operation crypto.Hash, str string) string {
	strPtr = &str
	switch operation {
	case crypto.MD5:
		md5Hex()
	case crypto.SHA1:
		sha1Hex()
	case crypto.SHA256:
		sha256Hex()
	case crypto.SHA384:
		sha384Hex()
	case crypto.SHA512:
		sha512Hex()
	}

	return hex.EncodeToString(*sum)
}

func md5Hex() {
	md := md5.New()
	md.Write([]byte(*strPtr))
	bytes := md.Sum(nil)
	sum = &bytes
}

func sha1Hex() {
	sha := sha1.New()
	sha.Write([]byte(*strPtr))
	bytes := sha.Sum(nil)
	sum = &bytes
}

func sha256Hex() {
	sha := sha256.New()
	sha.Write([]byte(*strPtr))
	bytes := sha.Sum(nil)
	sum = &bytes
}

func sha384Hex() {
	sha := crypto.SHA384.New()
	sha.Write([]byte(*strPtr))
	bytes := sha.Sum(nil)
	sum = &bytes
}

func sha512Hex() {
	sha := sha512.New()
	sha.Write([]byte(*strPtr))
	bytes := sha.Sum(nil)
	sum = &bytes
}

func Base64UTF8Encoder(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64UTF8Decoder(str string) string {
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("base64 decode error:", err)
		return ""
	}
	return string(decodeString)
}
