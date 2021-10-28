package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i,n1:= range nums{
		for k,n2:= range nums{
			if n1+n2 ==target && i!=k{

				return []int{i,k}
			}
		}
	}
	return nil
}




func main() {

	var nums = []int{3,4,3}

fmt.Println(twoSum(nums,6))



}
