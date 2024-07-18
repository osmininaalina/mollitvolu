package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
