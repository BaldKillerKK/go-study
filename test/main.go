package main

import (
	"fmt"
)

func main() {

	fn := Colsure("田波")
	fn()
}

func Colsure(name string) func() string {
	fmt.Printf("Colsure : %p\n", &name)
	return func() string {
		fmt.Printf("Fn : %p\n", &name)
		return name
	}
}
