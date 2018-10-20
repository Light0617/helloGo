package main
import (
	"fmt"
)
/*
len()
cap()
*/
func main() {
	var x [5]int
	x[0] = 2
	fmt.Println(x);
	var y [5]int = [5]int{1, 2, 3, 4 ,5}
	fmt.Println(y);
	z := [3]int {1,2,3,}
	for i, v := range z {
		fmt.Printf("ind %d, val %d\n", i, v);
	}
	s1 := y[1:3]
	s2 := y[2:5]
	fmt.Println(s1);
	fmt.Println(s2);
	fmt.Println(len(s1));
	fmt.Println(cap(s1));
	s1[0] = 100
	fmt.Println(s1);
	fmt.Println(y);
	sli := make([]int, 10)
	fmt.Println(sli);
	sli3 := make([]int, 3, 4)
	fmt.Println(sli3);
	sli3 = append(sli3, 4)
	fmt.Println(sli3)
	sli = append(sli, 6)
	fmt.Println(sli)

	//hash table
	// var idMap1 map[string]int
	// idMap1 = make(map[string]int)
	idMap := map[string]int {"joe": 123}
	fmt.Println(idMap["joe"])
	idMap["joe"] = 456
	fmt.Println(idMap["joe"])
	//delete(idMap, "joe")
	id, p := idMap["joe"]
	fmt.Println(id, p)
	for key, val := range idMap {
		fmt.Println(key, val)
	}
}