package main
import "fmt"
func rectangle(l *int, w *int){
	*l = *l +39
	*w = *w + 32

	return 
}


func main(){

	l:= 10
	w:= 14

	rectangle(&l, &w)

	

	fmt.Println(l+w, &l)


}