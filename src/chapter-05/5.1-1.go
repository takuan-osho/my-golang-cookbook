package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("somefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
