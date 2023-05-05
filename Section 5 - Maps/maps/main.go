package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#00FF00",
	}

	// Other ways of declaring/initializing maps
	//  - These will actually be empty maps
	// var colors map[string]string
	// colors := make(map[string]string)

	// Updating maps
	colors["white"] = "#FFFFFF"

	// Deleting from maps
	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for the given %s is %s\n", color, hex)
	}
}