package remote

import (
	"errors"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 15:57
 * @Title:
 * --- --- ---
 * @Desc:
 */

func ResolveRemote(remote Remote) (interface{}, error) {
	if remote.Code == http.StatusOK {
		return remote.Model, nil
	}
	return nil, errors.New("resolve remote 数据异常")
}
