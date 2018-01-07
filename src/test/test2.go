package main

import (
	"fmt"
)

type person struct {
	name  string
	color string
}

type car struct {
	brand string
	color string
}

func (p person) mycolor() string {
	return p.color
}
func (p person) myname() string {
	return p.name
}

func (c car) mycolor() string {
	return c.color
}
func (c car) mybrand() string {
	return c.brand
}

type myobj interface {
	mycolor() string
}

func printColor(cc myobj) {
	fmt.Println("my color is:", cc.mycolor())
}

func main() {
	p1 := person{"Yves", "white"}
	c1 := car{"Ford", "red"}

	fmt.Println(p1.myname(), c1.mybrand())
	fmt.Println(p1.mycolor(), c1.mycolor())

	printColor(p1)
	printColor(c1)
}
