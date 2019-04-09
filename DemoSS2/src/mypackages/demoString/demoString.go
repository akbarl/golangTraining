package demoString

import (
	"fmt"
	"sort"
	"strings"
)

func Demo1() {
	s := "abcdef"
	fmt.Println("len: ", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func Demo2() {
	s1 := "abc"
	s1 = strings.ToUpper(s1)
	fmt.Println("s1: ", s1)

	s2 := "DEF"
	s2 = strings.ToLower(s2)
	fmt.Println("s2: ", s2)
}

func Demo3() {
	s1 := "computer"
	s2 := "com"
	result1 := strings.HasPrefix(s1, s2)
	fmt.Println("result 1:", result1)
	s3 := "ter"
	result2 := strings.HasSuffix(s1, s3)
	fmt.Println("result 2: ", result2)
	s4 := "put"
	result3 := strings.Contains(s1, s4)
	fmt.Println("result 3: ", result3)
}

/*
	Tao slice chua danh sach cac ten san pham. xay dung cac ham thuc hien
	cac yeu cau sau:
	1. Xay dung ham tim cac ten san pham co chua 1 keyword. Tra
	ve slice chua cac ten do
	2. Xay dung ham sap xep cac ten san pham theo thu tu giam dan
	3. Xay dung ham tra ve slice chua danh sach so nguyen to trong khoang
	tu a den b
*/

func Demo4() {
	s := []string{"San pham 1", "San pham 2", "San pham 3", "San pham 4"}
	result := findWithKeyWord(s, "San pham")
	Display(result)
	fmt.Println()
	sortProduct(s)
}

func findWithKeyWord(a []string, keyword string) (result []string) {
	result = make([]string, 5, 5)
	for _, value := range a {
		if strings.Contains(value, keyword) {
			result = append(result, value)
		}

	}
	return result
}

func sortProduct(a []string) {
	sort.Slice(a, func(i, j int) bool {
		return strings.Compare(a[i], a[j]) > 0
	})
	fmt.Println(a)
}

func Display(a []string) {
	for _, value := range a {
		fmt.Printf("%s ", value)
	}
}
