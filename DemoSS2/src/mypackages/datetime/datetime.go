package datetime

import (
	"fmt"
	"time"
)

func Demo1() {
	today := time.Now()
	fmt.Println("today: ", today)
	fmt.Println("year: ", today.Year())
	month := today.Month()
	fmt.Println("month: ", month)
	fmt.Println("month: ", int(month))
	fmt.Println("day: ", today.Day())
}

func Demo2() {
	today := time.Now()
	fmt.Println("today: ", today.Format("02/01/2006 15:04:05"))

	s := "20/10/1980"
	d, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d.Format("2006-01-02"))
	}
}

func Demo3() {
	today := time.Now()
	date1 := today.AddDate(0, 0, 3)
	fmt.Println("date 1: ", date1.Format("02/01/2006"))
	date2 := today.Add(4 * 24 * time.Hour)
	fmt.Println("date 2: ", date2.Format("02/01/2006"))
	result := date2.Sub(date1).Hours()
	fmt.Println("result: ", result)
}

/*
Khai bao 1 chuoi chua ngay sinh, Xay dung ham thuc hien cac yeu cau sau
1. Kiem tra hom nay co phai la sinh nhat hay khong ?
2. Kiem tra tuoi co du 18 hay khong
*/

func Demo4() {
	s := "22/11/1995"
	myBirth, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println(err)
	} else {
		today, err := time.Parse("02/01/2006", time.Now().Format("02/01/2006"))
		if err != nil {
			fmt.Println(err)
		} else {
			if myBirth.Day() == today.Day() && myBirth.Month() == today.Month() {
				fmt.Println("today is your birthday. Cheers!!!")
			} else {
				fmt.Println("Have a good day")
			}
		}
	}

	fmt.Println("is birthday: ", isBirthday(s))
	fmt.Println("is Age 18:  ", isAge18(s))
}

func isBirthday(s string) bool {
	today := time.Now()
	date, _ := time.Parse("02/01/2006", s)
	return int(today.Month()) == int(date.Month()) && today.Day() == date.Day()
}

func isAge18(s string) bool {
	today := time.Now()
	date, _ := time.Parse("02/01/2006", s)
	return today.Year()-date.Year() >= 18
}
