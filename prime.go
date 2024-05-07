package main

import (
	"fmt"
)

func prime(a, b int) (result []int) {
	for i := a; i <= b; i++ {
		if i%2 != 0 && i%3 != 0 {
			result = append(result, i)
		}
	}
	return
}

func main() {
	fmt.Println(prime(11, 20))
}
