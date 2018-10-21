package main
import (
	"fmt"
	"math"
)
func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}
//two parameters: afunct, val
func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}

func MakeDistOrigin (o_x, o_y float64) func (float64, float64) float64 {
	fn := func (x, y float64) float64 {
		return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
	}
	return fn
}

func getMax(vals ...int) int {
	maxV := -1 
	for _, v := range vals {
		if v > maxV { 
			maxV = v
		}
	}
	return maxV
}

func main() {
	var funcVar func(int) int
	funcVar = incFn
	fmt.Println(funcVar(1))
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))

	v := applyIt(func (x int)int {return x + 10}, 2)
	fmt.Println(v)

	//get the distance to origin
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))

	fmt.Println(getMax(1, 3, 6, 4))
	vslice := []int {1, 3, 6, 4}
	fmt.Println(getMax(vslice...))

	defer fmt.Println("bye!")
	fmt.Println("hi")

	i := 1
	defer fmt.Println(i+1)
	i++
	fmt.Println("hello")

}