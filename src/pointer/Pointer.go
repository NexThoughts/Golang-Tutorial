package main

import "fmt"



//Agenda
//1. Creating pointer
//2. Dereferencing pointer
//3. The new function
//4. Working with nil
//5. Types with internal pointers

func main() {

	example1()
	example2()
	example3()
	fmt.Println("main end")

}

func example1() {
	var a int = 42
	var b  *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)

	a = 27
	fmt.Println(a, *b)

	*b = 42
	fmt.Println(a, *b)

}

func example2() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Println(a, b, c)

}
func example3() {
	var st student = student{
		"john", 42,
	}

	fmt.Println(st)

	st1 := st
	st1.name = "steve"
	fmt.Println(st)

}

type student struct {
	name   string
	rollNo int
}

//Summary
//1. creating pointers
//	pointer types use an asterisk(*) as a prefix to type pointed to
//		*int - a pointer to an integer
//	Use the addressof operator (&) to get address of a variable
//2 Dereferencing Pointers
//	Dereference a pointer by preceding with an asterisk(*)
//	Complex types (eg: structs) are automatically dereferenced
//3 Type with internal pointers
//	All Assignment operations in go  are copy operations.
//	Slices and maps comtains internal pointers,  so copies point to same underlying data

