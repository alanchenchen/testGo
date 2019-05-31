package array

import (
	"errors"
)

type callback func(a, b int, c []int) bool

func FindIndex(s []int, fn callback) int {
	index := -1 // 通过index为-1表示未找到对应的值
	for i, v := range s {
		status := fn(v, i, s)
		if status {
			index = i
		}
	}
	return index
}

func Find(s []int, fn callback) (int, error) {
	var err error // 因为s切片中可能存在负数，所以不能以此作为错误依据。单独返回一个error类型的值
	target := 0
	isExist := false

	for i, v := range s {
		status := fn(v, i, s)
		if status {
			target = v
			isExist = true
		}
	}
	if isExist == false {
		err = errors.New("couldn't find the corresponding value in target slice")
	}

	return target, err
}
