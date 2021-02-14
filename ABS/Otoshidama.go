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
  var n, d int
  n = scanInt()

  m := make(map[int]bool, n)
  for i := 0; i < n; i++ {
    d = scanInt()

    if !m[d] {
      m[d] = true
    }
  }

  fmt.Println(len(m))
}
