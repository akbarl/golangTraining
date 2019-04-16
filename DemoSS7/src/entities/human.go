package entities

import "fmt"

type Human struct {
	Name   string
	Gender string
}

func (human Human) ToString() string {
	return fmt.Sprintf("name: %s\ngender: %s", human.Name, human.Gender)
}
