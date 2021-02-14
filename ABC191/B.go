package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	a, x := scanInt(), scanString()

	ans := ""
	for i := 0; i < a; i++ {
		j := scanString()
		if j != x {
			ans += j + " "
		}
	}

	fmt.Println(strings.TrimSpace(ans))
}
