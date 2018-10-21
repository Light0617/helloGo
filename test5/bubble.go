package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func Swap(arr []int, i int) {
	tmp := arr[i]
	arr[i] = arr[i+1]
	arr[i+1] = tmp
}
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr[:], j)
			}
		}
	}
}

func getArray(strs []string) []int {
	arr := []int{}
	for i, value := range strs {
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Input contains a not integer valued", err)
		}
		arr = append(arr, num)
		if i == 9 {
			break
		}
	}
	return arr
}

func processInput() []string {
	fmt.Println("Please input up to 10 numbers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	return strings.Split(s, " ")
}

func main() {
	strs := processInput()
	arr := getArray(strs)

	fmt.Println("before sorted: ", arr);
	BubbleSort(arr[:])
	fmt.Println("after sorted: ", arr);
}