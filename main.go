package main

import "fmt"

type contactInfo struct {
	email   string
	zioCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
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
		contactInfo: contactInfo{
			email:   "jim.halpert@dundermiflin.com",
			zioCode: 18501,
		},
	}

	// & = operator to give the memory address of the value (jim in this case)
	// jimPointer := &jim 
	// Now jimPointer is holding the address of jim
	// jimPointer.updateName("Dwight")
	// jim.printPerson()

	// Shortcut to do the above code:
	jim.updateName("Dwight")
	jim.printPerson()
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

// * = operator that gives the value the memory address is pointing at 
// jimPointer is passed and the * gets the value found at the mem address from jimPointer
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}