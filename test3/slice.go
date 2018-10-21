package main
import (
	"fmt"
	"strconv"
	"sort"
)
func main() {
	arr := make([]int, 3)
	fmt.Println(arr)
	var input string
	fmt.Println("please input until X...");
	count := 0
	for input != "X" {
		fmt.Scan(&input)
		if input == "X" {break}
		num, _ := strconv.Atoi(input)
		if count < 3 {
			arr[count] = num
		} else {
			arr = append(arr, num)
		}
		count++
		sli := arr[0: count]
		sort.Ints(sli)
		fmt.Println("sorted array: ", sli)
	}
}