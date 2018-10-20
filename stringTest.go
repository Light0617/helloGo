package main
import "fmt"
/*
IsDigit
IsSpace
IsLetter
IsLetter
IsPunct
ToUpper
ToLower
Compare(a, b)
Contains(s, sub)
HasPrefix(s, pre)
Index(s, sub)
Replace(s, old, new, n)
TrimSpace(s)
Atoi(s)
Itoa(s)
ParseFloat(s, bitsize)
*/
type Grades int
const (
	A Grades = iota
	B
	C
	D
	E
)
func main() {
	var x int = 6;
	if x > 5 {
		fmt.Printf("yup\n");
	}
	for i:=0; i < 10; i ++ {
		fmt.Printf("hy\n");
	}
	switch x {
	case 1:
		fmt.Printf("case1\n");
	case 2:
		fmt.Printf("case2 \n");
	default:
		fmt.Printf("nocase\n");
	}
	switch  {
	case x > 1:
		fmt.Printf("case1\n");
	case x > 2:
		fmt.Printf("case2 \n");
	default:
		fmt.Printf("nocase\n");
	}
	i := 0
	for i < 10 {
		i++
		if i == 5 {break}
		fmt.Printf("no!")
	}
	var content string = "Hello, world!";
	fmt.Printf(content + "\n");

	var appleNum int
	fmt.Printf("number of apples?")
	fmt.Scan(&appleNum)
	fmt.Printf("apple= %d\n", appleNum);
}