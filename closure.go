package main

import "fmt"

func main() {
	printer := makePrint()
	printer()
}

func makePrint() func() {
	return func() {
		fmt.Println("Hello world!")
	}
}
