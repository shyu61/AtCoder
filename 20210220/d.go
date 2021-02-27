package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer _w.Flush()
	X, M := ReadString(), ReadInt()
	if len(X) == 1 {
		v, _ := strconv.Atoi(X)
		if uint64(v) <= M {
			println(1)
		} else {
			println(0)
		}
		return
	}

	var s uint64
	for _, x := range X {
		s = Max(s, uint64(x)-'0'+1)
	}
	var l uint64
	r := M + 1
	for r-l > 1 {
		// c := (l + r) / 2
		c := l + ((r - l) / 2)
		if check(s, X, c, M) {
			l = c
		} else {
			r = c
		}
	}
	if l < s {
		println(0)
	} else {
		println(l - s + 1)
	}
}

func check(s uint64, X string, c, M uint64) bool {
	if c < s {
		return true
	}
	if c > M {
		return false
	}

	base := uint64(1)
	for i := len(X) - 1; i >= 0; i-- {
		if M == 0 {
			return false
		}
		d := uint64(X[i] - '0')
		v := d * base
		if v > M {
			return false
		}
		M -= v
		if i >= 1 && base > M/c {
			return false
		}
		base *= c
	}
	return true
}

func Max(xs ...uint64) uint64 {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}
	return max
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func print(a ...interface{})            { fmt.Fprint(_w, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(_w, f, a...) }
func println(a ...interface{})          { fmt.Fprintln(_w, a...) }
func printInts(a ...int)                { s := fmt.Sprint(a); println(s[1 : len(s)-1]) }
func Scan(a ...interface{}) {
	if _, err := fmt.Fscan(_r, a...); err != nil {
		panic(err)
	}
}
func ReadString() (s string) { Scan(&s); return }
func ReadInt() (i uint64)    { Scan(&i); return }
