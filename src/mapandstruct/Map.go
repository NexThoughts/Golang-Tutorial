package main

//Map
//	what are they?
//	creating
//	manipulation
//


import "fmt"

func main() {
	students := make(map[string]int)
	students = map[string]int{
		"cse":100,
		"ece":50,

	}
	students["civil"] = 80
	fmt.Printf("%v , %T", students, students)
	fmt.Println()
	fmt.Println(students)

	delete(students, "cse")
	fmt.Println(students)

	pop, ok := students["cse"]
	fmt.Println(pop, ok)





	//copies refer to same underlying array

}

//Summary
//Maps
//1 Collections of value types that are accessed via keys.
//2 Created via literals or via make function
//3 Members accessed via [key] syntax
//	myMap["key"]="value"
//4 check for presence with "value, ok" form of result
//5 multiple assignment  refer to same underlying data