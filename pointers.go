package main

import "fmt"

func main() {
	x, y := 10, 20
	fmt.Println(x, y)

	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
