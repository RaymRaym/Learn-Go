package main

import "fmt"

func main() {

	colors := map[string]int{
		"Red":   0,
		"Green": 1,
		"Blue":  2,
	}
	fmt.Printf("colors: %v\n", colors)

	delete(colors, "Green")

	fmt.Printf("colors: %v\n", colors)

}
