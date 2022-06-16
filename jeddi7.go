package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(cars)
	//fmt.Println(JSNconvert(cars))
	byt, err := json.Marshal(cars)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byt))

	err2 := json.Unmarshal([]byte(str), &people)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(people)

	// Starting with this code, sort the []int and []string for each person

	numbers := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	sort.Ints(numbers)
	fmt.Println(numbers)

	people2 := []person{p1, p2, p3}
	fmt.Println(people2)

	sort.Sort(sortByAge(people2))
	for _, p := range people2 {
		fmt.Println(p.First)
		fmt.Println(p.Last)
		fmt.Println(p.Age)
		for _, say := range p.Sayings {
			fmt.Println("\t", say)
		}
	}

	fmt.Print("------------------")

	sort.Sort(sortByLast(people2))
	for _, p := range people2 {
		fmt.Println(p.First)
		fmt.Println(p.Last)
		fmt.Println(p.Age)
		for _, say := range p.Sayings {
			fmt.Println("\t", say)
		}
	}

}

/*
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this
hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a
package
*/
type car struct {
	Cilindraje float32
	Color      string
}

var c1 = car{
	Cilindraje: 1.6,
	Color:      "rojo",
}

var c2 = car{
	Cilindraje: 2.0,
	Color:      "plateado",
}

var cars []car = []car{c1, c2}

/*
func JSNconvert(v []car) ([]byte || error){
	var byt, err = json.Marshal(v)
	if err != nil {
		return err
	}
	return byt
}
*/
/*
Starting with this code,
[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]
unmarshal the JSON into a Go data structure. Hint
*/

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

var people []person

var str string = `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

// Starting with this code, sort the []int and []string for each person

/*
Starting with this code, sort the []user by
● age
● last
Also sort each []string “Sayings” for each user
● print everything in a way that is pleasant
*/

var p1 = person{
	First: "jhon",
	Last:  "second",
	Age:   25,
	Sayings: []string{
		"noooppo, what to do",
		"georgeos",
	},
}

var p2 = person{
	First: "Sandra",
	Last:  "cardona",
	Age:   43,
	Sayings: []string{
		"noooppo, what to do",
		"georgeos",
	},
}

var p3 = person{
	First: "Manolo",
	Last:  "diaz",
	Age:   33,
	Sayings: []string{
		"noooppo, what to do",
		"georgeos",
	},
}

type sortByAge []person

func (p sortByAge) Len() int           { return len(p) }
func (p sortByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p sortByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }

type sortByLast []person

func (p sortByLast) Len() int           { return len(p) }
func (p sortByLast) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p sortByLast) Less(i, j int) bool { return p[i].Last < p[j].Last }
