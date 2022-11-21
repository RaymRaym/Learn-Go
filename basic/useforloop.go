package main

import "fmt"

func main() {

	s := "qwer"
	for i := range s {
		fmt.Println(i)
	}
	fmt.Println("---------------------")

	for i, ch := range s {
		fmt.Println(i, ch)
	}
	fmt.Println("---------------------")

	r := []rune(s)
	for _, ch := range r {
		fmt.Println(ch)
	}

}
