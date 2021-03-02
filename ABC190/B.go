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
	n, s, d := readInt(), readInt(), readInt()

	var win bool
	for i := 0; i < n; i++ {
		x, y := readInt(), readInt()
		if x < s && y > d {
			win = true
		}
	}
	if win {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
