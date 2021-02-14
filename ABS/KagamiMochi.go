package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, y := scanInt(), scanInt()

	a, b, c := -1, -1, -1

	for i := 0; i <= n; i++ {
		j := sort.Search(n-i+1, func(j int) bool {
			return 10000*i + 5000*j + 1000*(n-i-j) >= y
		})

		if j == n-i+1 {
			continue
		}

		if 10000*i + 5000*j + 1000*(n-i-j) == y {
			a, b, c = i, j, n-i-j
			break
		}
	}
	
	fmt.Printf("%d %d %d\n", a, b, c)
}
