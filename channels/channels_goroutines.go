package main

import (
	"fmt"
	"log"
)

func fibonacci(max int, channel chan int) {
	fibo := make([]int, max)
	fibo[0] = 0
	fibo[1] = 1

	channel <- fibo[0]
	channel <- fibo[1]

	for i := 2; i < max; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]

		channel <- fibo[i]
	}

	close(channel)
}

func main() {
	var max int

	fmt.Println("----- Fibonacci sequence printing by using channels and goroutines -----")
	fmt.Println("Enter the number till which you want the sequence to be printed: ")
	n, err := fmt.Scanln(&max)
	if err != nil || n <= 0 {
		log.Fatal("Error while reading input from the user !")
	}

	channel := make(chan int)
	go fibonacci(max, channel)

	for message := range channel {
		fmt.Println(message)
	}
}
