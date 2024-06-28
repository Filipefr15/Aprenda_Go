package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				answer := []int{i, j}
				return answer
			}
		}
	}
	return nums
}

func main() {
	numeros := []int{3, 2, 4}
	printarei := twoSum(numeros, 6)
	fmt.Println(printarei)
}
