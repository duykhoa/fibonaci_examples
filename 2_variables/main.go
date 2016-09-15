package main

import "fmt"

/*
 3 variables version
  for i:=2; i < n; i++ {
    c = a + b
    a = b
    b = c
  }
*/

func fib(n int) int {
  a,b := 1,1

  for i:=2; i < n; i++ {
    a, b = b, a + b
  }

  return b
}

func main() {
  fmt.Println(fib(10))
}
