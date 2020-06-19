package jwtUtility

import (
	"log"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/19 11:01
 * @Title:
 * --- --- ---
 * @Desc:
 */

var (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcHAiLCJleHAiOjE1OTE4NzIyMjEsInN1YiI6IjExNDUxNCIsInVpZCI6IjExNDUxNCIsImNsaWVudElkIjoiYXBwIn0.CagmtKcf-HmV5dT9QyaNjEv2_MEEODlDyHo7BHqD-r8"
)

func TestMap(t *testing.T) {
	claims := GetMap(token)
	log.Printf("%+v", claims)
	client := GetClient(token)
	log.Printf("%+v", client)
	uid := GetUid(token)
	log.Printf("%d", uid)
	property := GetProperty(token, "uid")
	log.Printf("%s", property)
}
