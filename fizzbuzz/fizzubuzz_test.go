package fizzbuzz_test

import (
	"fmt"
	"testing"
	"todo/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	cases := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		6:  "Fizz",
		9:  "Fizz",
		5:  "Buzz",
		10: "Buzz",
		15: "FizzBuzz",
		30: "FizzBuzz",
	}

	for given, want := range cases {
		t.Run(fmt.Sprintf("given %d want %q", given, want), func(t *testing.T) {
			get := fizzbuzz.Say(given)
			if want != get {
				t.Errorf("%q %q", want, get)
			}
		})
	}
}

type stubIntn struct{}

func (stubIntn) Intn(int) int {
	return 2
}

type IntnFunc func(int) int

func (f IntnFunc) Intn(n int) int {
	return f(n)
}

func TestRandomFizzBuzz(t *testing.T) {
	want := "Buzz"

	get := fizzbuzz.RandomFizzBuzz(IntnFunc(func(int) int { return 4 }))

	if want != get {
		t.Error()
	}
}
