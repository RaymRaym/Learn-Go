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

	fmt.Println(123123)

	// 被append的数组必须是[]，而非[3]or[...]
	// append 必须要有返回值
	var y = []int{1, 2, 3}
	y = append(y, 4)

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

}
