package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := "a"
	b := "ghhhh"
	sl := []string{"a"}

	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(sl))

	fmt.Println(len(a))
	fmt.Println(len(b))
}
