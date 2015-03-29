package main

import "fmt"

func main() {
	age := make(map[string]int)

	age["jeremy"] = 12
	age["josh"] = 13
	age["john"] = 15

	fmt.Println(age)

	fmt.Println("Jeremy's age", age["jeremy"])

	delete(age, "jeremy")
	fmt.Println(age)

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(m)

	for n, a := range m {
		fmt.Printf("%v is %v years old", n, a)
	}

	fmt.Println(keys(m))
}

func keys(m map[string]int) []string {
	keys := []string{}

	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}
