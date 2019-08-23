package main

import "fmt"

func fixKey(arr [][]string) (result [][]string) {
	var fixArr []string
	//var resultArr [][]string
	//s := arr
	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr[1]); j++ {
			fixArr = append(fixArr, arr[0][i]+arr[1][j])
		}
	}
	arr = append(arr[2:], fixArr)
	//fmt.Println("[arr]",arr)
	if len(arr) > 1 {
		result = fixKey(arr)
		fmt.Println("[if]", arr)
	} else {
		fmt.Println(arr)
		result = arr
		return result
	}
	return result
}



func fixKeyTwo(arr [][]string) [][]string {
	var fixArr []string
	nn := 1

	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr[1]); j++ {
			fixArr = append(fixArr, arr[0][i]+arr[1][j])
		}
	}
	arr = append(arr[2:], fixArr)
	fmt.Println(arr)

	if len(arr) > 1 {
		fmt.Println("fixKey前:", nn, arr)
		nn++
		fixKeyTwo(arr)
		fmt.Println("fixKey后:", nn, arr)

	} else {
		fmt.Println("fixKey中:", nn, arr)
		return arr

	}

	return arr
}


func main() {
	arr := [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		//{"j", "k", "l"},
	}

	fmt.Println("result:", fixKey(arr))
}
