package main

import (
	"fmt"
)

func main() {
	var a = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)
	fmt.Printf("%o \n", a)
	fmt.Printf("%x \n", a)

	var b = 0777
	fmt.Printf("%d \n", b)
	fmt.Printf("%b \n", b)
	fmt.Printf("%o \n", b)
	fmt.Printf("%x \n", b)
	c := 0x1234
	fmt.Printf("%d \n", c)
	fmt.Printf("%b \n", c)
	fmt.Printf("%o \n", c)
	fmt.Printf("%x \n", c)

	d := int8(9)
	fmt.Printf("%d \n", d)
	fmt.Printf("%b \n", d)
	fmt.Printf("%o \n", d)
	fmt.Printf("%x \n", d)
}
