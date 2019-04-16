package entities

import "fmt"

type Employee struct {
	Id     string
	Human  Human
	Salary float64
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("id: %s\nhuman: %s\nSalary: %0.2f", employee.Id, employee.Human.ToString(), employee.Salary)
}
