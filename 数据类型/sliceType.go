package main

import "fmt"

func main()  {
	var intArr [5]int = [...]int{1,22,33,66,99}
	slice := intArr[1:3]
	fmt.Println(slice) //intArr[1:3]切片可取到数组下标为1，2元素slice = {22,33}
	fmt.Println(cap(slice)) //内置函数cap为切片指向的内存空间的最大容量4（对应元素的个数，而不是字节数）
	slice = append(slice, 2) //cap=4
	slice = append(slice, 2) //cap=4
	slice = append(slice, 2) //cap=8
	slice = append(slice, 2)
	slice = append(slice, 2)
	slice = append(slice, 2)
	slice = append(slice, 2)
	fmt.Println(cap(slice))
}
