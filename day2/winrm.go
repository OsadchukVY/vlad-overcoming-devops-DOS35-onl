package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"github.com/masterzen/winrm"
	"log"
	"os"
)

var (
	server   string
	user     string
	password string
	port     int
)

func init() {
	flag.StringVar(&server, "server", "127.0.0.1", "set  destination server IP")
	flag.StringVar(&user, "user", "", "Set username")
	flag.IntVar(&port, "port", 5985, "Set port")
}

func winrm_connect(server, user, password, command string, port int) {
	endpoint := winrm.NewEndpoint(server, port, false, false, nil, nil, nil, 0)
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

func enterPassword() (string, error) {
	fmt.Println("Enter password :>")
	inputReader := bufio.NewReader(os.Stdin)
	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return userInput, nil
}

func main() {
	var password, err = enterPassword()
	if err != nil {
		log.Fatal(err)
	}

	for {

		command, err := userInput()

		if err != nil {
			log.Fatal(err)
		}

		winrm_connect(server, user, password, command, port)
		fmt.Println(command)

	}
}
