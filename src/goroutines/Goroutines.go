package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)


//Topics covered include:
//- creating goroutines
//- synchronizing goroutines

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("start")

	example1()

	example2()


	//********************
	wg.Add(2)
	example3()
	example3()
	wg.Wait()


	//*******************
	example4()

	//**********************
	example5()

}

func example1() {

	go func() {
		fmt.Println("example1")
	}()
	time.Sleep(100 * time.Millisecond)
}

func example2() {
	wg.Add(1)
	go func() {
		fmt.Println("example2")
		wg.Done()
	}()

	wg.Wait()
}

func example3() {
	go func() {
		fmt.Println("example3")
		wg.Done()
	}()

}



//**********************************
var counter int = 0
var m = sync.RWMutex{}

func example4() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

}
func sayHello() {
	fmt.Println("Hello # ", counter)
	m.RUnlock()
	wg.Done()

}
func increment() {
	counter++
	m.Unlock()
	wg.Done()

}


func example5() {
	//runtime.GOMAXPROCS(100)
	fmt.Printf("Threads : %v",runtime.GOMAXPROCS(-1))
	fmt.Println()

}

//Summary
//1 Creating goroutine
//	Use go keyword in front of function call
//2 Synchrinization
//	Use sync.WaitGroup to wait for groups of goroutines to complete
//	Use sync.RWMutex to protect data access
//3. Parallelism
//	By default, Go will use CPU threads equal to available cores
//	Change with runtime.GOMAXPROCS
//	More threads can increase performance, but too many can slow it down