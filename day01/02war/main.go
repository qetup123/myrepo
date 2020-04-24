package main

import (
	"fmt"
)

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "zlz"
	age = 23
	isOk = true

	fmt.Print(isOk)
	fmt.Printf("name:%s", name)
	fmt.Println(age)
}
