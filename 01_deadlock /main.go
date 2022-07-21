package main

import "fmt"

// This program panics because there is no goroutine.
// outside of the main interacive with 'c' channel:
// 
// fatal error: all goroutines are asleep - deadlock!
func main() {
	var c chan int
	c = make(chan int)
	// two line of above can define as:
	// c := make(chan int)

	c <- 10

	v := <- c

	fmt.Println("Recieved: ", v)
	

}