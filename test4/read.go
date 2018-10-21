package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
type Name struct {
	fname string
	lname string
}
func main() {
	file, err := os.Open("test.txt")
	check(err)

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		name := Name{fname: "", lname: ""}
		name.fname = strings.Split(content, " ")[0]
		name.lname = strings.Split(content, " ")[1]
		fmt.Println(name)
		names = append(names, name)
	}
	if err := scanner.Err(); err != nil {
		check(err)
	}	
	
	fmt.Println(names)
	file.Close()
}