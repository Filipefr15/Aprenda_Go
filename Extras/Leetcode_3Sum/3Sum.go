package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	answer := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 && i != j && i != k && j != k {
					row := []int{nums[i], nums[j], nums[k]}
					sort.Ints(row)
					answer = append(answer, row)
				}
			}
		}
	}
	answerSorted := separador(answer)
	return answerSorted
}

func separador(nums [][]int) [][]int {
	//numsQueSeraDevolvido := nums
	numOriginais := nums[0]
	for i := 0; i < len(nums); i++ {
		numNums := nums[i]
		if numOriginais[0] == numNums[0] {
			nums = append(nums[:1], nums[2:]...)
		}
	}
	return nums
}

func main() {
	numeros := []int{-1, 0, 1, 2, -1, -4}
	printarei := threeSum(numeros)
	fmt.Println(printarei)
}
