package main

import "fmt"

func main() {
 
	fmt.Println("Hello, Bangladesh! This is Siam.")

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

	primes := []string{"Siam", "Rahat", "Rupa", "Mahad", "Sajid", "Jayed"}

	var s []string = primes[1:6]
	fmt.Println(s)
}

	