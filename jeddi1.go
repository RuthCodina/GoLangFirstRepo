package main

import "fmt"

var a int
var b string
var c bool

type jeddaType int

var d jeddaType

const (
	f int = 31
	g     = true
	h     = ""
)

//iota es una condición para constatnes que permite manejar incrementales.
const (
	aa = 2022 + iota
	ab = 2022 + iota
	ac = 2022 + iota
	ad = 2022 + iota
)

func main() {
	x := 42
	y := "james bond"
	z := true
	s := fmt.Sprintf("%T\t%T\t%T\t", a, b, c)
	d = 42

	e := (12 < 12)
	// shift operator: (>>) pone a la izquierda  del numero entero tantos bits (en  0) como se especifíque, y (<<) pone a la derecha tantos bits como se le específique.
	d = 42 << 1

	fmt.Println(x, y, z)
	fmt.Println(s)
	fmt.Println("d ->", d)
	fmt.Printf("%b\t%d\t%#X\n", d, d, d)
	fmt.Print(e, "\n")
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
}
