package _136只出现一次的数字

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/
/*
示例 1:
输入: [2,2,1]
输出: 1

示例 2:
输入: [4,1,2,1,2]
输出: 4
*/

func singleNumber(nums []int) int {
	//numsMap := make(map[int]int)
	//var result int
	//for _,v := range nums {
	//	if numsMap[v] > 0  {
	//		fmt.Println("前:",numsMap)
	//		delete(numsMap,v)
	//		fmt.Println("后:",numsMap)
	//		continue
	//	}
	//	numsMap[v] = v
	//	result = numsMap[v]
	//	//fmt.Println(numsMap)
	//}

	//return result

	var out = 0

	for i := 0; i < len(nums); i++ {
		out = out ^ nums[i]
	}

	return out
}

func main() {
	inputData := []int{4, 1, 2, 1, 2}
	single := singleNumber(inputData)
	fmt.Println(single)
}
