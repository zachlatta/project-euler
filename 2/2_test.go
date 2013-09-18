package main

import "testing"

func TestEvenFibonacciNumbersUnder(t *testing.T) {
  actual := evenFibonacciNumbersUnder(4000000)
  expected := 4613732
  if actual != expected {
    t.Error("Got:", actual, "Expected:", expected)
  }
}

func BenchmarkEvenFibonacciNumbersUnder(b *testing.B) {
  for i := 0; i < b.N; i++ {
    evenFibonacciNumbersUnder(4000000)
  }
}
