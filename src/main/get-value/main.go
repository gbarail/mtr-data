package main

import (
	"fmt"
	"reflect"
	"strings"
)

func GetValue[T any](x any, key string) (*T, error) {
	var (
		v   any = x
		err error
	)

	keys := strings.Split(key, "/")
	for _, k := range keys {
		v, err = getValue(v, k)
		if err != nil {
			return nil, err
		}
	}
	return check[T](v)
}

func getValue(x any, key string) (any, error) {
	if key == "" {
		return x, nil
	}

	v := reflect.ValueOf(x)
	var v1 any
	switch reflect.TypeOf(x).Kind() {
	case reflect.Map:
		v1 = v.MapIndex(reflect.ValueOf(key)).Interface()
	case reflect.Struct:
		v1 = v.FieldByName(key).Interface()
	case reflect.Ptr:
		v1 = v.Elem().Interface()
		return getValue(v1, key)
	default:
		return nil, fmt.Errorf("cannot get value from %+v with key \"%s\"", x, key)
	}
	return v1, nil
}

func check[T any](x any) (*T, error) {
	v, ok := x.(T)
	if !ok {
		var y T
		return nil, fmt.Errorf("cannot convert value %+v to type \"%T\"", x, y)
	}
	return &v, nil
}

func main() {
	var a = struct {
		X int
		Y *struct {
			Z1 string
			Z2 float64
		}
	}{
		X: 1,
		Y: &struct {
			Z1 string
			Z2 float64
		}{
			Z1: "hello",
			Z2: 3.14,
		},
	}

	v, err := GetValue[float64](a, "Y/Z2")
	fmt.Println(*v, err)
}
