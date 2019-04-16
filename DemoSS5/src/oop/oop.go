package oop

import (
	"entities"
	"fmt"
)

func Demo1() {
	student1 := entities.Student{
		Id:    "st01",
		Name:  "name 1",
		Age:   20,
		Score: 6,
	}

	fmt.Println(student1.ToString())
}

func Demo2() {
	student2 := entities.NewStudent("st02", "name 2", 22, 5.6)
	fmt.Println(student2.ToString())
	student2.SetAge1(25)
	fmt.Println(student2.ToString())

	student2.SetAge2(25)
	fmt.Println(student2.ToString())
}

/*
Khai bao 1 struct Point co 2 thuoc tinh nhu sau
x kieu so nguyen
y kieu so nguyen
1) Xay dung ham tinh do dai giua 2 diem theo cong thuc sau
AB = Math.sqrt((Xb - Xa) * (Yb - Ya) + (Yb + Ya) * (Yb + Ya))
2) thuc hien tinh chu vi tao thanh giua 3 diem A,B va C
*/
