package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)

func readString() string {
	sc.Scan()
	return sc.Text()
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func main() {
	s := readString()

	for index, i := range s {
		if index % 2 == 0 && check_regexp(`[A-Z]`, string(i)) {
			fmt.Println("No")
			return
		}
		if index % 2 == 1 && check_regexp(`[a-z]`, string(i)) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
