package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// boolean
	// integer
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for  int32
	// float32 float64
	// complex64 complex128

	// creating variable using var keyword
	// var name  = "Ernest"
	var age int32 = 37
	isCool := true
	isCool = false

	// shorthand var decalration

	// name := "Brad"
	// email := "ernest.com"
	size := 1.3

	name, email := "Brad", "ernest@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", email)
}
