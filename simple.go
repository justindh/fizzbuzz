package main

import (
	"flag"
	"fmt"

	"github.com/justindh/fizzbuzz/fizzbuzz"
)

func main() {
	amount := flag.Int("amount", 16, "Amount of Numbers to FizzBuzz test")
	flag.Parse()

	for i := 1; i <= *amount; i++ {
		fmt.Println(fizzbuzz.Fizzbuzz(i))
	}
}
