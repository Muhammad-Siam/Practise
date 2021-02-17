package main

import "fmt"


func main(){
	
	fmt.Println("Enter your Sunject name to know the requirments:")


	var subject string

	fmt.Scanln(&subject)


	fmt.Println("Your subject is:", subject)

	switch subject{
	
	case "Physics", "Chemistry":

	fmt.Println("You need 90% of marks")

	case "Math", "Higher Math":

	fmt.Println("You need 95% of marks")

	case "Biology", "Bangladesh And Global Studies":

	fmt.Println("You need 80% of marks")

	default:

	fmt.Println("You need 80% of marks")


	}


}