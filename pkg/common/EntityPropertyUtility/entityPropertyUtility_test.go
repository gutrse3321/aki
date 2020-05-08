package EntityPropertyUtility

import (
	"log"
	"testing"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 11:47
 * @Title:
 * --- --- ---
 * @Desc:
 */
type testStrcut struct {
	Name   string `fuck:"shit"`
	Age    string
	Gender int
	sm     int
	fun    func()
}

type test2Struct struct {
	Name   string `fuck:"shit"`
	Age    string
	Gender int
	sm     int
	fun    func()
}

func TestCheckTag(test *testing.T) {
	test1 := &testStrcut{Name: "tomo", Age: "18", Gender: 1}
	keyExist, valExist, err := CheckTagKey(test1, "Name", "fuck")
	if err != nil {
		panic(err)
	}
	log.Println("key:", keyExist, " val:", valExist)
}

func TestCopyNotNull(test *testing.T) {
	test1 := &testStrcut{Name: "tomo", Age: "18", Gender: 1}
	test2 := &test2Struct{}
	err := CopyNotNull(test1, test2)
	if err != nil {
		panic(err)
	}
	log.Println("main target:", test2)
}

func TestStructToMap(test *testing.T) {
	test1 := &testStrcut{Name: "tomo", Age: "18", Gender: 1}

	resultMap, _ := StructToMap(test1)
	resultMap2, _ := StructToMap(test1)
	log.Println(resultMap)
	log.Println(resultMap2)
}