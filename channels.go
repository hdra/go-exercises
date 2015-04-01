package main

import "fmt"

func main() {
	c := make(chan string, 4)
	c <- "Hello"
	c <- "world"
	c <- "world"

	v := <-c
	fmt.Println(v)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
