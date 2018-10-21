package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter up to ten integers separated by spaces:")
	scnr := bufio.NewScanner(os.Stdin)
	scnr.Scan()
	s := scnr.Text()
	a := strings.Split(s, " ")
	slc := []int{}
	for i, v := range a {
		j, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Input contains a not integer valued", err)
		}
		slc = append(slc, j)
		if i == 9 {
			break
		}
	}
	BubbleSort(slc)
	fmt.Print("Sorted numbers: ")
	fmt.Println(slc)
}

func BubbleSort(slc []int) {
	n := len(slc)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if slc[i] > slc[i+1] {
				Swap(slc, i)
				swapped = true
			}
		}
	}
}

func Swap(slc []int, i int) {
	slc[i+1], slc[i] = slc[i], slc[i+1]
}