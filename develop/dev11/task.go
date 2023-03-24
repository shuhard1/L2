package main

import (
	"sort"
)

func main() {

	nums := []int{3, 1, 4, 6, 2, 6, 5}

	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			println(nums[i-1])
		}
	}
}
