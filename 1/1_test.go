package main

import "testing"

func Test_sum_of_multiples(t *testing.T) {
	actual := sum_of_multiples()
	expected := 233168
	if actual != expected {
		t.Error("Got:", actual, "Expected:", expected)
	}
}
