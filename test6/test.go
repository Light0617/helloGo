package main
import (
	"fmt"
)
var f func(string) int
func test(x string) int {
	return len(x)
}
func fA() func() int {
	i := 0
	return func() int {
			i++
			return i
	}
}

func main() {
	f = test
	fmt.Println(f("1234"))

	fB := fA()
  fmt.Println(fB())
	fmt.Println(fB())
	
	i := 1
	fmt.Print(i)
	i++
	defer fmt.Print(i+1)
	fmt.Print(i)
}