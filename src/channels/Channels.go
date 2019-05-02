package main

import (
	"fmt"
	"sync"
)

//Agenda
//- creating channels
//- sending and receiving messages
//- buffered channels
//- for...range loops with channels
//- closing channels

var wg = sync.WaitGroup{}

func main() {

	example1()
	//**************************
	//example2()
	example3()

}

func example1() {
	ch := make(chan int)
	wg.Add(2)
	fmt.Println("********************************")

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()

}

func example2() {
	ch := make(chan int)
	//ch := make(chan int, 50)
	wg.Add(2)
	fmt.Println("********************************")

	go func() {
		i := <-ch
		fmt.Println(i)
		//fmt.Println(<-ch)
		wg.Done()
	}()

	go func() {
		ch <- 42
		ch <- 52
		wg.Done()
	}()
	wg.Wait()
}

func example3() {
	ch := make(chan int, 50)
	wg.Add(2)
	fmt.Println("********************************")

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42
		ch <- 50
		ch <- 51
		ch <- 52
		ch <- 53
		wg.Done()
	}()
	wg.Wait()
	close(ch)


	for i := range ch {
		fmt.Println("ch # ", i)

	}


}


//Summary
//1 channel basics
//	create a channel with make command
//		make(chan int)
//	send message into channel
//		ch <- val
//	receive message from channel
//		val <- ch
//	can have multiple sender and receivers
//2 Buffered channels
//	channels block sender side till receiver is available
//	block receiver side till message is available
//	can decouple sender and receiver with buffered channels
//		make(chan int,50)
//	Use buffered channels when sender and receiver have assymmetric loading
//3 For range loop with channels

