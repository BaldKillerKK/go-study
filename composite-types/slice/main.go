package main

import "fmt"

func main() {
	// 初始化方式
	s1 := init01()
	s2 := init02()
	s3 := init03()
	s4 := init04()

	// 遍历方式
	for i := 0; i < len(s1); i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------------------")

	for _, val := range s2 {
		fmt.Println(val)
	}
	fmt.Println("---------------------------------")

	// 使用:读取切片全部
	fmt.Println(s3[:])
	// 读取一部分
	fmt.Println(s4[0:1])
}

// 使用数组初始化
func init01() []string {
	// 初始化数组
	arr := [5]string{"a", "b", "c", "d", "e"}

	// 使用数组初始化切片 左闭右开[)
	slice := arr[1:3]

	return slice
}

// 使用{}方式初始化
func init02() []string {
	// 可以使用{}方式直接初始化切片
	slice := []string{"a", "b", "c", "d", "e"}
	return slice
}

// 使用make方式初始化
func init03() []string {
	// 使用make进行初始化并指定切片大小
	slice := make([]string, 5)
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	slice[3] = "d"
	slice[4] = "e"
	return slice
}

// 使用append方式初始化
func init04() []string {
	var slice []string
	slice = append(slice, "a")
	slice = append(slice, "b")
	return slice
}

func read(slice []string) {

}
