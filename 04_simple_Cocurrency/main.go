package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters () {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {

	// For better understanding of the following code(Cocurrency in golang), comment out the two following line:
	// and see the result.
	//there will be no cocurrency in the program second program will run after finishig the first program.
	// but in when you run them with this syntax (Cocurrency consept), you will see they run simultaneously and concurrently.
	
	go printNumbers()
	go printLetters()

	
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("\nPrinting from the main goroutine")
}