package main

import (
	"fmt"
	"reflect"
)

func max(num ...uint64) uint64 {
	max := num[0]
	for _, i := range num[1:] {
		if i > max { max = i }
	}
	return max
}

func main() {
	ja := "こんにちは"
	// en := "hello"

	for i, s := range ja {
		fmt.Println(string(s))
		fmt.Println(reflect.TypeOf(i))
	}
}