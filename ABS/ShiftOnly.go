package main

import (
	"fmt"
)

func main() {
	var length int
  fmt.Scan(&length)

	slice := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&slice[i])
	}

	var flag bool
	for count := 0; ; count++ {
		for i := 0; i < length; i++ {
			if slice[i] % 2 == 0 {
				slice[i] /= 2
			} else {
				flag = true
			}
		}
		if flag {
			fmt.Println(count)
			break
		}
	}
}
