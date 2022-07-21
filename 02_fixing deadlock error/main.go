package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	// defining a goroutine
	go func() {
		fmt.Println(time.Now(), "Taking a nap.")
		time.Sleep(time.Second * 2)
		c <- "Hello" 
	}()

	fmt.Println(time.Now(),"Wating for a massage ... ")

	v := <- c
	fmt.Println(time.Now(), "Massage received: ", v)

}