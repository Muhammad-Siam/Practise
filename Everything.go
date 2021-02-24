package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//simple hello world!
	fmt.Println("HEllO, WORLD!")

	//variables

	var siam string
	siam = "Name: Auhidul Islam Siam"

	var rahat int

	rahat = 15

	var siam1 float32

	siam1 = 34.5

	var rahat1 rune

	rahat1 = 'A'
	//one line declaration

	var siam2 rune = 'A'

	var rahat2 bool

	rahat2 = true

	var siam3 bool

	siam3 = true

	fmt.Println(siam, siam1, string(siam2), siam3, rahat, rahat1, rahat2)

	//short hand variable declaration

	var x = 1
	fmt.Println(x)

	y := 1
	fmt.Println(y)

	g := "Name: Auhidul Islam Siam"

	r := 'a'

	fmt.Println(g, r)
	//ASCII number ke convert korar niyom another is string(r)
	fmt.Printf("%c", r)
	//another simple way of doing this

	var (
		auhid = "siam"
		bd    = 23
		math  = 34
	)

	fmt.Println(auhid, bd, math)

	//booleans

	Na := 11
	Mg := 12

	fmt.Println("11 == 12:", Na == Mg)
	fmt.Println("11 > 12", Na > Mg)
	fmt.Println("11 >= 12", Na >= Mg)

	Na1 := "Sodium"
	Na2 := "Natrium"

	fmt.Println("Sodium == Natrium", Na1 == Na2)
	fmt.Println("Sodium != Natrium", Na1 != Na2)
	fmt.Println(reflect.TypeOf(Na1))

	//Operators: == , &&, !, <> , ||

	//Primitive DataTypes : string, int32, int64, int, float32, float, float64, rune, booleans,uint, uint32, uint64, uintptr
	//composite dataTypes : array , struct, slice,  map, function, channels, pointers, method, interface
	//to know about the type we can use reflect.TypeOf
	//to know about the size we can use unsafe.Sizeof
	var tu int = 2
	fmt.Println(reflect.TypeOf(tu))
	fmt.Println(unsafe.Sizeof(tu))
	fmt.Println(tu)

	//Taking user input

	/*fmt.Println("Enter your age to know about your voting:")

	var si int64

	fmt.Scanln(&si)

	fmt.Println("Your voting age is:", si)

	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Your name is:", name)*/

	//ARRAY

	var array [3]int

	array[0] = 23
	array[1] = 25
	array[2] = 34

	c3 := array[0:3]
	fmt.Println(c3)
	fmt.Println("My scores are:", array[0:3])
	a1 := [...]int{1, 2, 3, 4}
	fmt.Println(a1)

	a2 := [3]int{1, 2, 4}

	fmt.Println(len(a2))

	//slices

	sw := []int{1, 23, 32, 3, 6}

	fmt.Println(sw)

	var fruites []int

	fruites = append(fruites, 23, 23, 434, 43, 4, 4, 34, 3, 4, 3, 4, 3, 43)

	fmt.Println(fruites)
	fmt.Println(reflect.TypeOf(fruites))
	fmt.Printf("%T")

	//maps

	z := make(map[string]int)

	z["Siam"] = 232
	z["Rahat"] = 3434

	fmt.Println(z)

	ca := map[string]int{

		"Name": 342,
		"saim": 344,
	}
	fmt.Println(ca)

}
