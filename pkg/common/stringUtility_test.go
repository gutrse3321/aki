package common

import (
	"crypto"
	"fmt"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 12:20
 * @Title:
 * --- --- ---
 * @Desc:
 */

func TestReplaceSpace(b *testing.T) {
	s := `他  fu你 	c
k妈`
	space := ReplaceSpace(s)
	count := GetChineseCharCount(s)
	fmt.Println(space)
	fmt.Println(count)
}

func TestEncryptor(b *testing.T) {
	str := "123456"
	fmt.Println(Encrypt(crypto.SHA1, str))
	fmt.Println(Encrypt(crypto.MD5, str))
	fmt.Println(Encrypt(crypto.SHA256, str))
	fmt.Println(Base64UTF8Encoder(str))
	println(Base64UTF8Decoder(Base64UTF8Encoder(str)))
}
