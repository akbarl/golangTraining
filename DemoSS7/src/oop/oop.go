package oop

import (
	"entities"
	"fmt"
)

func Demo1() {
	student1 := entities.Student{
		Id: "st01",
		Human: entities.Human{
			Name:   "Name 1",
			Gender: "Male",
		},
		Score: 8.1,
	}

	fmt.Println("Student 1 Info")
	fmt.Println(student1.ToString())

	employee1 := entities.Employee{
		Id: "e01",
		Human: entities.Human{
			Name:   "Name 1",
			Gender: "Male",
		},
		Salary: 200000.0,
	}
	fmt.Println("Employee 1 Info")
	fmt.Println(employee1.ToString())

}

func Demo2() {
	var animal entities.Animal
	cat := entities.Cat{}
	duck := entities.Duck{}

	animal = cat
	fmt.Println(animal.Sound())

	animal = duck
	fmt.Println(animal.Sound())
}

func Demo3() {
	cat1 := entities.Cat{}
	cat2 := entities.Cat{}

	duck1 := entities.Duck{}
	duck2 := entities.Duck{}
	duck3 := entities.Duck{}

	animals := []entities.Animal{cat1, duck1, duck2, cat2, duck3}

	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}

func Demo4() {
	hinhTron := entities.HinhTron{
		BanKinh: 4,
		Ten:     "Hinh Tron",
	}
	hinhVuong := entities.HinhVuong{
		Canh: 4,
		Ten:  "Hinh Vuong",
	}

	danhsachHinh := []entities.HinhHoc{hinhTron, hinhVuong}

	for _, hinh := range danhsachHinh {
		fmt.Println(hinh.TenHinh())
		fmt.Printf("ChuVi: %0.2f\nDienTich: %0.2f", hinh.ChuVi(), hinh.DienTich())
		fmt.Println()
	}
}
