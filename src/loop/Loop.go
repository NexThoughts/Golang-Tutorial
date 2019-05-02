package main

import "fmt"




//For Statement
//	simple loops
//	exiting early
//	looping through collections


func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
	fmt.Println("********************************")

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("********************************")

	s := []int{1, 2, 3}
	//s := "hello"
	for k, v := range s {                       //for range loop
		fmt.Println(k, v)
		//fmt.Println(k, string(v))

	}

}


//Summary
//1 Simple loops
//	for initializer; test; incrementer {}
//	for test {}
//	for {}
//2 exiting early
//	break
//	continue
//	labels
//3 looping over collections
//	array,slices, maps, string ,channels
//	for k,v := range collection {}
//
//



