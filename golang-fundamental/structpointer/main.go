package main

import "fmt"

func main() {
	student := Student{
		ID:   1,
		Name: "Adi Nugraha Putra",
		GPA:  4.0,
	}

	fmt.Println(student.Name)
	student.Graduate()
	fmt.Println(student.Name)
}

type Student struct {
	ID int
	Name string
	GPA	 float32
}

func (student *Student) Graduate()  {
	student.Name = student.Name + " S.T"
}