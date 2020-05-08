package EntityPropertyUtility

import "reflect"

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 11:40
 * @Title:
 * --- --- ---
 * @Desc:
 */
type PropertyBase struct {
	Name      string
	Type      string
	OtherInfo reflect.StructField
	ValueOf   reflect.Value
}
