package testGo

import (
	"fmt"
	"github.com/alanchenchen/testGo/appendInt"
	_ "github.com/alanchenchen/testGo/array"
)

type Man struct {
	Name string
	Age  int
}

func (a Man) Show(msg string) {
	s := make([]int, 1)
	s1 := appendInt.AppendInt(s, 1)

	fmt.Println(a.Name)
	fmt.Println(s1)
}
