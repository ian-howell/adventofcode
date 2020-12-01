package main

import "fmt"

func main() {
	nums := make([]int, 200)
	for i := 0; i < 200; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			if nums[i]+nums[j] == 2020 {
				fmt.Println(nums[i] * nums[j])
			}
		}
	}
}
