package main

import "fmt"

// Pet 宠物店的宠物
type Pet struct {
	// Name 宠物的名字
	Name string
	// Type 宠物的品种
	Type string
	// Age 宠物的年龄
	Age int
}

func main() {
	pet1 := Pet{
		Name: "小白",
		Type: "哈士奇",
		Age:  3,
	}

	fmt.Println(pet1)
}
