package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func userInput() {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		userInput, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(userInput)
	}

}
func main() {
	userInput()
}
