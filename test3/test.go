package main

import (
	"fmt"
	"os"
	"strconv"
)

// insert an int into a slice
func insertToSlice(nums []int, index int, x int) []int {
	if index <= len(nums) {
		nums = append(nums, 0)
		copy(nums[(index+1):], nums[index:])
		nums[index] = x
	}
	return nums
}

// insert an int into the correct position in a sorted slice of ints
func insertToSortedSlice(nums []int, x int) []int {
	if len(nums) < 1 {
		return append(nums, x)
	}

	low := 0
	hi := len(nums) - 1
	if x <= nums[low] {
		return insertToSlice(nums, low, x)
	}
	if x >= nums[hi] {
		return insertToSlice(nums, (hi + 1), x)
	}

	for mid := 0; hi-low > 1; {
		mid = low + (hi-low)/2
		if x < nums[mid] {
			hi = mid
		} else {
			low = mid
		}
	}
	return insertToSlice(nums, hi, x)
}

func main() {
	nums := make([]int, 0, 3)
	num := 0
	var err error
	for input := ""; ; {
		fmt.Println("Enter an Integer, 'X' to quit: ")
		if _, err = fmt.Scan(&input); err != nil {
			fmt.Printf("err: %v\n", err)
		} else if input == "X" {
			os.Exit(0)
		} else if num, err = strconv.Atoi(input); err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			nums = insertToSortedSlice(nums, num)
		}
		fmt.Printf("nums: %v\n", nums)
	}
}