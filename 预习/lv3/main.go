package main

import "fmt"

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

// go run main.go < data
func main() {
	nums := make([]int, 0)
	var n int
	for {
		k, _ := fmt.Scanf("%v", &n)
		if k != 1 {
			break
		} else {
			nums = append(nums, n)
		}
	}

	quickSort(nums, 0, 99)
	for _, v := range nums {
		println(v)
	}
}
