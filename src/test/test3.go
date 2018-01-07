package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	email string
}

type driver struct {
	person
	car string
}

func (p person) getfullname() (string, string) {
	p.fname = "Always Yves"
	p.lname = "Always Hoebeke"
	return p.fname, p.lname
}

func (d driver) getfullname() (string, string) {
	return d.fname, d.lname
}

type human interface {
	getfullname() (string, string)
}

func printfullname(h human) (string, string) {
	return h.getfullname()
}

func main() {
	p1 := person{} //{"yves","hoebeke","yh@example.com"}
	d1 := driver{person{"lete", "miranda", "lm@example.com"}, "Hyundai Accent"}
	fmt.Println(p1.fname, d1.fname)
	fmt.Println(printfullname(p1))
}
