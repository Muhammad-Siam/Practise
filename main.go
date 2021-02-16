package main

import "fmt"

func main() {
	//How to print?
	fmt.Println("Hello, Bangladesh! This is Siam.")
	//variable declaration with array
	var x [10]string

	x[0] = "Siam,"
	x[1] = "Rahat,"
	x[2] = "Sajid,"
	x[3] = "Safa,"
	x[4] = "Rupa,"
	x[5] = "Mahad,"
	x[6] = "Jayed,"
	x[7] = "Anaf,"
	x[8] = "Zara,"
	x[9] = "Farzana"

	fmt.Println(x)
	//short variable declaration with slice
	names := []string{"Siam", "Rahat", "Rupa", "Mahad", "Sajid", "Jayed"}

	var s []string = names[1:6]
	fmt.Println(s)

	//another way of variable declaration
	var (
		age    int    = 17
		name   string = "Siam"
		height int    = 15
	)

	fmt.Println(age, name, height)

	const (
		age1    int    = 16
		name1   string = "Rahat"
		height1        = 15
	)

	fmt.Println(age1, name1, height1)
	//short array
	numbers := [...]int{100, 97, 76, 78, 74}
	fmt.Println(numbers[0:5])
	fmt.Println(len(numbers))

	a := make([]int, 5)
	fmt.Println(a)

	fmt.Print("My name is Siam.")

	fmt.Print("I am a boy.")

	fmt.Println("I am learning Golang.")

	fmt.Println("I am a programmer.")

	fmt.Printf("%T\n ", numbers)
	fmt.Printf("%T\n", names)

	population := map[string]int{
		"dhaka":      3000323,
		"Chattagram": 3394343,
		"Sylhet":     1345532,
		"Khulna":     43536,
	}
	fmt.Println(population)
	fmt.Println(len(population))

	//struct

	type Doctor struct {
		number    int
		actorName string
	}
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
	}
	fmt.Println(aDoctor)

}
