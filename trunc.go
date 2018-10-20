package main
import "fmt"

func main() {
	var num float32
	fmt.Printf("second float=?")
	fmt.Scan(&num)
	fmt.Println("first float is= %d", int(num));
	
	fmt.Printf("second float=?")
	fmt.Scan(&num)
	fmt.Println("second float is= %d", int(num));
}
