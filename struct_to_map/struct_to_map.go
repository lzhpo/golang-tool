package struct_to_map

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// UnmarshalToMap 方法1：JSON序列化的方式
func UnmarshalToMap(inBody interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	inBodyBytes, inBodyMarshalError := json.Marshal(&inBody)
	if inBodyMarshalError != nil {
		log.Println(fmt.Sprintf("json Marshal %v Error:%v", inBody, inBodyMarshalError))
	}

	var inBodyMap map[string]interface{}
	inBodyUnmarshalError := json.Unmarshal(inBodyBytes, &inBodyMap)
	if inBodyUnmarshalError != nil {
		log.Println(fmt.Sprintf("json Unmarshal %v Error:%v", inBody, inBodyUnmarshalError))
	}

	for k, v := range inBodyMap {
		fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
		out[k] = v
	}

	return out, nil
}

// ReflectToMap 方法2：反射
func ReflectToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		// Elem返回接口v包含：Interface、Ptr
		// 或指针v指向的值。
		// 如果v的Kind不是Interface或Ptr，就报错panics。
		// 如果v为nil，则返回零值。
		inValue = inValue.Elem()
	}

	// 非结构体返回错误提示
	if inValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("toMap only accepts struct or struct pointer; got %T", inValue)
	}

	inValueType := inValue.Type()
	for i := 0; i < inValue.NumField(); i++ {
		field := inValueType.Field(i)
		log.Println("Type:", inValue.Field(i).Type())
		if tagValue := field.Tag.Get(tagName); tagName != "" {
			out[tagValue] = inValue.Field(i).Interface()
			fmt.Printf("key:%v value:%v value type:%T\n", out[tagValue], inValue.Field(i).Interface(), inValue.Field(i).Interface())
		}
	}

	return out, nil
}
