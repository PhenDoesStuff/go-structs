package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var stephen person



	fmt.Println(stephen)
	fmt.Printf("%+v", stephen)
}