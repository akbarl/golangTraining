package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	//Demo1()
	//Demo2()
	//Demo3()
	//Demo4()
	//Demo5()
	//Demo6()
	//Demo7()
	//Demo8()
	//Demo9()
	//Demo10()
	//Demo11()
	//Demo12()
	//Demo13()
	//Demo14()
	tongChan, tongLe := demSoChanVaSoLe(1, 5)
	fmt.Println(tongChan)
	fmt.Println(tongLe)
}

func Demo1() {
	var age int
	var price float32
	var fullName string
	var status bool
	var key rune

	age = 20
	price = 4.5
	fullName = "ABC"
	status = true
	key = 'D'

	fmt.Println("age: ", age)
	fmt.Println("full name: ", fullName, ", status: ", status)
	fmt.Println("price: ", price)
	fmt.Println("key: ", key)

	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("full name: %s\n", fullName)
	fmt.Printf("%c position %d\n", key, key)
}

func Demo2() {
	var age = 20
	var fullName = "ABC"
	var price = 4.5
	var key = 'C'
	var status = true

	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("full name: %s\n", fullName)
	fmt.Printf("%c position %d\n", key, key)
}

func Demo3() {
	var (
		age      = 20
		fullName = "ABC"
		price    = 4.5
		key      = 'C'
		status   = true
	)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("full name: %s\n", fullName)
	fmt.Printf("%c position %d\n", key, key)
}

func Demo4() {
	var a1, a2, a3 = 4, 10, -6
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
}

func Demo5() {
	a := 3
	fmt.Println(a)
	price, fullName, status, age := 4.5, "DEF", true, 20
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("full name: %s\n", fullName)
}

func Demo6() {
	var a int = 5
	var b float32 = 5.6
	var result = float32(a) + b
	fmt.Println(result)

	firstName, lastName := "FN", "LN"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	var s string = "age: "
	var result2 = s + fmt.Sprintf("%d", a)
	fmt.Println(result2)
}

func Demo7() {
	a := 5
	if a > 3 {
		fmt.Println("a > 3")
	} else {
		fmt.Println("a <= 3")
	}

	if b := 2; b%2 == 0 {
		fmt.Println("b chan")
	} else {
		fmt.Println("b le")
	}
}

func Demo8() {
	menu := 2
	if menu == 1 {
		fmt.Println("Menu 1 is selected")
	} else if menu == 2 {
		fmt.Println("Menu 2 is selected")
	} else {
		fmt.Println("Invalid")
	}
}

func Demo9() {
	menu := 2
	switch menu {
	case 1:
		fmt.Println("Menu 1 is selected")
	case 2:
		fmt.Println("Menu 2 is selected")
	default:
		fmt.Println("Menu 3 is selected")
	}
}

func Demo10() {
	a := 2
	switch a {
	case 1, 2, 3:
		fmt.Println("a = 1, 2, 3")
	case 4, 5, 6, 7:
		fmt.Println("a = 4, 5, 6, 7")
	default:
		fmt.Println("Invalid")
	}
}

func Demo11() {
	a := 5
	switch {
	case a >= 1 && a <= 10:
		fmt.Println("a >=1 && and a <= 10")
	case a >= 11 && a <= 20:
		fmt.Println("a >= 11 && a <= 20")
	default:
		fmt.Println("Invalid")
	}
}

func Demo12() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%5d", i)
	}
}

func Demo13() {
	i := 1
	for i <= 10 {
		fmt.Printf("%5d", i)
		i++
	}
}

func Demo14() {
	i := 1
	for {
		fmt.Printf("%5d", i)
		i++
		if i > 10 {
			break
		}
	}
}

func Add1(a int, b int) int {
	return a + b
}

func Add2(a, b int) int {
	return a + b
}

func Add3(a, b int) (s int) {
	s = a + b
	return s
}

func calculate1(a, b int) (int, int, int, float32) {
	result1 := a + b
	result2 := a - b
	result3 := a * b
	result4 := float32(a) / float32(b)
	return result1, result2, result3, result4
}

func calculate2(a, b int) (result1 int, result2 int, result3 int, result4 float32) {
	result1 = a + b
	result2 = a - b
	result3 = a * b
	result4 = float32(a) / float32(b)
	return result1, result2, result3, result4
}

func Demo15() {
	value, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

/*
1. Xay dung ham tinh dien tich va chu vi HCN
	dt = a * b
	cv = (a + b ) * 2
2. xay dung ham tinh tong cac so chan, tong so le trong khoan tu m den n
3. Xay dung ham dem co bao nhieu so chan, bao nhieu so le trong khoan tu m den n
*/

func tinhDienTich(a, b int) (dt, cv int) {
	dt = a * b
	cv = (a + b) * 2
	return dt, cv
}

func tinhTongChanVaTongLe(m, n int) (tongChan, tongLe int) {
	for i := m; i <= n; i++ {
		if i%2 == 0 {
			tongChan += i
		} else {
			tongLe += i
		}
	}
	return tongChan, tongLe
}

func demSoChanVaSoLe(m, n int) (soLe, soChan int) {
	for i := m; i <= n; i++ {
		if i%2 == 0 {
			soChan++
		} else {
			soLe++
		}
	}
	return soChan, soLe
}
