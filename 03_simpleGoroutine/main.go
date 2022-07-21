package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Printing from goroutine")
}

func main() {
	go print()
	fmt.Println(time.Now(), "Taking a nap and watting for complating goroutine print")
	
	// If we comment out the following line, the program will not wait for the goroutine to complete and will exit.
	// and the program will not print "Printing from goroutine(main)"
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now(),"Printing from main")
}