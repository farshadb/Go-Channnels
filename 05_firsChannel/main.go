package main

import "fmt"

func Print(c chan bool) {
	fmt.Println("Printing form goroutine")
	c <- true
}

func main() {

	c :=make(chan bool)
	// this line will print c addres in the memory
	fmt.Println(c)
	// Below line will return fatal error  because deadlock - no gorutine has been started.
	//a := <-c
	//fmt.Println(a)
	// Sending channel c as parameter to Print Goroutine 

	go Print(c)
	<- c
    fmt.Println("Printing form main(main also is a goroutine).")

	
}