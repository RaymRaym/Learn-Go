package main

import "fmt"

func main() {
	var nums = [3]int{10, 20, 30}
	for i := 0; i < 3; i++ {
		fmt.Println(nums[i])
	}

	var nums2 = [...]int{10, 20, 30}
	fmt.Println(nums == nums2)

	for i := 0; i < 3; i++ {
		fmt.Println(nums2[i])
	}
}
