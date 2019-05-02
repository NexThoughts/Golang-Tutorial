package main

//* declaration
//* redeclaration and shadowing
//* visibility rules                -> variable visibility i.e lower case and upper case
//* type conversions
//* naming conventions
//   ALL declaration must be used in application


import (
	"fmt"
	"strconv"
)

func main() {
	//3 ways to declare variables

	//1.
	var i int
	i = 42
	fmt.Println(i)


	//2.
	var j int = 42
	fmt.Println(j)



	//3.
	k := 42
	fmt.Println(k)

	fmt.Printf("%v , %T", j, j) //provide formatting string
	fmt.Println()

	var l float32 = 42.55
	fmt.Printf("%v , %T", l, l) //provide formatting string
	fmt.Println()

	fmt.Printf("%v , %T", k, k) //provide formatting string
	fmt.Println()

	k = 42.0
	fmt.Println(k)

	fmt.Printf("%v , %T", j, j) //provide formatting string
	fmt.Println()




	//type conversions


	var m int
	//m = l
	m = int(l)

	fmt.Printf("%v , %T", m, m) //provide formatting string
	fmt.Println()

	var n string
	//m = l
	n = string(m)

	fmt.Printf("%v , %T", n, n) //provide formatting string
	fmt.Println()


	n = strconv.Itoa(m)

	fmt.Printf("%v , %T", n, n) //provide formatting string
	fmt.Println()


}
//Grouping
var
(
	example string = "Variables"  // Global Variable
)
var
(
	counter int = 0  // Global Variable
)







//Summary:
//1 variable declaration
//   	var foo int
//	var foo int =42
//	foo:=42
//
//2 Cant redeclare variables but can shadow them
//3 all variables must be used
//4 visibility
//	lower case first letter for package scope
//	upper case first letter to export
//	no private scope
//5 naming convention
//	pascal or camelCase
//		capitalize acronyms(HTTP,URL)
//	As short as reasonable
//		longer names for longer lives
//6 type conversions
//	destinationType(variable)
//	use strconv for strings