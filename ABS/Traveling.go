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
	n := scanInt()

  var tdiff, xydiff, prevT, prevX, prevY int
	for i := 0; i < n; i++ {
		t, x, y := scanInt(), scanInt(), scanInt()
		tdiff = t - prevT
		xydiff = (x + y) - (prevX + prevY)

		switch {
		case (xydiff > tdiff) || (xydiff < 0 && -xydiff > tdiff):
			fmt.Println("No")
			return
		case (xydiff + tdiff) %2 == 0 && (tdiff - xydiff) %2 == 0:
			prevT, prevX, prevY = t, x, y
		default:
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
