package main

import "fmt"

func main() {
	fmt.Println(CenturyFromYear(1705))
	fmt.Println(CenturyFromYear(1900))
	fmt.Println(CenturyFromYear(1601))
	fmt.Println(CenturyFromYear(2000))
}

func CenturyFromYear(year int) int {
	century := 0
	if year%100 == 0 {
		century = year / 100
	} else {
		century = year/100 + 1
	}

	return century
}
