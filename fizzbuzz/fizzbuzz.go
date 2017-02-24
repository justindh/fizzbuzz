package fizzbuzz

import (
	"fmt"
)

// Fizzbuzz accepts an Int and returns the appropriate FizzBuzz Test compliant string
func Fizzbuzz(i int) string {
	var result string
	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	if len(result) == 0 {
		result = fmt.Sprintf("%v", i)
	}
	return result
}
