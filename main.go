package main

import "fmt"

type contactInfo struct {
	email   string
	zioCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Way of creating a struct
	// stephen := person{firstName: "Stephen", lastName: "Montague"}

	// Another way of creating a struct:
	// var stephen person

	// stephen.firstName = "Stephen"
	// stephen.lastName = "Montague"

	// fmt.Println(stephen)

	// --------------------------------------

	// New Person with a Contact:
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contactInfo{
			email:   "jim.halpert@dundermiflin.com",
			zioCode: 18501,
		},
	}

	fmt.Printf("%+v", jim)
}