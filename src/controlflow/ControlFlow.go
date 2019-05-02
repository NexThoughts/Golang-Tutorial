package main

import "fmt"

//Agenda
//1 Defer   -> it executes after function exists -> LIFO
//2 Panic
//3 Recover

func main() {

	//deferExample()
	//panicExample1()
	panicExample2()

	fmt.Println("main end")

}

func deferExample() {

	fmt.Println("first")
	defer fmt.Println("middle")
	fmt.Println("last")

	a := "start"
	defer fmt.Println(a)
	a = "end"

}

func panicExample1() {

	defer fmt.Println("this was deferred")

	var a, b int = 1, 0
	c := a / b
	fmt.Println(c)
}
func panicExample2() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("Error : ", error)

			//panic(error)     // If you dont want to continue
		}
	}()
	fmt.Println("first")
	panic("something bad happened")
	fmt.Println("last")
}

//Summary
//1 Defer
//	used to delay the execution of a statement until function exits.
//	Useful to close the resourses we use
//	Run in LIFO order
//	arguments are evaluated at time defer is executed, not at the time of called function execution
//2 Panic
//	Occur when program cannot continue at all.
//		Use for unrecoverable events - cannot obtain TCP port for web server
//	Function will stop executing
//		Deferred function will still fire
//	If nothing handles panic, program will exit
//3 Recover
//	Used to recover from panics
//	Only useful in deferred functions
//	current function will not attempt to continue, but higher function in call stack will
