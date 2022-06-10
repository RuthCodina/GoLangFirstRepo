package main

import "fmt"

func main() {
	//print number from 1 to 100
	/*x := 1
	for {
		if x <= 100 {
			fmt.Println(x)
			x++
		}
	}*/

	/*Print every rune code point of the uppercase alphabet three times. Your output should look like
	this:
		65
			U+0041 'A'
			U+0041 'A'
			U+0041 'A'
		66
			U+0042 'B'
			U+0042 'B'
			U+0042 'B'
	… through the rest of the alphabet characters
	*/

	/*for i := 65; i <= 90; i++ {
		fmt.Print(i, "\n")
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}*/
	/*Create a for loop using this syntax
	● for condition { }
	Have it print out the years you have been alive.
	*/
	/*year := 2000

	for year <= 2022 {
		fmt.Println(year)
		year++
	}*/
	/*
		Create a for loop using this syntax
			● for { }
		Have it print out the years you have been alive.

	*/
	/*bd := 2000
	for {
		if bd <= 2022 {
			fmt.Println(bd)
			bd++
		}
	}
	*/
	/*
		Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4
	*/
	/*
		for i := 0; i <= 100; i++ {
			fmt.Println(i % 4)
		}
	*/
	/*
		Create a program that shows the “if statement” in action
	*/
	/*
		b := true
		if reflect.TypeOf(b) == reflect.TypeOf("") {
			fmt.Print("es un string")
		} else if reflect.TypeOf(b) == reflect.TypeOf(0) {
			fmt.Print("es un number")
		} else {
			fmt.Print("es un booleano")
		}
	*/
	/*
		Create a program that uses a switch statement with the switch expression specified as a
		variable of TYPE string with the IDENTIFIER “favSport”.
	*/
	favsport := ""
	switch favsport {
	case "tennis":
		fmt.Print("the tennis is fav")
	case "soccer":
		fmt.Print("the soccer is fav")
	default:
		fmt.Print("no hay fav")
	}
}
