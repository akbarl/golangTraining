package entities

import "fmt"

type Student struct {
	Id    string
	Name  string
	Age   int
	Score float64
}

func NewStudent(id string, name string, age int, score float64) Student {
	student := Student{id, name, age, score}
	return student
}

func (student Student) ToString() string {
	return fmt.Sprintf("id : %s \n name: %s \n age: %d \n score: %0.2f", student.Id,
		student.Name, student.Age, student.Score)
}

func (student Student) SetAge1(age int) {
	student.Age = age
}

func (student *Student) SetAge2(age int) {
	student.Age = age
}
