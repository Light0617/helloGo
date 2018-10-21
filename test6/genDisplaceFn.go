package main
import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)
func processInput() (float64, float64, float64, float64) {
	msg := "Please input Acceleration a, Initial Velocity v0, Initial Displacement s0, Enter the value of time t with separated by spaces:"
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	strs := strings.Split(s, " ")
	arr := []float64{}
	for i, value := range strs {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("Input contains a no-float64", err)
		}
		arr = append(arr, num)
		if i == 4 {
			break
		}
	}
	if len(arr) < 4 {
		log.Fatal("the number of input is not enough")
	}
	a, v0, s0, t := arr[0], arr[1], arr[2], arr[3]
	return a, v0, s0, t
}

func GenDisplaceFn(a, v0, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}
	return fn
}

func main() {
	a, v0, s0, t := processInput()
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("the Displacement is = ", fn(t))
}