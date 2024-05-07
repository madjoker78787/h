package main

import (
	"fmt"
)

func multiplication(n int) {
	fmt.Printf("%4s", "")
	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%4d", i)

		for x := 1; x <= n; x++ {
			result := i * x
			fmt.Printf("%4d", result)
		}
		fmt.Println()
	}
}

func main() {
	multiplication(5)

}


// out
   //     1   2   3   4   5
   // 1   1   2   3   4   5
   // 2   2   4   6   8  10
   // 3   3   6   9  12  15
   // 4   4   8  12  16  20
   // 5   5  10  15  20  25
