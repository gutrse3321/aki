package remote

import "net/http"

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 15:35
 * @Title: 客户端统一响应模型
 * --- --- ---
 * @Desc:
 */

type Remote struct {
	Code    int
	Message string
	fields  map[string]string

	Model interface{}
}

func Init(ptr *Remote, model interface{}) {
	InitWithCode(ptr, model, http.StatusOK)
}

func InitWithCode(ptr *Remote, model interface{}, code int) {
	ptr.Code = code
	ptr.Model = model
}
