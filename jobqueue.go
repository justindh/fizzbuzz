package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/justindh/fizzbuzz/fizzbuzz"
)

var wg sync.WaitGroup

func main() {
	// get cli args
	amount := flag.Int("amount", 16, "Amount of Numbers to FizzBuzz test")
	workers := flag.Int("workers", 1, "Amount of concurrent workers")
	flag.Parse()

	wg.Add(*amount)

	//  create the job queue
	jobs := make(chan int, *amount)

	// please to collect the results
	results := make(chan string, *amount)

	// spin up workers that will wait for jobs
	for w := 1; w <= *workers; w++ {
		go worker(w, jobs, results)
	}

	// build the job queue
	for j := 1; j <= *amount; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	for i := 1; i <= *amount; i++ {
		fmt.Println(<-results)
		//<-results
	}
}

func worker(workerid int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		//log.Printf("Worker %d starting %d", workerid, j)
		results <- fizzbuzz.Fizzbuzz(j)
		//log.Printf("Worker %d finished %d", workerid, j)
		wg.Done()
	}
}
