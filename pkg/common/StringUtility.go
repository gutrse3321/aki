package common

import (
	"net/url"
	"regexp"
	"strings"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 12:09
 * @Title:
 * --- --- ---
 * @Desc:
 */
func ComposeUrl(builder strings.Builder, param map[string]interface{}) {
	if len(param) == 0 {
		return
	}

	builder.WriteString("?")
	builder.WriteString(urlParameterEncoder(param))
}

func urlParameterEncoder(param map[string]interface{}) string {
	urlValues := url.Values{}

	for key, value := range param {
		urlValues.Add(key, value.(string))
	}

	return urlValues.Encode()
}

/**
	验证字符串是否为空
 */
func IsNull(str string) bool {
	str = ReplaceSpace(str)
	return len(str) == 0
}

/**
	返回是否存在制表符等特殊字符
 */
func ReplaceSpace(str string) string {
	if len(str) == 0 {
		return ""
	}

	reg := regexp.MustCompile(`[\\s*|\t*|\r*|\n*]`)
	str = reg.FindString(str)
	return str
}

func GetChineseCharCount(str string) (count int) {
	reg := regexp.MustCompile(`[一-龥]`)
	findString := reg.FindAllString(str, -1)
	if findString != nil {
		count = len(findString)
	}
	return
}
