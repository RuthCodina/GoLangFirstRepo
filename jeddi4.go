package main

import "fmt"

func main() {
	/*
		Create your own type “person” which will have an underlying type of “struct” so that it can store
		the following data:
		● first name
		● last name
		● favorite ice cream flavors
		Create two VALUES of TYPE pers
	*/
	type person struct {
		first    string
		last     string
		favorite []string
	}

	person1 := person{
		first:    "amanda",
		last:     "lopez",
		favorite: []string{"chocolate", "vanilla"},
	}
	fmt.Println("person1 -->", person1)

	person2 := person{
		first:    "jorge",
		last:     "lorenz",
		favorite: []string{"donuts", "almonds"},
	}
	fmt.Println("person2 -->", person2)

	/*
		Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
	*/
	map1 := map[string]person{
		"p1": person1,
		"p2": person2,
	}

	for key, value := range map1 {
		fmt.Println(key, value.first)
		fmt.Println(key, value.last)
		for i, v := range value.favorite {
			fmt.Println(i, v)
		}
		fmt.Print("--------------\n")
	}

	/*
		● Create a new type: vehicle.
			○ The underlying type is a struct.
			○ The fields:
				■ doors
				■ color
		● Create two new types: truck & sedan.
			○ The underlying type of each of these new types is a struct.
			○ Embed the “vehicle” type in both truck & sedan.
			○ Give truck the field “fourWheel” which will be set to bool.
			○ Give sedan the field “luxury” which will be set to bool. solution
		● Using the vehicle, truck, and sedan structs:
			○ using a composite literal, create a value of type truck and assign values to the fields;
			○ using a composite literal, create a value of type sedan and assign values to the fields.
		● Print out each of these values.
		● Print out a single field from each of these values.
	*/
	// la definición de los types, se deberían hacer por fuera de la func main.
	type vehicule struct {
		doors int
		color string
	}
	type sedan struct {
		vehicule
		luxury bool
	}
	type truck struct {
		vehicule
		fourWheel bool
	}

	truck1 := truck{
		vehicule: vehicule{
			doors: 7,
			color: "yellow",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicule: vehicule{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println("truck1, sedan1 -->", truck1, sedan1)

	/*
		Create and use an anonymous struct.
	*/

	women := struct {
		first          string
		profession     string
		yearsExpertise int
		friends        map[string]string
	}{
		first:          "susana",
		profession:     "lawyer",
		yearsExpertise: 12,
		friends: map[string]string{
			"Jhon":   "lawyer",
			"Martha": "chef",
			"Eddy":   "Engineer",
		},
	}

	for key, value := range women.friends {
		fmt.Println(key, value)
	}
}
