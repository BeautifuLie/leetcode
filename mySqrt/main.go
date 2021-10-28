package main

import (
	"fmt"
	"math"
)

func main() {
	x:=8
fmt.Println(mySqrt(x))
}

func mySqrt(x int) int {
	res:= math.Sqrt(float64(x))

	return int(res)
}