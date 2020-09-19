package main

import (
	"fmt"
)

func initail(nums []int, target int) []int {
	for i, num := range nums {
		for j, num2 := range nums {
			if num+num2 == target {
				if i != j {
					fmt.Println(i, j)
					return ([]int{i, j})
				}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 10)
	for i, num := range nums {
		val, ok := hash[num]
		if ok {
			fmt.Println(val, i)
			return []int{val, i}
		}
		hash[target-num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
	nums = []int{3, 2, 4}
	target = 6
	twoSum(nums, target)
	nums = []int{3, 3}
	target = 6
	twoSum(nums, target)
}
