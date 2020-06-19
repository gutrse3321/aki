package jwtUtility

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 10:47
 * @Title:
 * --- --- ---
 * @Desc:
 */

func GetMap(token string) jwt.MapClaims {
	tokenInfo, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return "aki", nil
	})
	claims, ok := tokenInfo.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}
	return claims
}

func GetProperty(token, key string) string {
	hashMap := GetMap(token)
	if hashMap != nil {
		return hashMap[key].(string)
	}
	return ""
}

func GetUid(token string) int64 {
	hashMap := GetMap(token)
	uid, _ := strconv.ParseInt(hashMap["uid"].(string), 10, 64)
	return uid
}

func GetClient(token string) string {
	hashMap := GetMap(token)
	return hashMap["clientId"].(string)
}
