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
	bd := 2000
	for {
		if bd <= 2022 {
			fmt.Println(bd)
			bd++
		}
	}
}
