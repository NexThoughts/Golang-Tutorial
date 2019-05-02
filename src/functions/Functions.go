package main

import "fmt"



//Agenda
//1. Basic Syntax
//2. Parameters
//3. Return Values
//4. Anonymous function
//5. Function as Types
//6. Methods


type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	g.name = "john"
	fmt.Println(g.greeting, g.name)
}

func main() {
	name := "steve"
	example1(&name)
	fmt.Println(name)

	sum("Sum is ", 1, 2, 3, 4, 5)



	//Anonymous function
	func() {
		fmt.Println("hello")

	}()



	//6. Methods

	methodExample()



	//3. Return Values

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	methodExample()

}

func example1(name *string) {
	*name = "john"
	fmt.Println(*name)

}

func sum(msg string, values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result = result + v
	}
	fmt.Println(msg, result)
	return
}
func divide(a, b float64) (float64, error) {
	if (b == 0.0) {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return (a / b), nil
}






func methodExample() {
	g := greeter{
		greeting:"hello",
		name:"go",
	}
	fmt.Println("******************************")

	g.greet()
	fmt.Println(g.name)

}



//Summary
//
//1 Parameters
//	When pointers are passed in, the function can change the value in the caller
//		this is always true for data of slices and maps because they work as internal pointers
//	Use variadic parameter to send list of same types in
//		Must be last parameter
//		Received as a slice
//2 Return
//	Single return value
//		func example() int
//	Multiple return value
//		func example() (int,error)
//	Can use named return values
//		initializes return variable
//		return using return keyword on its own
//
//3 Anonymous function
//
//4 Methods
//	Function that executes in context of a type
//	Receiver can be value or pointer
//		value receiver gets a copy of type
//		pointer receiver get pointer to type


