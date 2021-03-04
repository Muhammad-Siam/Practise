package main

import "fmt"

func main() {
	//var declaration
	var siam string
	siam = "Siam"
	var rahat int = 45
	var rupa float32 = 40.4

	var n, m int = 11, 12

	fmt.Println(siam, rahat, rupa, n, m)

	var (
		name  string
		email string
		age   int32
	)
	name = "Siam"

	email = "rootersiam2255@gmail.com"
	age = 24

	fmt.Println(name, email, age)

	//short hand variable declaration

	country := "Bangladesh"
	indp := 1971

	fmt.Println(country, indp)

	//Arithmetic operators

	var (
		value  int
		value1 int
		value2 int
	)

	value = 12
	value1 = 13
	value2 = 14

	//addition
	c := value + value1 + value2
	fmt.Printf("%d + %d + %d = %d \n", value, value1, value2, c)

	//subtraction

	d := value - value1 - value2

	fmt.Printf("%d - %d - %d = %d", value, value1, value2, d)

	//divission = /
	// multiplication = *

}
