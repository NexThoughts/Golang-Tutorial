package main

import "fmt"
//* The boolean type
//* Numeric types
//* Integers
//* Floating point
//* Complex numbers
//* Text types


func main() {
	var n bool = false
	fmt.Printf("%v , %T", n, n)
	fmt.Println()

	var m bool
	fmt.Printf("%v , %T", m, m)
	fmt.Println()

	var o int
	fmt.Printf("%v , %T", o, o)
	fmt.Println()

	var q float32 = 30.1
	fmt.Printf("%v , %T", q, q)
	fmt.Println()

	var r string = "hello"
	fmt.Printf("%v , %T", r, r)
	fmt.Println()

	var s []byte = []byte(r)
	fmt.Printf("%v , %T", s, s)
	fmt.Println()



	p := 'h'
	fmt.Printf("%v , %T", p, p)
	fmt.Println()

}


//summary
//1 boolean types
//	values are true or false
//	zero value is false
//2 numeric types
//	integer
//		Signed integers
//		int type has varying size, but min 32 bits
//		8 bit (int8) through 64 bit (int64)
//
//		UnSigned integers
//		8 bit (uint8) through 32 bit (uint32)
//
//		Arithmetic Operations
//
//	floating point numbers
//		32 and 64 bit versions
//
//	text types
//		Strings
//			UTF 8
//			Immutable
//			Can be concatenated with + operator
//			can be converted to byte