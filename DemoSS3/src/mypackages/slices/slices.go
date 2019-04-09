package slices

import "fmt"

func Demo1() {
	var a = [5]int{5, 10, 20, 30, 40}
	slice1 := a[0:3]
	fmt.Println(slice1)
	slice1[1] = 999
	fmt.Println(a)

	slice2 := a[:3]
	fmt.Println(slice2)

	slice3 := a[3:]
	fmt.Println(slice3)

	slice4 := a[:]
	fmt.Println(slice4)
}

func Demo2() {
	a := []int{5, 10, 30, 50}
	fmt.Println(a)
	a = append(a, 30)
	fmt.Println(a)
	a = append(a, 40, 60, 70, 80)
	fmt.Println(a)
	b := []int{50, 10, 50}
	a = append(a, b...)
}

func Demo3() {
	var a = [5]int{5, 10, 7, 6, 7}
	slice1 := a[0:2]
	fmt.Println("slice1: ", slice1)
	fmt.Println("\tlen: ", len(slice1))
	fmt.Println("\tCap: ", cap(slice1))

	slice2 := a[2:5]
	fmt.Println("slice2: ", slice2)
	fmt.Println("\tlen: ", len(slice2))
	fmt.Println("\tCap: ", cap(slice2))
}

func Demo4() {
	slice1 := make([]int, 3, 3)
	//slice1 = append(slice1, 5)
	//slice1 = append(slice1, 8)
	slice1[0] = 4
	slice1[1] = 20
	slice1[2] = 7
	fmt.Println("=============")
	fmt.Println("\tlen: ", len(slice1))
	fmt.Println("\tCap: ", cap(slice1))
	fmt.Println(slice1)
	slice1 = append(slice1, 7)
	fmt.Println("\tlen: ", len(slice1))
	fmt.Println("\tCap: ", cap(slice1))
	fmt.Println(slice1)
	fmt.Println("=============")
	slice1 = append(slice1, 8, 9, 7)
	fmt.Println("\tlen: ", len(slice1))
	fmt.Println("\tCap: ", cap(slice1))
	fmt.Println(slice1)
}
