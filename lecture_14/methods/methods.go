package main

import "fmt"

//Human contains the information of a human
type Human interface {
	//Speak()
}

//Book contains the information of a book
type Book struct {
	Name  string
	Price int
}

//Doctor contains the information of a doctor
type Doctor struct {
	Name   string
	Degree string
	Age    int
}

//Engineer contains the information of a engineer
type Engineer struct {
	Name   string
	Degree string
	Age    int
}

//Speak prints a corresponding message
func (d Doctor) Speak() {
	fmt.Println(d.Name, "can speak")
}

//Surgery prints a corresponding message
func (d Doctor) Surgery() {
	fmt.Println(d.Name, "can surgery")
}

//Speak prints a corresponding message
func (e Engineer) Speak() {
	fmt.Println(e.Name, "can speak")
}

//Programming prints a corresponding message
func (e Engineer) Programming() {
	fmt.Println(e.Name, "can code")
}

func main() {
	d := Doctor{"Nahid", "DMF", 28}
	d.Speak()

	e := Engineer{Name: "Hasan", Degree: "BSc.", Age: 28}
	e.Speak()

	humans := [2]Human{d, e}
	fmt.Println(humans)

	for _, v := range humans {
		fmt.Println(v)

		switch v.(type) {
		case Doctor:
			dd := v.(Doctor)
			fmt.Println(dd.Degree)
			dd.Surgery()
		case Engineer:
			fmt.Println(v.(Engineer).Degree)
			v.(Engineer).Programming()
		}
	}
}
