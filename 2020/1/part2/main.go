package main

import "fmt"

func main() {
	nums := make([]int, 200)
	for i := 0; i < 200; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			for k := j + 1; k < 200; k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					fmt.Println(nums[i] * nums[j] * nums[k])
				}
			}
		}
	}
}
