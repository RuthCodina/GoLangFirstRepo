package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		● in addition to the main goroutine, launch two additional goroutines
			○ each additional goroutine should print something out
		● use waitgroups to make sure each goroutine finishes before your program exists
	*/
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())

	wg.Add(2)

	go routine()
	go routine1()

	saySomething(&p)
	// saySomething(p) //doesn't work
	fmt.Println("GoRoutines\t\t", runtime.NumGoroutine())
	wg.Wait()
}

func routine() {
	fmt.Println("segunda Rutina")
	wg.Done()
}

func routine1() {
	fmt.Println("Tercera rutina")
	wg.Done()
}

/*
This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
	○ *person
● create a type human interface
	○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
	○ have it take in a human as a parameter
	○ have it call the speak method
● show the following in your code
	○ you CAN pass a value of type *person into saySomething
	○ you CANNOT pass a value of type person into saySomething
*/

type persona struct {
	first string
	age   int
}

func (p *persona) speak1() {
	fmt.Println("Hallo word")
}

type human interface {
	speak1()
}

func saySomething(h human) {
	h.speak1()
}

var p = persona{
	first: "Martha",
	age:   33,
}

/*
 Using goroutines, create an incrementer program
	○ have a variable to hold the incrementer value
	○ launch a bunch of goroutines
		■ each goroutine should
● read the incrementer value
	○ store it in a new variable
● yield the processor with runtime.Gosched()
● increment the new variable
● write the value in the new variable back to the incrementer
variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag
*/
