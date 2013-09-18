package main

import "fmt"

func main() {
  fmt.Println(evenFibonacciNumbersUnder(4000000))
}

func evenFibonacciNumbersUnder(max int) int {
  a := 1
  c := 1
  b := 1

  var sum int

  for c < max {
    c = a + b
    a = c

    if c % 2 == 0 {
      sum += c
    }

    c = a + b
    b = c

    if c % 2 == 0 {
      sum += c
    }
  }

  return sum
}
