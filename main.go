package main

import (
	"fmt"
	"github.com/alanchenchen/testGo/appendInt"
	_ "github.com/alanchenchen/testGo/array"
)

type Man struct {
	Name string
	Age int
}

func (a Man) show(msg string) {
	s := make([]int, 1)
	s1 := appendInt.AppendInt(s, 1)

	fmt.Println(a.Name, msg)
	fmt.Println(s1)
}

func main() {
	man := Man {
		Name: "alan",
		Age: 25,
	}

	man.show("hello")
}