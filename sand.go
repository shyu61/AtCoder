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

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := readString()

	var size, ans int
	for i := 0; i < len(s); i++ {
		str := s[i]
		if str == 'A' || str == 'T' || str == 'G' || str == 'C' {
			size++
			if ans < size {
				ans = size
			}
		} else {
			size = 0
		}
	}
	fmt.Println(ans)
}
