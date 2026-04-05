package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userInput)
}
