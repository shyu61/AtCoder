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
	a, b, c := readInt(), readInt(), readInt()

	var output string
	if a > b {
		output = "Takahashi"
	} else if a < b {
		output = "Aoki"
	} else {
		if c == 0 {
			output = "Aoki"
		} else {
			output = "Takahashi"
		}
	}

	fmt.Println(output)
}
