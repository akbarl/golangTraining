package entities

import "fmt"

type Student struct {
	Id    string
	Human Human
	Score float32
}

func (student Student) ToString() string {
	return fmt.Sprintf("id: %s\nhuman: %s\nScore: %0.2f", student.Id, student.Human.ToString(), student.Score)
}
