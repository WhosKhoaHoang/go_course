package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}
type person2 struct {
	firstName string
	lastName  string
	contactInfo  // Equivalent to contactInfo contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// You don't explicitly need to call updateName() on a person pointer. Instead,
// You can call this function on a variable that directly refers to a person
// struct and Go will behave as though it's calling this function on the pointer
// to that person struct! A handy shortcut.
func (ptrToPerson *person) updateName(newFirstName string) {
	(*ptrToPerson).firstName = newFirstName
}

func main() {
	// // Less ideal way of defining a struct (positional values)
	// //alex := person{ "Tang", "Sanzang" }
	// // More ideal way of defining a struct (keyword values)
	// tang := person{firstName: "Tang", lastName: "Sanzang"}
	// tang.firstName = "Tanglol"
	// tang.lastName = "Sanzanglol"
	// fmt.Println(tang)

	// var sanzo person
	// fmt.Println(sanzo)
	// //fmt.Printf("%s %s", sanzo.firstName, sanzo.lastName)
	// fmt.Printf("%+v", sanzo)

	jim := person{
		firstName: "Jim",
		lastName: "Hopper",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94303,
			},
		}
	jim.updateName("Jimmy")
	jim.print()
}