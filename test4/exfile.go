package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, e := ioutil.ReadFile("ftest.txt")
	if e != nil {
		fmt.Println("error:", e)
	}
	fmt.Println(string(data))

	dat := []byte("hello bro")
	//0777 is permision
	err5 := ioutil.WriteFile("ftest2.txt", dat, 0777)
	if err5 != nil {
		fmt.Println("error:", err5)
	}

	//os
	//read
	f, err0 := os.Open("ftest.txt")
	if err0 != nil { fmt.Println("error:", err0)}
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	if err != nil { fmt.Println("error:", err)}
	fmt.Println("nb:", nb)
	f.Close()

	//write
	infile, err2 := os.Create("ftest2.txt")
	if err2 != nil { fmt.Println("error:", err2)}
	barr = []byte{1, 2, 3}
	nb, err = infile.Write(barr)
	nb, err = infile.WriteString("Hi")
}