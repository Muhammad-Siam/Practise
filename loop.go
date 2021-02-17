package main

import "fmt"

func main(){
	
	
	for i:=5; i<=6; i++ {

		fmt.Println("Hello!",i)
}



	subjects := []string {"Physics", "Chemistry", "Higher Math", "Higher Math"}


	for i, sub := range subjects {

	fmt.Println(i, sub)


}

	i:=0

	for i<10 {

	fmt.Println("I am gonna use that")

	i++


}

}