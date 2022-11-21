package main

import "fmt"

func main() {
	fmt.Println(test("asdfasdfasdflasdfasdfasdfe", "ale"))
	forloopstring("1234")
}

func findLongestWord(s string, dictionary []string) string {
	max := ""
	for _, word := range dictionary {
		if test(s, word) {
			if len(max) < len(word) || (len(max) == len(word) && max > word) {
				max = word
			}
		}
	}
	return max
}

func test(s string, word string) bool {
	p := 0
	for i := 0; i < len(s) && p < len(word); i++ {
		if s[i] == word[p] {
			p++
		}
	}
	return p == len(word)
}
func forloopstring(s string) {
	for i := range s {
		fmt.Println(s[i], i)
	}
}
