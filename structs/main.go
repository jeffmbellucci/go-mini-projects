package main

import "fmt"

type person struct {
	first string
	last string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}


func main() {
  jeff := person{
		first: "Jeff",
		last: "BeLlucci",
		contactInfo: contactInfo{
			email: "jeffmbellucci@gmail.com",
			zipCode: 95945,
		},
	}

	jeff.updateName("Jeffrey")
	jeff.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.first = newFirstName
}
