package main

import "fmt"

func main() {
	n := 1298734
	fmt.Println(NthEven(n))
}

func NthEven(n int) int {

	if n == 1 {
		return 0
	}

	res := n*2 - 2
	return res

}
