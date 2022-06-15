package main

import "fmt"

/*
Create a value and assign it to a variable.
● Print the address of that value
*/
var value = 45

func main() {
	fmt.Println(&value)

	fmt.Println("p1 -->", p1)

	a := changeme(&p1)
	fmt.Println(a)
}

/*
● create a person struct
● create a func called “changeMe” with a *person as a parameter
	○ change the value stored at the *person address
■ important
● to dereference a struct, use (*value).field
	○ p1.first
	○ (*p1).first
● “As an exception, if the type of x is a named pointer type and (*x).f
is a valid selector expression denoting a field (but not a method),
x.f is shorthand for (*x).f.”
	○ https://golang.org/ref/spec#Selectors
● create a value of type person
	○ print out the value
● in func main
	○ call “changeMe”
● in func main
	○ print out the value
*/

type person struct {
	first string
	last  string
	edad  int
}

func changeme(pd *person) person {
	//*&pd.first = "changed"
	(*pd).first = "firsChanged"
	return *pd
}

var p1 = person{
	first: "Antonella",
	last:  "Fuji",
	edad:  37,
}
