package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	v, t, s, d := scanInt(), scanInt(), scanInt(), scanInt()
	if (v*t < d && v*s < d) || (v*t > d && v*s > d) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
