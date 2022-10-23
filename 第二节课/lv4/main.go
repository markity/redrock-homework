package main

import (
	"fmt"
	"sort"
)

type StringListMapSorted []string

func (lis *StringListMapSorted) Len() int {
	return len(*lis)
}

func (lis *StringListMapSorted) Swap(i, j int) {
	temp := (*lis)[i]
	(*lis)[i] = (*lis)[j]
	(*lis)[j] = temp
}

func (lis *StringListMapSorted) Less(i, j int) bool {
	return (*lis)[i] < (*lis)[j]
}

func main() {
	var lis StringListMapSorted
	lis = append(lis, "你好")
	lis = append(lis, "世界")
	lis = append(lis, "不如")
	lis = append(lis, "来玩")
	lis = append(lis, "沙")

	sort.Sort(&lis)

	fmt.Println(lis)
}
