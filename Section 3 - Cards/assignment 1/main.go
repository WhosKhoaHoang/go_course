package main

import "fmt"

func main() {
	// Determining parity on the fly
	// nums := []int{}
	// for i := 0; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i, " is even")
	// 	} else {
	// 		fmt.Println(i, " is odd")
	// 	}
	// 	nums = append(nums, i)
	// }

	// Alternative way using make()
	nums := make([]int, 11)
	for i, _ := range nums {
		nums[i] = i
	}
	for _, val := range nums {
		if val % 2 == 0 {
			fmt.Println(val, " is even")
		} else {
			fmt.Println(val, " is odd")
		}
	}
}