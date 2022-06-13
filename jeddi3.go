package main

import "fmt"

func main() {
	/*
	   Using a COMPOSITE LITERAL:
	   	● create an ARRAY which holds 5 VALUES of TYPE int
	   	● assign VALUES to each index position.
	   	● Range over the array and print the values out.
	   	● Using format printing ○ print out the TYPE of the array
	*/
	array := [5]int{}

	array[0] = 10
	array[1] = 9
	array[2] = 8
	array[3] = 7
	array[4] = 6

	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("array --> %T\n", array)

	/*
		Using a COMPOSITE LITERAL:
		● create a SLICE of TYPE int
		● assign 10 VALUES
		● Range over the slice and print the values out.
		● Using format printing
		○ print out the TYPE of the slice

	*/
	slice := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("slice --> %T\n", slice)

	/*
		Using the code from the previous example, use SLICING to create the following new slices which are then printed:
		● [42 43 44 45 46]
		● [47 48 49 50 51]
		● [44 45 46 47 48]
		● [43 44 45 46 47]
	*/
	slice2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice2[:5])
	fmt.Println(slice2[5:])
	fmt.Println(slice2[2:7])
	fmt.Println(slice2[1:6])

	/*
		Follow these steps:
		● start with this slice
		○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		● append to that slice this value
		○ 52
		● print out the slice
		● in ONE STATEMENT append to that slice these values
		○ 53
		○ 54
		○ 55
		● print out the slice
		● append to the slice this slice
		○ y := []int{56, 57, 58, 59, 60}
		● print out the slice
	*/
	slice3 := append(slice2, 52, 54, 55)
	fmt.Println("slice3 -->", slice3)

	slice4 := []int{56, 57, 58, 59, 60}

	slice5 := append(slice3, slice4...)
	fmt.Println("slice5 -->", slice5)

	/*
		To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
		● start with this slice
		○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		● use APPEND & SLICING to get these values here which you should ASSIGN to a
		variable “y” and then print:
		○ [42, 43, 44, 48, 49, 50, 51]
	*/
	slice6 := append(slice2[:3], slice3[6:10]...)
	fmt.Println("slice6 -->", slice6)

	/*
		Create a slice to store the names of all of the states in the United States of America. Use make and append to do this. Goal: do not have the array that underlies the slice created more than once. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position, without using the range clause
	*/
	slice7 := make([]string, 50, 100)

	slice7 = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	for i, v := range slice7 {
		fmt.Println(i, v)
	}

	/*
		Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "Helloooooo, James."
		Range over the records, then range over the data in each record
	*/

	slice8 := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Print("slice8 -->\n")
	for i, v := range slice8 {
		fmt.Println(i, v)
	}

	/*
		Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
			`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
			`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
			`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
	*/
	map1 := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("map1 -->")
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	/*
		Using the code from the previous example, add a record to your map. Now print the map out
		using the “range” loop
	*/
	map1["agent086"] = []string{`being funny`, `women's impression`} //áca sí obligatoriamente se debe poner el composite literal
	for key, value := range map1 {
		fmt.Println(key, value)
	}

	/*
		Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
	*/
	delete(map1, "agent086")
	fmt.Println(map1)
}
