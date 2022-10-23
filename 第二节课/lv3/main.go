package main

import "fmt"

// 冒泡排序

func main() {
	ints := []int{1, 2, 3, 5, 67, 8, 5, 3, 1, -2}
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-1; j++ {
			if ints[j] > ints[j+1] {
				temp := ints[j]
				ints[j] = ints[j+1]
				ints[j+1] = temp
			}
		}
	}
	for _, v := range ints {
		fmt.Println(v)
	}
}
