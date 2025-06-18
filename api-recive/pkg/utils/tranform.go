package utils

import (
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
    objValue := reflect.ValueOf(obj)
    objType := objValue.Type()

    if objType.Kind() != reflect.Struct {
        return nil
    }

    result := make(map[string]interface{})
    for i := 0; i < objValue.NumField(); i++ {
        field := objType.Field(i)
        fieldValue := objValue.Field(i).Interface()
        result[field.Name] = fieldValue
    }

    return result
}
