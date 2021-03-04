package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20.2
	b := 25

	c := math.Pow(a, 2)

	fmt.Println(c, b)

	c = math.Sin(a)

	fmt.Println(c)

	r := 1.1
	p := 1.2

	t := math.Sqrt(r * p)

	fmt.Println(t)

	q := 8

	fmt.Println(q)

	q = q + 1
	fmt.Println(q)

	q++
	fmt.Println(q)

	//User input

	fmt.Print("Enter the value of a radius:")

	var radius float32

	fmt.Scanf("%d", &radius)

}
