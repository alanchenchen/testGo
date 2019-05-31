package appendInt

import "fmt"

func AppendInt(x []int, y int) []int {
	var clone []int
	xLen := len(x)
	targetLen := xLen + 1
	xCap := cap(x)
	if targetLen <= xCap {
		fmt.Println("当前切片容量足够扩展一个元素的长度")
		clone = x[:targetLen]
	} else {
		fmt.Println("当前切片容量不够扩展一个元素的长度, 需要扩容")
		clone = make([]int, targetLen, xLen*2)
		copy(clone, x)
	}
	clone[xLen] = y
	return clone
}
