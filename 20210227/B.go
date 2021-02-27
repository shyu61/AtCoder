package main

import (
	"bufio"
	"fmt"
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
	sc.Split(bufio.ScanWords)
	n := readInt()

	min := -1
	for i := 0; i < n; i++ {
		a, p, x := readInt(), readInt(), readInt()
		if x-a > 0 {
			if min == -1 {
				min = p
			} else if min > p {
				min = p
			}
		}
	}
	fmt.Println(min)
}
