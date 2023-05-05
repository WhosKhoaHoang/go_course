package main

import "fmt"

// No need to define any attributes. We just want to
// create structs that will act as receivers.
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// FOCUS HERE!!!
// Specify an additional type that can perform a function called getGreeting
// that returns a string. We've already specified two types that fulfull this
// criteria (englishBot and spanishBot).
type bot interface {
	getGreeting() string
}
// THINK: b bot WILL CORRESPOND TO SOME VALUE THAT IMPLEMENTS THE bot INTERFACE.
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
// // FUNCTION OVERLOADING IS NOT ALLOWED IN GO!
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// // FUNCTION OVERLOADING IS NOT ALLOWED IN GO!
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// You can omit the parameter assoicated with englishBot if
// you're not gonna use it within the function
func (englishBot) getGreeting() string {
	// Custom logic for generating an English greeting
	return "Sup dude"
}

// You can omit the parameter assoicated with spanishBot if
// you're not gonna use it within the function
func (spanishBot) getGreeting() string {
	// Custom logic for generating a Spanish greeting
	return "HOLA BOI"
}
