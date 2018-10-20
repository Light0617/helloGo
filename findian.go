package main
import "fmt"
import "strings"

func checkInput(input string) {
	fmt.Println("input= %s", input);
	var found bool = strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Index(input, "a") >= 0;
	if found {
		fmt.Println("Found!");
	} else {
		fmt.Println("Not Found!");
	}
}

func main() {
	for i:=0; i < 10; i ++ {
		var input string;
		fmt.Println("type the string");
		fmt.Scan(&input)
		fmt.Println("input= ", input);
		checkInput(strings.ToLower(input));
	}
}