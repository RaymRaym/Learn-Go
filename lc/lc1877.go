package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("minPairSum([]int{3, 5, 2, 3}): %v\n", minPairSum([]int{3, 5, 2, 3}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	ans := 0
	for left < right {
		if ans < nums[left]+nums[right] {
			ans = nums[left] + nums[right]
		}
		left++
		right--
	}
	return ans
}
