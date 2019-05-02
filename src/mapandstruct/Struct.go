package main


//Struct
// 	what are they?
//	creating
//	naming convention
//	embedding

import "fmt"

type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	student := Student{
		name:"John",
		rollNo:55,
		subjects:[]string{"S1", "S2", "S3"},
	}

	fmt.Printf("%v", student)
	fmt.Println()
	fmt.Println(student.name)

	doctor1 := struct {
		name string
	}{name:"John"}

	fmt.Printf("%v", doctor1)
	fmt.Println()

	doctor2 := doctor1
	doctor2.name = "Steve"
	fmt.Printf("%v", doctor2)
	fmt.Println()
	fmt.Printf("%v", doctor1)
	fmt.Println()

	//createBird()

}

type Animal struct {
	name   string `required`
	origin string
}

type Bird struct {
	Animal
	speedKPH float32
	canFly   bool
}

//func createBird() {
//	b := Bird{}
//	b.name = "Emu"
//	b.origin = "Australia"
//	b.speedKPH = 48
//	b.canFly = false
//
//	fmt.Println(b)
//	fmt.Println(b.name)
//
//}


//Summary
//Struct
//1 Collections of disparate data types that
//2 Normally created as types, but anonymous struct are allowed. Eg: generating a json response to a web service call
//3 Struct are valued type i.e. copied.
//4 No Inheritance, but can use composition via embedding.


