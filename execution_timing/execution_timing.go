package main

import (
	"fmt"
	"log"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s function took %v \n", name, time.Since(start).Microseconds())
	}
}

func fibIterative(max int) {
	defer timer("iterative")()
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
}

func fibRecursive(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return fibRecursive(number-1) + fibRecursive(number-2)
}

func main() {
	var max int

	fmt.Println("----- Fibonacci sequence calculation benchmarking -----")
	fmt.Println("Enter the number till which you want the sequence to be calculated recursively and iteratively: ")
	n, err := fmt.Scanln(&max)
	if err != nil || n <= 0 {
		log.Fatal("Error while reading input from the user !")
	}

	fibIterative(max)
	defer timer("recursive")()
	fibRecursive(max)
}
