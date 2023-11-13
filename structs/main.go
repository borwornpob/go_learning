package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zip: 94000,
		},
	}
	
	// (&jim).updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}