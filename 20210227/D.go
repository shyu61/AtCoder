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

func calcSum(s string) []int {
	var count, sum int
	sumSlice := make([]int, 9)
	for i := 1; i <= 9; i++ {
		for _, j := range s[:4] {
			for _, k := range s[:4] {
				if j == k { count++ }
			}
			sum += int(j * math.Pow10(count))
		}
		sumSlice[i-1] = sum
	}
	return sumSlice
}

func main() {
	k, s, t := readInt(), readString(), readString()
	sumS := make([]int, 9)
	sumT := make([]int, 9)

	sumS = calcSum(s)
	sumT = calcSum(t)

	pattern := make(map[int]int)
	for xIndex, x := range sum_s {
		for yIndex, y := range sum_t {
			if x > y {
				pattern[xIndex] = yIndex
			}
		}
	}
	
}
