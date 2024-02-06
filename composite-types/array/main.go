package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 初始化
	arr01 := [3]int{1, 2, 3}
	// 编译器自动计算大小
	arr02 := [...]int{1, 2, 3, 4}
	// 选择性初始化 其他赋类型的零值
	arr03 := [3]string{1: "b"}

	// 遍历
	for i := 0; i < len(arr01); i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------------")

	for _, val := range arr02 {
		fmt.Println(val)
	}
	fmt.Println("---------------------------")

	// 计算数组长度和数组大小
	fmt.Println(len(arr03))
	fmt.Println(unsafe.Sizeof(arr03))
}
