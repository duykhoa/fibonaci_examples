package main

import "fmt"

func fib(n int) int {
  fibs := make([]int, n)
  fibs[0] = 1; fibs[1] = 1

  for i := 2; i < n; i++ {
    fibs[i] = fibs[i-1] + fibs[i-2]
  }

  return fibs[n - 1]
}

func main() {
  fmt.Println(fib(10))
}
