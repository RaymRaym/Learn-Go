package main

import "fmt"

// you don't have to use these variables, which are outside of the main functions
var n100 = 100
var (
	n101 = 101
	n102 = 102
)

// string is a primitive type of GO
func main() {
	var age int
	age = 18
	fmt.Println(age)

	// if you declared a variable but don't use -> compile failed
	var name string = "123"
	//fmt.Println(name)
	fmt.Println(name)

	// default value is 0
	var num int
	fmt.Println(num)

	//
	var num2 = 10
	fmt.Println(num2)

	//
	num3 := 3
	fmt.Println(num3)

	//
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	//
	var n4, s, n5 = 10, "sss", 11
	fmt.Println(n4)
	fmt.Println(s)
	fmt.Println(n5)

	//
	n6, ss := 17, "ss"
	fmt.Println(n6)
	fmt.Printf("ss: %v\n", ss)

	// error
	// var a int8 = 234
	// println(a)

	var b int8 = 15
	println(b)

}
