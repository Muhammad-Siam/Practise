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

	i:=1+1

	for i<10 {

	fmt.Println(i)

	i++


}

}