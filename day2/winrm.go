package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/masterzen/winrm"
	"log"
	"os"
)

// comment
func winrm_connect(server, user, password, command string) {
	endpoint := winrm.NewEndpoint(server, 5986, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, user, password)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client.RunWithContext(ctx, command, os.Stdout, os.Stderr)
}

func userInput() (string, error) {
	fmt.Println("shell :>")
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
