package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main() {
	var name, address string
	scanner := bufio.NewScanner(os.Stdin)
	// Reads name and address from input
	fmt.Printf("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("Enter your address: ")
	scanner.Scan()
	address = scanner.Text()
	// Creates map and builds JSON
	var userMap map[string]string
	userMap = make(map[string]string)
	userMap["name"] = name
	userMap["address"] = address
	userJSON, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(userJSON))
}