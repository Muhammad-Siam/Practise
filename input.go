package main

import "fmt"


func main(){
	fmt.Println("Enter your name:")

	var first string
	
	fmt.Scanln(&first)

	fmt.Println("Enter your age:")

	var second float32
	fmt.Scanln(&second)

	fmt.Println("Your name is:", first)
	fmt.Println("Your age is:", second)



}
