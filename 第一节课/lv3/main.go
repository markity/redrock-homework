package main

import (
	"math/rand"
	"time"
)

func quickSort(nums []int, low int, high int) {
	if low < high {
		pivotloc := partition(nums, low, high)
		quickSort(nums, low, pivotloc-1)
		quickSort(nums, pivotloc+1, high)
	}
}

func partition(nums []int, low int, high int) int {
	pivotkey := nums[low]
	for low < high {
		for low < high && nums[high] >= pivotkey {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= pivotkey {
			low++
		}
		nums[high] = nums[low]
	}

	nums[low] = pivotkey
	return low
}

func main() {
	rand.Seed(time.Now().Unix())
	nums := make([]int, 100)
	cur := 0
	for i := 0; i < 100; i++ {
		cur = rand.Intn(3000)
		nums[i] = cur
		rand.Seed(int64(cur) + time.Now().Unix())
	}
	quickSort(nums, 0, 99)
	for _, v := range nums {
		println(v)
	}
}
