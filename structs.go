package main
import (
	"fmt"
)
type Person struct {
	name string
	addr string
	phone string
}
func main() {
	//p1 := new(Person)
	p1 := Person{name : "joe", addr: "abv", phone: "123"}
	p1.name = "mike"
	fmt.Println(p1)
}