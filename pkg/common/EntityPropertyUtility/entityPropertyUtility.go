package EntityPropertyUtility

import (
	"errors"
	"reflect"
	"strings"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 11:37
 * @Title:
 * --- --- ---
 * @Desc:
 */
var (
	cacheStructProperties = make(map[string][]*PropertyBase)
)

/**
	获取所有可导出属性信息
 */
func GetProperties(origin interface{}) ([]*PropertyBase, error) {
	if !isStructType(origin) {
		return nil, errors.New("must be has type struct")
	}

	name := getStructName(origin)
	result := cacheStructProperties[name]

	if result == nil {
		typeOf, valueOf := getElem(origin)
		fieldNum := typeOf.NumField()
		var1 := 0
		for var1 < fieldNum {
			fieldType, fieldValue := typeOf.Field(var1), valueOf.Field(var1)
			if fieldValue.CanSet() {
				base := &PropertyBase{
					Name:      fieldType.Name,
					Type:      fieldType.Type.String(),
					OtherInfo: fieldType,
					ValueOf:   fieldValue,
				}
				result = append(result, base)
			}
			var1++
		}
		cacheStructProperties[name] = result
		return result, nil
	}

	return result, nil
}

//转换结构体属性和值为映射
func StructToMap(origin interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	properties, err := GetProperties(origin)
	if err != nil {
		return nil, err
	}

	for _, item := range properties {
		result[item.Name] = getRealValue(item.ValueOf)
	}
	return
}

//拷贝非空字段值
func CopyNotNull(origin, target interface{}) (err error) {
	if origin == nil || target == nil {
		return errors.New("struct not be null")
	}

	originProperties, err := GetProperties(origin)
	if err != nil {
		return err
	}
	targetProperties, err := GetProperties(target)
	if err != nil {
		return err
	}

	for _, originItem := range originProperties {
		for _, targetItem := range targetProperties {
			if targetItem.Name == originItem.Name && targetItem.ValueOf.CanSet() {
				targetItem.ValueOf.Set(originItem.ValueOf)
			}
		}
	}
	return nil
}

//获取结构体字段标签名称是否存在
func CheckTagKey(origin interface{}, field, tag string) (keyExist, valExist bool, err error) {
	if origin == nil || field == "" || tag == "" {
		return false, false, errors.New("Struct required or Field required or Tag required")
	}

	properties, err := GetProperties(origin)
	if err != nil {
		return false, false, err
	}

	var key, val bool
	for _, item := range properties {
		if item.Name == field && strings.HasPrefix(string(item.OtherInfo.Tag), tag+":") {
			key = true
			if item.OtherInfo.Tag.Get(tag) != "" {
				val = true
			}
			return key, val, nil
		}
	}
	return key, val, nil
}

/****************
	private
 */
func isStructType(entity interface{}) bool {
	typeOf, _ := getElem(entity)
	if typeOf.Kind() != reflect.Struct {
		return false
	}
	return true
}

func getElem(entity interface{}) (reflect.Type, reflect.Value) {
	//Elem() 如果取到值非Interface 或 pointer会panic错误，使用Elem()方法转换为源地址的reflect.Value或reflect.Type，才能进行后续操作
	//否则就是指针或接口的Value或Type了
	//而且用了这个必定要传指针或接口类型的参数
	typeOf := reflect.TypeOf(entity).Elem()
	valueOf := reflect.ValueOf(entity).Elem()
	return typeOf, valueOf
}

func getStructName(entity interface{}) string {
	entityType, _ := getElem(entity)
	return entityType.Name()
}

func getRealValue(valueOf reflect.Value) (result interface{}) {
	switch valueOf.Kind() {
	case reflect.Bool:
		result = valueOf.Bool()
	case reflect.Int:
		result = valueOf.Int()
	case reflect.Int64:
		result = valueOf.Int()
	case reflect.Float32:
		result = valueOf.Float()
	case reflect.Float64:
		result = valueOf.Float()
	case reflect.String:
		result = valueOf.String()
	}
	return
}
