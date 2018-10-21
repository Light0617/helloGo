package main
import ("fmt")
func foo2(x int) (int, int) {
	return x, x + 1
}
func foo(x int) int {
	x = x + 2
	return x
}
func foo3(x *int) {
	*x = *x + 2
}
func slow_foo(x [3]int) int {
	return x[0]
}
func fast_messy_foo(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}
func fast_foo(sli []int) {
	sli[0] = sli[0] + 10
}
func main() {
	a0, b0 := foo2(3)
	fmt.Println(a0, b0)
	x := 2
	fmt.Println(foo(x))
	fmt.Println(x)
	foo3(&x)
	fmt.Println(x)

	a := [3]int{1, 2, 3}
	fmt.Println(slow_foo(a))
	fast_messy_foo(&a)
	fmt.Println(a)

	b := []int{1, 2, 3}
	fast_foo(b[:])
	fmt.Println(b)
}