package main

//* Naming conventions
//* Typed and untyped constants
//* Enumerated constants and,
//* Enumeration expressions
import "fmt"

const (
	//* Enumerated constants and,

	a = iota
	b = iota
	c
)

func aaa() int {
	return 30 + 50
}
func main() {
	const myConst int = 0  // NOT MY_CONST because it will be expored to other packages
	fmt.Printf("%v , %T", myConst, myConst)
	fmt.Println()
	var c int = aaa()    // shadowed
	fmt.Println(a, b, c, )

}




//Summary
//1 Immutable but can be shadowed
//2 Replaced at compile time
//	value must be calculated at compile time
//3 named like variable
//	pascalCase for exported constants
//	camelcase for internal constants