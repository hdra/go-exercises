package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("Running..")
		time.Sleep(250 * time.Millisecond)
		done <- true
	}()

	go func() {
		fmt.Println("Running..")
		time.Sleep(400 * time.Millisecond)
		done <- true
	}()

	waiton(done)
	waiton(done)

}

func waiton(done chan bool) {
	select {
	case <-done:
		fmt.Println("Done")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout")
	}
}
