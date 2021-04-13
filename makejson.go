package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	x := map[string]string{
		"name":    "",
		"address": "",
	}

	inputreader := bufio.NewReader(os.Stdin)
	print("Type name : ")
	name, _ := inputreader.ReadString('\n')
	print("Type Address : ")
	addr, _ := inputreader.ReadString('\n')

	name = strings.ReplaceAll(name, "\r\n", "")
	addr = strings.ReplaceAll(addr, "\r\n", "")

	x["name"] = name
	x["address"] = addr

	barr, _ := json.MarshalIndent(x, "", "  ")

	println("\n\nHere is the JSON : ")
	fmt.Println(string(barr))
}
