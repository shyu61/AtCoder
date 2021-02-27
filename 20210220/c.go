package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func asc(num int) int {
	s := strconv.Itoa(num)
	list := make([]int, 0)
	for _, i := range s {
		j, _ := strconv.Atoi(string(i))
		list = append(list, j)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	ans := 0
	for index, i := range list {
		ans += i * int(math.Pow10(index))
	}

	return ans
}

func desc(num int) int {
	s := strconv.Itoa(num)
	list := make([]int, 0)
	for _, i := range s {
		j, _ := strconv.Atoi(string(i))
		list = append(list, j)
	}

	sort.Ints(list)

	ans := 0
	for index, i := range list {
		ans += i * int(math.Pow10(index))
	}

	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := readInt(), readInt()

	output := n
	for i := 0; i < k; i++ {
		output = desc(output) - asc(output)
	}

	fmt.Println(output)
}