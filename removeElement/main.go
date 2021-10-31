package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	new_nums := []int{}
	for i := range nums {
		if nums[i] != val {
			new_nums = append(new_nums, nums[i])

		}
	}
	fmt.Println(new_nums)
	return len(new_nums)
}
