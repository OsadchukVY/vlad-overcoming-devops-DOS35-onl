package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := fmt.Println("Hello, world")
	if err != nil {
		log.Fatal(err)
	}
}
