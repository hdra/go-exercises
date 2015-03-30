package main

import "fmt"

type Animal interface {
	Pet()
	Name() string
}

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("prr")
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("woof")
}

func (d Dog) Name() string {
	return d.name
}

func Compliment(a Animal) {
	fmt.Println("Great job", a.Name())
	a.Pet()
}

func main() {
	c := Cat{"john"}

	c.Pet()

	d := Dog{"bob"}
	d.Pet()

	Compliment(c)
	Compliment(d)
}
