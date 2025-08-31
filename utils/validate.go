package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidateStruct(s any) error {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		validateTag := field.Tag.Get("validate")
		if validateTag == "" {
			continue
		}

		if strings.Contains(validateTag, "required") && isZero(value) {
			return fmt.Errorf("поле %s обязательно для заполнения", field.Name)
		}
	}

	return nil
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
		
	case reflect.Bool:
		return !v.Bool()

	case reflect.Ptr, reflect.Interface:
		return v.IsNil()

	case reflect.Chan, reflect.Func:
		return v.IsNil()

	case reflect.Struct:
		zero := reflect.Zero(v.Type()).Interface()
		current := v.Interface()
		return reflect.DeepEqual(current, zero)

	case reflect.Array, reflect.Slice:
		if v.IsNil() && v.Len() == 0 {
			return true
		}

		return false
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}
