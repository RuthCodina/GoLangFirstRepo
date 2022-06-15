package main

import (
	"fmt"
	"math"
)

var arg = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	defer foo3()

	a := foo()
	b, c := bar()
	fmt.Println(a, b, c)

	d := foo1(arg...)
	fmt.Println(d)

	e := bar1(arg)
	fmt.Println(e)

	p1 := person{
		first: "Louis",
		last:  "lane",
		age:   32,
	}

	p1.speak()

	info(square1)
	info(circle1)

	/*
		Build and use an anonymous func
	*/
	func() {
		arg2 := append(arg, 10, 11)
		fmt.Println("arg2 --> ", arg2)
	}()

	fmt.Println("arg3 -->", fun2())

	fmt.Println("fun3 -->", fun3(arg))

	//f := fun4(fun3(arg), arg)
	//fmt.Println("fun4 -->", f())

	fmt.Println("fun5 -->", fun5())
}

//----------------------------------------------------------------------------------------

/*
	Hands on exercise
		○ create a func with the identifier foo that returns an int
		○ create a func with the identifier bar that returns an int and a string
		○ call both funcs
		○ print out their results
*/
func foo() int {
	return 42
}

func bar() (int, string) {
	return 31, "aprendiendo..."
}

/*
	create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
	create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/
func foo1(x ...int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum
}

func bar1(x []int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum
}

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/
func foo3() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}

/*
 Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ last
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("hola, mi nombre es ", p.first, " ", p.last, " y tengo ", p.age, " años")
}

/*
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/

type shape interface {
	area() float64
}

type square struct {
	lado float32
}
type circle struct {
	radio float32
}

func (s square) area() float64 {
	area := math.Pow(float64(s.lado), 2.0)
	return area
}

func (c circle) area() float64 {
	area := float64(c.radio) * math.Pow(math.Pi, 2.0)
	return area
}

func info(s shape) {
	fmt.Println("area es -->", s.area())
}

var square1 = square{
	lado: 3.451,
}

var circle1 = circle{
	radio: 68.45,
}

/*
 Assign a func to a variable, then call that func
*/
var fun2 = func() []int {
	arg3 := append(arg, 21, 22, 23, 24, 25)
	return arg3
}

/*
 Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/

func fun3(a []int) func() []int {
	return func() []int {
		arr := []int{}
		for _, v := range a {
			arr = append(a, v)
		}
		fmt.Println("arr -->", arr)
		return arr
	}
}

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
*/

func fun4(c func(a []int) func() []int, arr []int) func() []int {
	f := c(arr)
	return f
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this
hands-on exercise, create a func which “encloses” the scope of a variable:
*/

func fun5() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
