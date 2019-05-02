package main


//Struct
// 	what are they?
//	creating
//	naming convention
//	embedding

import "fmt"

func main() {

	switch 6 {
	case 4, 5:
		fmt.Println("four and five")
	case 6:
		fmt.Println("six")

	default:
		fmt.Println("greater than six")
	}


	//fallthrough                       //fallthrough is logic less


	//i := 5
	//switch  {
	//case i == 4, i == 5:
	//	fmt.Println("four and five")
	//break
	//	fmt.Println("four and five")
	//
	//case i == 6:
	//	fmt.Println("six")
	//
	//default:
	//	fmt.Println("greater than six")
	//}
}






