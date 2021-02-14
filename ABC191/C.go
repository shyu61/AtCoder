package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var intSc = bufio.NewScanner(os.Stdin)
var RowSc = bufio.NewScanner(os.Stdin)

func readInt() int {
	intSc.Scan()
	i, _ := strconv.Atoi(intSc.Text())
	return i
}

func readRow() string {
	RowSc.Scan()
	return RowSc.Text()
}

type coordinate struct {
	x int
	y int
}

func main() {
	intSc.Split(bufio.ScanWords)
	h, w := readInt(), readInt()

	blacks := make([]coordinate, 0)
	// vertext := make([]coordinate, 0)

	for i := 1; i <= h; i++ {
		row := readRow()

		for j := 1; j <= w; j++ {
			if row[j-1:j] == "#" {
				blacks = append(blacks, coordinate{x: j, y: i})
			}
		}
	}
	
	var vertext int
	for _, k := range blacks {
		surround := make([]coordinate, 0)
		surround = append(surround, coordinate{x: k.x+1, y: k.y})
		surround = append(surround, coordinate{x: k.x-1, y: k.y})
		surround = append(surround, coordinate{x: k.x, y: k.y+1})
		surround = append(surround, coordinate{x: k.x, y: k.y-1})

		var count int
		for _, l := range surround {
			for _, m := range blacks {
				if l == m {
					count++
				}
			}
		}
		fmt.Println(count)
		if count == 2 || count == 4 {
			vertext++
		}
	}

	fmt.Println(vertext)
}

// 5 5
// .....
// .###.
// .###.
// .###.
// .....
