package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x[2:])  //[3, 4]
	fmt.Println(x[:2])  //[1, 2]
	fmt.Println(x[1:3]) // [2, 3]
	fmt.Println(x[:])   // [1, 2, 3, 4]

	// slice之间 会共享内存
	y := x[:2]           //[1, 2]
	z := x[1:]           //[2, 3, 4]
	x[1] = 20            //[1, 20, 3, 4]
	y[0] = 10            //[10, 20]
	z[1] = 30            //[20, 30, 4]
	fmt.Println("x:", x) //[10, 20, 30, 4]
	fmt.Println("y:", y) //[10, 20]
	fmt.Println("z:", z) //[20, 30, 4]

	// 如何避免这个问题 -> 使用 copy 深复制
	y = make([]int, 3, 4) // 开辟一块 len=3,cap=4 的空间
	num := copy(y, x)     // y <- x 虽然cap是4，但是len是3，所以只会复制3个过去
	fmt.Println(y, num)   //[10,20,30] 3
}
