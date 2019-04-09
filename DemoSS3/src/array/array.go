package array

import "fmt"

func Demo1() {
	var a [5]int
	a[0] = 10
	a[1] = -5
	a[2] = 2
	a[3] = 4
	a[4] = 5
	fmt.Println(a)
	fmt.Println("len: ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%5d", a[i])
	}

	fmt.Println()

	for index, value := range a {
		fmt.Println(index, ": ", value)
	}

	for _, value := range a {
		fmt.Println("value: ", value)
	}
}

func Demo2() {
	var a = [5]int{5, 10, 15, 20, 25}
	for _, value := range a {
		fmt.Println("value: ", value)
	}

	names := [3]string{"name1", "name2", "name3"}
	for _, name := range names {
		fmt.Println("name: ", name)
	}
}

func Demo3() {
	a := [...]int{5, 10, 20, 30}
	fmt.Println("len: ", len(a))

	for i := 0; i < len(a); i++ {
		fmt.Printf("%5d", a[i])
	}

	fmt.Println()

	for index, value := range a {
		fmt.Println(index, ": ", value)
	}

	for _, value := range a {
		fmt.Println("value: ", value)
	}
}

func display(a [5]int) {
	for index, value := range a {
		fmt.Println(index, ": ", value)
	}
}

func Demo4() {
	var a = [5]int{5, 10, 20, 30, 40}
	display(a)
}

/*
1. Xay dung ham tinh tong cac so chan, so le, so am, so duong trong mang.
Tra ve danh sach gia tri
2. Xay dung ham dem cac so chan, so le, so am, so duong trong mang.
Tra ve danh sach gia tri
3. Xay dung ham tim gia tri min, max trong mang
Tra ve danh sach gia tri
4. Xay dung ham sap xep 1 mang theo thu tu tang dan
Tra ve mang sau khi sap xep
*/
func TinhTong(a [10]int) (tongChan int, tongLe int, tongAm int, tongDuong int) {
	for _, value := range a {
		if value%2 == 0 {
			tongChan += value
		} else {
			tongLe += value
		}

		if value > 0 {
			tongDuong += value
		} else if value < 0 {
			tongAm += value
		}

	}
	return tongChan, tongLe, tongAm, tongDuong
}

func Dem(a [10]int) (tongChan int, tongLe int, tongAm int, tongDuong int) {
	for _, value := range a {
		if value%2 == 0 {
			tongChan++
		} else {
			tongLe++
		}

		if value > 0 {
			tongDuong++
		} else if value < 0 {
			tongAm++
		}

	}
	return tongChan, tongLe, tongAm, tongDuong
}

func TimGiaTriMinMax(a [10]int) (min int, max int) {
	for _, value := range a {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return min, max
}

func SapXepMangTangDan(a [10]int) [10]int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}
