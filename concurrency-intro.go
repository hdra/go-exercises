package main

import (
	"fmt"
	"time"
)

func main() {
	go say("go")
	go say("for it")

	var in string
	fmt.Scanln(&in)
}

func say(msg string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(msg)
	}
}
