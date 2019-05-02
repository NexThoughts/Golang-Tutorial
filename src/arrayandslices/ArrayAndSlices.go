package main

//
//Arrays
//	creation
//	built in functions
//	working with arrays
//Slices
//	creation
//	built in functions
//	working with arrays
import "fmt"

func main() {
	grades := [3]int{97, 85, 60}
	fmt.Printf("%v , %T", grades, grades)
	fmt.Println()

	grades2 := [...]int{97, 85, 60}
	fmt.Printf("%v , %T", grades2, grades2)
	fmt.Println()

	students := [3]string{}
	students[0] = "a1"
	fmt.Printf("%v , %T", students, students)
	fmt.Println()
	fmt.Println(len(students))

	grades3 := grades
	grades3[0] = 99
	fmt.Printf("%v , %T", grades, grades)
	fmt.Println()
	fmt.Printf("%v , %T", grades3, grades3)
	fmt.Println()

	grades4 := &grades
	grades4[0] = 99
	fmt.Printf("%v , %T", grades, grades)
	fmt.Println()
	fmt.Printf("%v , %T", grades4, grades4)
	fmt.Println()










	//slice
	slice1 := []int{97, 85, 60}
	fmt.Printf("%v , %T", slice1, slice1)
	fmt.Println()

	slice2 := slice1   //it is pointing to same underlying data
	slice2[0] = 100

	fmt.Printf("%v , %T", slice1, slice1)
	fmt.Println()
	fmt.Printf("%v , %T", slice2, slice2)
	fmt.Println()







	//using make function
	a := make([]int, 0)     //type ,length

	a = append(a, 2,3,4,5,6,7,8)         //slice,value
	fmt.Printf("%v , %T", a, a)
	fmt.Println()
	fmt.Println(len(a))
	fmt.Println(cap(a))

}

//
//Summary
//1 Array
//	collection of items with same type
//	fixed size
//	declaration style
//		a:=[3]int{1,2,3}
//		a:=[...]int{1,2,3}
//		var a [3]int
//	access via zero based index
//		a:=[3]int {1,3,5}   //a[1] == 3
//
//	len function returns size of array
//	copies refer to different underlying data
//
//2 Slices
//	creation style
//	len function returns return length of slice
//	cap function returns length of underlying array
//	append function to add elements to slice
//		may cause expensive copy operation if underlying array is  too small
//
//	copies refer to same underlying array
