package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

type passenger struct {
	person
	destination string
}

func (p passenger) speak() {
	fmt.Println(p.fname, "says: Good morning. I,", p.fname, p.lname, "want to go to", p.destination)
}

type driver struct {
	person
	car string
}

func (d driver) speak() {
	fmt.Println(d.lname, "asks: Were do you want me,", d.fname, "to bring you in his", d.car, "?")
}

type human interface {
	speak()
}

func humanspeak(h human) {
	h.speak()
}

func main() {
	p1 := passenger{person{"Lete", "Hoebeke", 57}, "1040 East Putnam in Riverside."}
	d1 := driver{person{"Eval", "Knieval", 29}, "Chevrolet"}

	humanspeak(d1)
	humanspeak(p1)
}
