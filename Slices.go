package main

import "fmt"

func main() {
	var nums [5]int

	fmt.Println("Empty:", nums)

	nums[4] = 100
	fmt.Println("Set: ", nums)
	fmt.Println("Get: ", nums[4])

	ints := []int{1, 2, 3, 4, 5}
	fmt.Println("new slice:", ints)

	ints = append(ints, 6)
	fmt.Println("appended", ints)

	fmt.Println("0-2", ints[:2])
	fmt.Println("4-4", ints[2:4])
	fmt.Println("4-6", ints[4:])

	fmt.Println("sum", sum(ints))
}

func sum(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}

	return sum
}
