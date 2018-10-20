package main

import "fmt"
var y = 3;
func f() {
	var x = 4;
	fmt.Printf("%d", x);
	fmt.Printf("%d", y);
}
func g() {
	var x = 4;
	fmt.Printf("%d", x);
	fmt.Printf("%d", y);
}
func foo() *int {
	x := 1
	return &x
}
func main_foo() {
	var y *int
	y = foo()
	fmt.Printf("%d\n", *y);
}
func main() {
	main_foo();
	var content string = "Hello, world!";
	fmt.Printf(content + "\n");
}
