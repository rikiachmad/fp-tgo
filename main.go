package main

import (
	"fmt"
)

var arr = []int{4, 1, 13, 7, 0, 2, 8, 11, 3}
var ans = make([][]int, 0)
var temp = make([]int, 0)
var max = 0

func LongestIncreasingSubsequence(start, num int) {
	temp = append(temp, num)
	for i := start; i < len(arr); i++ {
		if arr[i] > num {
			LongestIncreasingSubsequence(i, arr[i])
		}
	}
	var temp2 []int
	temp2 = append(temp2, temp...)
	length := len(temp2)
	if length > max {
		max = length
	}
	ans = append(ans, temp2)
	temp = temp[:len(temp)-1]
}

func main() {
	for i := 0; i < len(arr); i++ {
		LongestIncreasingSubsequence(i, arr[i])
	}
	fmt.Printf("Longest Monotonically Increasing Subsquence is %d\n", max)
	for i := 0; i < len(ans); i++ {
		if len(ans[i]) == max {
			fmt.Println(ans[i])
		}
	}
}
