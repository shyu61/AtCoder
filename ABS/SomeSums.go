package main

import (
  "fmt"
)

func main() {
  var n, a, b, ans int
  fmt.Scan(&n, &a, &b)

  for i := 0; i <= n; i++ {
    sum := 0
    s := i
    for s > 0 {
      sum += s % 10
      s /= 10
    }

    if (sum >= a && sum <= b) {
      ans += i
    }
  }

  fmt.Println(ans)
}
