package prog100

import "fmt"

type User interface {
	PrintName()
	PrintDetails()
}
type Admin struct {
	Name string
	Branch string
}

func (admin Admin) PrintName()  {
	fmt.Println("PrintName")
}
func (admin Admin) PrintDetails()  {
	fmt.Println("PrintDetails")
}
func speak(user User)  {
	user.PrintName()
	user.PrintDetails()
	fmt.Println("helloooooo")

}
