package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
	}
	colors["green"] = "#4bf745"
	//delete(colors, "red")
	printMap(colors)
}

func printMap (c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}