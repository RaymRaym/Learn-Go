package main

// 此二分搜索适用于以下场景：
// 给出一个sorted array ： nums 和一个target
// 输出target的上下边界 如果target的没有上界或下界，输出-1.如果有输出上界和下界的索引
func main() {
	var nums = []int{0, 2, 4}
	left := 0
	right := len(nums)
	target := 4
	index := -1
	for {
		if left > right {
			break
		}
		mid := left + (right-left)/2
		if mid >= len(nums) {
			break
		}
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			index = mid
			break
		}
	}
	// 没找到这个数
	if index == -1 {
		index = left
		if index == 0 {
			println(-1, index)
		} else if index >= len(nums) {
			println(index-1, -1)
		} else {
			println(index-1, index)
		}
	} else { //找到了这个数
		println(index, index)
	}
	println(index)

}
