package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("maxBoxesInWarehouse([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}): %v\n", maxBoxesInWarehouse([]int{4, 3, 4, 1}, []int{5, 3, 3, 4, 1}))
}

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	sort.Ints(boxes)
	n := len(warehouse)
	leftmin := make([]int, n)
	min := warehouse[0]
	for i := 0; i < n; i++ {
		if min > warehouse[i] {
			min = warehouse[i]
		}
		leftmin[i] = min
	}
	p := 0
	n1 := len(boxes)
	for i := n - 1; i >= 0 && p < n1; i-- {
		if leftmin[i] >= boxes[p] {
			p++
		}
	}
	return p

}
