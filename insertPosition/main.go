package main

import (
	"fmt"
	"sort"
)

func main() {
	nums :=[]int{1,3,5,6}
	target:=4
fmt.Println(searchInsert(nums,target))
}

func searchInsert(nums []int, target int) int {
	for i:= range nums{

		if target==nums[i]{

			return i
		}else {
			if target != nums[i]{

				a:=sort.SearchInts(nums,target)
				return a
			}
		}
	}

	return 0
}