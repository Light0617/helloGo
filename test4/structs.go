package main
import (
	"fmt"
	"encoding/json"
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

	barr, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(barr)
	var p2 Person
	err2 := json.Unmarshal(barr, &p2) 
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	fmt.Println(p2)


	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	fmt.Println("jsonBlob:", jsonBlob)
	err3 := json.Unmarshal(jsonBlob, &animals)
	if err3 != nil {
		fmt.Println("error:", err3)
	}
	fmt.Println(animals)
}