package main

import (
	"fmt"
	"math/rand"
)

func hei() {
	fmt.Print("calling method")
}

func main() {
	hei()
	fmt.Println("hello world is", rand.Intn(100))
}
