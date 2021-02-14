package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	sort.Slice(cards, func(i, j int) bool { return cards[i] > cards[j] })

	var alice, bob int
	for i, c := range cards {
		if i%2 == 0 {
			alice += c
		} else {
			bob += c
		}
	}

	fmt.Println(alice - bob)
}
