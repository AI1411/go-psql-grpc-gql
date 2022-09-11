package helper

import (
	"fmt"
	"reflect"
	"time"
)

func IsNilOrEmpty(x interface{}) bool {
	if x == nil {
		return true
	}
	wkTypeName := fmt.Sprintf("%T", x)
	switch wkTypeName {
	case "bool", "uint", "uint8", "uint16", "uint32", "uint64", "int", "int8", "int16", "int32", "int64",
		"float32", "float64", "complex64", "complex128", "byte", "uintptr", "error":
		return false
	case "rune", "string":
		if x == "" {
			return true
		}
		return false
	case "time.Time":
		return x.(time.Time).IsZero()
	}
	xVal := reflect.ValueOf(x)
	if xVal.IsNil() {
		return true
	}

	if xVal.Kind() == reflect.Slice {
		if xVal.Len() == 0 {
			return true
		}
	}

	return false
}
