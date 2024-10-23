package main

import "fmt"

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i, x := range nums {
		for j := i + 1; j < l; j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("请输入数组大小")
	var n int
	fmt.Scan(&n)
	fmt.Println("请输入数组元素")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println("请输入目标数")
	var target int
	fmt.Scan(&target)
	result := twoSum(nums, target)
	fmt.Println(result)
}
