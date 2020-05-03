package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

//It will throw error if declare and not use
func main() {
	//const s = 1
	var s = 2
	s = s + 1 //cannot assign to constant

	num1, num2 := 5.4, 9.8 // declared but not used if not calling
	w1, w2 := "hai", "hello"

	var q int = 9
	var p float64 = float64(q)
	m := q // f will int type

	fmt.Println(multiple(w1, w2))
	fmt.Println(add(num1, num2))
}
