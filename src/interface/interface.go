package main

import "fmt"

func main() {
	example1()
	example2()

}

type Vehicle interface {
	totalSpeed(int) (int)
}
type Bicycle struct {
	speed int
}

func (bicycle Bicycle) totalSpeed(increment int) (int) {
	totalSpeed := bicycle.speed + increment
	fmt.Println(totalSpeed)
	return totalSpeed
}

func example1() {
	vehicle := Bicycle{
		speed:50,
	}
	vehicle.totalSpeed(50)
}

//******************************************************************************
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (intCounter IntCounter) Increment() int {
	intCounter++
	return int(intCounter)
}

func example2() {
	var incrementer Incrementer = IntCounter(1)
	fmt.Println(incrementer.Increment())

}
