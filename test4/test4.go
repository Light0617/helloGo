package main
import "fmt"
import "log"
import "os"
import "bufio"

type Name struct {
fname, lname string
}

func main() {
	var fileName string
	fmt.Printf("Enter the filename: ")
	fmt.Scan(&fileName)
	data, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
	}
	defer data.Close()
	data1 := make([]string,0)
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		data1 = append(data1, scanner.Text())
	}
	nameStructs := make([]Name,0)
	for i := 0; i < len(data1); i+=2 {
		data2 := Name{
			fname: data1[i],
			lname: data1[i+1],
		}
		nameStructs = append(nameStructs, data2)
	}
	for _, value := range nameStructs {
		fmt.Printf("%s\n",value)
	}
}