package pointer

import "fmt"

func Demo1() {
	var a int = 5
	fmt.Println("a: ", a)
	fmt.Println("address a: ", &a)

	var p *int = &a

	fmt.Println("address p: ", p)
	fmt.Println("value p: ", *p)

	*p = 10

	fmt.Println("a: ", a)
	q := &a

	fmt.Println("address q: ", q)
	fmt.Println("value q: ", *q)

}

func Demo2() {
	a, b := 5, 10
	p, q := &a, &b
	fmt.Println("address p: ", p)
	fmt.Println("address q: ", q)
}

func Demo3() {
	a := 5
	changeValue1(a)
	fmt.Println("a = ", a)
	changeValue2(&a)
	fmt.Println("a = ", a)
}

func changeValue1(a int) {
	a = 6
}

func changeValue2(p *int) {
	*p = 7
}

func Demo4() {
	a, b := 5, 10
	swap(&a, &b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}

func swap(p, q *int) {
	temp := *p
	*p = *q
	*q = temp
}
