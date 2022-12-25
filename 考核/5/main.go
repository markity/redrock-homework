package main

import (
	"fmt"
	"math"
)

const n = 6

func main() {
	for i := 1; i <= n; i++ {
		do(i)
	}
}

// 一位数 0 ~ 9			10
// 二位数 10 ~ 99		90
// 三位数 100~999		900
// 四位数 1000~9999 	9000
// 五位数 10000~99999	90000
func do(n int) {
	var t int
	if n == 1 {
		t = 10
	} else {
		t = 9 * int(math.Pow10(n-1))
	}

	var start = int(math.Pow10(n - 1))

	for i := 0; i < t; i++ {
		if check(start+i, n) {
			fmt.Println(start + i)
		}
	}
}

func check(num int, n int) bool {
	origin := num
	result := float64(0)
	for {
		remain := num % 10
		num /= 10

		result += math.Pow(float64(remain), float64(n))
		if num == 0 {
			break
		}
	}

	return int(result) == origin
}
