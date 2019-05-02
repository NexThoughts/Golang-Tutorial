//package main
//
//import "fmt"
//
//type User interface {
//	PrintName()
//	PrintDetails()
//}
//type Admin struct {
//	name string
//}
//
//func (admin Admin) PrintName()  {
//	fmt.Println("PrintName")
//}
//func (admin Admin) PrintDetails()  {
//	fmt.Println("PrintDetails")
//}
//func speak(user User)  {
//	user.PrintName()
//	user.PrintDetails()
//	fmt.Println("helloooooo")
//
//}
//func main() {
//	admin := Admin{
//		name:"Rohit",
//	}
//	speak(admin)
//	fmt.Println(admin)
//
//}