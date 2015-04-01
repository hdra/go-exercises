package main

import (
	"fmt"
	"time"
)

func main() {
	jobcount := 100
	workercount := 20
	jobs := make(chan int)
	results := make(chan int)

	go func() {
		for i := 0; i < jobcount; i++ {
			jobs <- i
		}
	}()

	for i := 1; i < workercount; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < jobcount; i++ {
		<-results
	}
}

func worker(id int, jobs, result chan int) {
	for jobid := range jobs {
		fmt.Printf("Worker %v : job :%v\n", id, jobid)
		time.Sleep(1 * time.Second)

		result <- jobid
	}
}
