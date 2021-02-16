package main

import "fmt"

func main() {


	fmt.Println("Hello, Bangladesh. This is Siam. Today I am gonna use some condition")

	//How numbers Do you need to get chance into BUET

	fmt.Println("Enter your SSC GPA and Physics Number:")

	var GPA int
           var numberP int
		


	fmt.Scanf("%d %d", &GPA, &numberP)
	fmt.Println("Your GPA is:", GPA)
	fmt.Println("Your SSC Physics Number is:", numberP)


		

	

	if GPA ==5 {
	fmt.Println("You will able to take form for BUET")
	}else if GPA >=4 {
	fmt.Println("You will able to take form for BUET, but it will more harder")
	}else{
	fmt.Println("You won't able to take form for BUET")

}





	if numberP >=90{

	fmt.Println("You are welcome to BUET")
	}else{

	fmt.Println("Go for another UNIVERSITY")

	}



}

