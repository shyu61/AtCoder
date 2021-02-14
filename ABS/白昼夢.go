package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func scanLine() string {
	sc.Buffer(nil, math.MaxInt64)
	sc.Scan()
	return sc.Text()
}

func main() {
	s := scanLine()

	l := len(s)
	var msg string
	for {
		if msg != "" {
			fmt.Println(msg)
			break
		}

		l = len(s)
		switch {
		case l == 0:
			msg = "YES"
		case l < 5:
			msg = "NO"
		case l < 10:
			switch {
			case l == 5 && s[:5] == "dream":
				s = s[5:]
			case l == 5 && s[:5] == "erase":
				s = s[5:]
			case l == 6 && s[:6] == "eraser":
				s = s[6:]
			case l == 7 && s[:7] == "dreamer":
				s = s[7:]
			default:
				msg = "NO"
			}
		case l >= 10:
			switch {
			case s[:5] == "dream":
				switch {
				case s[5:10] == "erase":
					s = s[5:]
				case s[5:7] == "er":
					s = s[7:]
				default:
					s = s[5:]
				}
			case s[:6] == "eraser":
				s = s[6:]
			case s[:5] == "erase":
				s = s[5:]
			default:
				msg = "NO"
			}
		}
	}
}
