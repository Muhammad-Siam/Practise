package main

import "fmt"

func main() {
	fmt.Println("Hello, This Siam from Bangladesh...\n")

	fmt.Println("If you want to know about your age please fill up the input at below:\n")

	fmt.Println("Enter your age:")
	var age1 int
	fmt.Scanf("%d", &age1)
	fmt.Println("Your age is:", age1)

	switch age1 {
	case 1, 2:
		fmt.Println("Infant")
		fallthrough

	case 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
		fmt.Println("Children")

	case 13, 14, 15, 16, 17, 18:
		fmt.Println("teen age")

	default:
		fmt.Println(" You are an adult")

	}

	type EnglishBook struct {
		Title   string
		Classes string

		Writers string

		Writers1 string
		Writers2 string
		Writers3 string
		Writers4 string
		Writers5 string

		Editor  string
		Editor1 string

		Footer string
	}

	var b1 EnglishBook

	b1.Title = "English Grammer and Composition"
	b1.Classes = "Classes Nine-Ten"

	b1.Writers = "Writers"

	b1.Writers1 = "Shaheen M.Kabir"
	b1.Writers2 = "Jahurul Islam"
	b1.Writers3 = "M.Shahidullah"
	b1.Writers4 = "Goutam Roy"
	b1.Writers5 = "Mohammad Ali"

	b1.Editor = "Editor"
	b1.Editor1 = "Shaheen M.Kabir"
	b1.Footer = "National Curriculum and textbook board, Bangladesh"

	fmt.Println(b1.Title, "\n")
	//"\n" for giving new line.
	fmt.Println(b1.Classes, "\n")

	fmt.Println(b1.Writers)

	fmt.Println(b1.Writers2)

	fmt.Println(b1.Writers3)

	fmt.Println(b1.Writers4)

	fmt.Println(b1.Writers5)

	fmt.Println(b1.Editor)

	fmt.Println(b1.Editor1)

	fmt.Println(b1.Footer)
	// struct literals

	c1 := struct {
		Name       string
		Price      int
		Downloaded int
		Author     string
	}{
		Name:       "Career Education",
		Price:      132,
		Downloaded: 150,
		Author:     "Siam Ahmed",
	}
	fmt.Println(c1)

	//struct and literal struct done

	//Right now we will know about map

	//nill map
	var siam = make(map[string]int)
	fmt.Println(siam)
	//map.Example1
	var s = map[string]string{
		"Full_name":   "Auhidul Islam Siam",
		"School_name": "Kk govt institution",
	}
	fmt.Println(s)

	//Literal map Example
	x := make(map[string]string)
	x["Name"] = "SIAM"
	x["age"] = "5.7"
	delete(x, "age")
	fmt.Println(x)
}
