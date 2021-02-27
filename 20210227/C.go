package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	n := readInt()

	slice := make([]int, 0)
	escape := false
	for i := 2; i <= n; i++ {
		if escape == true { break }
		for j := 2; j <= n; j++ {
			num := int(math.Pow(float64(i), float64(j)))
			if num > n {
				if j == 2 {
					escape = true
					break
				} else {
					break
				}
			}
			slice = append(slice, num)
		}
	}

	m := make(map[int]bool)
	count := 0

	for _, elem := range slice {
		if _, ok := m[elem]; !ok {
			m[elem] = true
			count++
		}
	}

	fmt.Println(n - count)
}
