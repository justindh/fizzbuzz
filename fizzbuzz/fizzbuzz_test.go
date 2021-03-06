package fizzbuzz

import (
	"testing"
)

func TestFizzbuzzOne(t *testing.T) {
	if Fizzbuzz(1) != "1" {
		t.Fatal("Testing 1 did not return 1")
	}
}

func TestFizzbuzzThree(t *testing.T) {
	if Fizzbuzz(3) != "Fizz" {
		t.Fatal("Testing 3 did not return Fizz")
	}
}
func TestFizzbuzzFive(t *testing.T) {
	if Fizzbuzz(5) != "Buzz" {
		t.Fatal("Testing 5 did not return Buzz")
	}
}
func TestFizzbuzzFifteen(t *testing.T) {
	if Fizzbuzz(15) != "FizzBuzz" {
		t.Fatal("Testing 15 did not return FizzBuzz")
	}
}

func BenchmarkStringCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := Fizzbuzz(1)
		_ = result
	}

}

func Benchmark3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := Fizzbuzz(3)
		_ = result
	}
}

func Benchmark5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := Fizzbuzz(5)
		_ = result
	}
}
func Benchmark15(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result := Fizzbuzz(15)
		_ = result
	}
}
