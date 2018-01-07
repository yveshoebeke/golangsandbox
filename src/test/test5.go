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

func (p *person) mycolor(col string) bool {
	p.color = col
	return true
}

func (c car) mycolor(col string) bool {
	c.color = col
	return true
}

type myobj interface {
	mycolor() string
}

func printColor(cc myobj) {
	fmt.Println("my color is:", cc.mycolor())
}

func main() {
	p1 := person{"Yves", "white"}
	//	c1 := car{"Ford", "red"}

	fmt.Println(p1)
	fmt.Println(p1.color)

	fmt.Println(p1.mycolor("brown"))
	fmt.Println(p1.color)

	fmt.Println(p1)
	//fmt.Println(c1)

}
