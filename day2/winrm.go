package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func userInput() (string, error) {
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return userInput, nil
}
func main() {

	for {
		command, err := userInput()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(command)
	}
}
