package redis

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/5/29 16:34
 * @Title:
 * --- --- ---
 * @Desc:
 */
var WireSet = wire.NewSet(New, NewOptions)
