package main

import "fmt"

//the ZERO VALUE is False for boolean, 0 for integers, 0.0 for floats, "" for strings, nil for (pointers, functions, interfaces, slices, channels, maps)

var y = 30
var a = `camino va "camilo"`

// En Go puede crear su propio tipo de dato
type ruthType string

var c ruthType

func main() {
	// Para saber el tipo de dato de una variable %T
	fmt.Printf("%T\t%x\t%b\n", y, y, y) //usa Printf, para imprimir con formato
	fmt.Println(y)                      //usa PrintLN, para imprimir con un salto de linea al final.
	fmt.Print(a, "\n")

	c = "Fabuloso"
	fmt.Printf("%T\n", c)

	//An explicit conversion is an expression of the form T(x) where T is a type and x is an expression that can be converted to type T.
	a = string(c)
	fmt.Println(a)

}
