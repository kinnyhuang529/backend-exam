package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	// 用來記錄處理過的指標 防止無限循環
	visited := make(map[uintptr]bool)

	var walk func(v reflect.Value)
	walk = func(v reflect.Value) {
		// 如果是指標 進去拿內容 並記錄位址防止重複
		if v.Kind() == reflect.Ptr {
			if v.IsNil() || visited[v.Pointer()] {
				return
			}
			visited[v.Pointer()] = true
			walk(v.Elem())
			return
		}

		// 根據型別做簡單處理
		switch v.Kind() {
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				walk(v.Field(i))
			}
		case reflect.String:
			if v.CanSet() {
				v.SetString(strings.TrimSpace(v.String()))
			}
		}
	}

	walk(reflect.ValueOf(a))
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
