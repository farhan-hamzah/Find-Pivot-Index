package main
import (
	"fmt"
)
func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	leftSum := 0
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	result := pivotIndex(nums)
	fmt.Println("Pivot Index:", result)
}