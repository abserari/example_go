package main

import (
	"fmt"
	"reflect"
	"sync"
)

func add(a, b int) (int, error) { return a + b, nil }

func main() {
	var services sync.Map
	services.LoadOrStore("add", add)

	call, _ := services.Load("add")
	rep := reflect.ValueOf(call).Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})

	fmt.Println(rep)
	for _, val := range rep {
		fmt.Println(val.Interface())
	}
}
