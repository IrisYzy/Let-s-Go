package main

import "fmt"

func main() {
	result1 := twoSumFun1([]int{1, 5, 3, 7}, 10)
	fmt.Println("方法一：", result1)
	result2 := twoSumFun2([]int{4, 5, 2, 6, 3, 7}, 12)
	fmt.Println("方法二：", result2)

}

func twoSumFun1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSumFun2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	res := []int{}
	for k, v := range nums {
		if _, val2 := m[target-v]; val2 {
			res = []int{m[target-v], k}
			break
		}
		m[v] = k
	}
	return res
}
