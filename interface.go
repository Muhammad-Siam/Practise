package main


import "fmt"

type Email interface{
	Write(string, string, string)
	Read()
	Send()string 
}

type Test struct{

	To string
	From string
	subject string
	Messege string
	Send string
}


func (t Test) Write(to string, from string, subject string){
 fmt.Println(to, from, subject)
}

func (t Test) Read(){
 fmt.Println(t.From, "received from")
}


func (t Test) Send()string {

	return t.To

}
type Manager struct{
	Name string

	Age int

	Education string

	Grades float32
}

func (m Manager) speak(){
	fmt.Println(m.Name, "can speak")

}


func (m Manager) getName() string{


	return m.Name

}

func (m Manager) setAge()int {

	return m.Age

}

func (m Manager) getEducation()string {

	return m.Education

}

func (m Manager) getGrades() float32{

	return m.Grades

}


func main() {

	var m Manager
	m.Name = "Siam Ahmed"
	m.Age = 17
	m.Education = "BUET"
	m.Grades = 5.00

	fmt.Println(m.getName())
	fmt.Println(m.getGrades())
	fmt.Println(m.getEducation())
	fmt.Println(m.setAge())
	m.speak()

	var t Test
	t.To = "rootersiam2255@gmail.com"
	t.From = "siam3434@gmail.com"
	t.Subject = "Air Pollution"
	t.Message = "Air pollution is very harmful for the society"
	t.Send = "Email sent"

	t.Write(t.To, t.From, t.Subject)
	fmt.Println(t.Message, t.Send)

}