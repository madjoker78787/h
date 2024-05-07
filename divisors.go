package main

import (
	"fmt"
)

func commonDivisors(nums []int) []int {
	resultBool := make(map[int]bool)
	for i := 1; i <= nums[0]; i++ {
		isBool := true
		for _, num := range nums {
			if num % i != 0 {
				isBool = false
				break
			}
		}
		if isBool {
			resultBool[i] = true
		}
	}
	var resultList []int
	for val := range resultBool {
		if val == 1 {
			continue
		}
		resultList = append(resultList, val)
	}
	return resultList
}

func main() {
	nums := []int{42, 12, 18}
	fmt.Println(commonDivisors(nums))
}

