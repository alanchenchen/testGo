package testGo

import (
	"fmt"
	"testGo/appendInt"
	_ "testGo/array"
)

type Man struct {
	Name string
	Age int
}

func (a Man) show(msg string) {
	s := make([]int, 0)
	s1 := appendInt.AppendInt(s, 1)

	fmt.Println(a.Name)
	fmt.Println(s1)
}