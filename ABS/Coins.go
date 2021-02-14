package main

import (
  "fmt"
)

func main() {
  var a, b, c, x int

  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Scan(&c)
  fmt.Scan(&x)

  var count int
  for i := 0; i <= a; i++ {
    if 500 * i > x {
      break
    }

    for j := 0; j <= b; j++ {
      if 500 * i + 100 * j > x {
        break
      }

      for k := 0; k <= c; k++ {
        if 500 * i + 100 * j + 50 * k > x {
          break
        }

        if 500 * i + 100 * j + 50 * k == x {
          count++
        }
      }
    }
  }

  fmt.Println(count)
}
