package main

import "fmt"

func main() {
	var s string = "Rui Wang 🤯"
	s1 := s[:3]
	s2 := s[4:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s[0])

	// len() 并不是真正长度 是byte的长度
	fmt.Println(len(s)) // 13

	var a rune = 'x'
	var b byte = 'y'
	fmt.Println(a) //120
	fmt.Println(b) //121
	c := string(a)
	d := string(b)
	fmt.Println(c) // x
	fmt.Println(d) // y

	var x int = 65
	var y = string(x)
	fmt.Println(y) // A

	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Printf("bs: %v\n", bs)
	fmt.Printf("rs: %v\n", rs)

	fmt.Println("1" < "2")
	fmt.Println("1" > "2")

}
