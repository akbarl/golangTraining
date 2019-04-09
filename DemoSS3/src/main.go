package main

import (
	"array"
	"fmt"
	"mypackages/circle"
	"mypackages/package1"
	"mypackages/package2"
	"mypackages/slices"
	"mypackages/square"
)

func main() {
	package1.Hello()
	fmt.Println("a + b = ", package2.Add(1, 2))
	fmt.Println("Circle")
	fmt.Println("Area = ", circle.Area(3.2))
	fmt.Println("Square")
	fmt.Println("Area = ", square.Area(3))
	array.Demo1()
	array.Demo2()
	array.Demo3()
	array.Demo4()

	a := [10]int{1, 2, 3, 4, 5, -1, -2, -3, -4, -5}
	tongChan, tongLe, tongAm, tongDuong := array.TinhTong(a)
	fmt.Println("tongChan =", tongChan)
	fmt.Println("tongLe = ", tongLe)
	fmt.Println("tongAm = ", tongAm)
	fmt.Println("tongDuong =", tongDuong)

	tongChan, tongLe, tongAm, tongDuong = array.Dem(a)
	fmt.Println("tongChan =", tongChan)
	fmt.Println("tongLe = ", tongLe)
	fmt.Println("tongAm = ", tongAm)
	fmt.Println("tongDuong =", tongDuong)

	min, max := array.TimGiaTriMinMax(a)
	fmt.Println("min =", min)
	fmt.Println("max = ", max)

	fmt.Println("arraySorted = ", array.SapXepMangTangDan(a))

	slices.Demo1()

	slices.Demo2()

	slices.Demo3()

	slices.Demo4()

}
