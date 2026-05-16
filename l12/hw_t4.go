package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	args, _ := script.Args().Slice()
	if len(args) != 3 {
		log.Fatal("Less then 3 arguments")
	}

	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])

	_, err := script.File(args[0]).Stdout()

	if err != nil {
		log.Fatal(err)
	}

	regexp_string := fmt.Sprintf("%s%s%s", ".*", args[2], "$")
	regex := regexp.MustCompile(regexp_string)
	_, err = script.FindFiles(args[1]).MatchRegexp(regex).Stdout()

	if err != nil {
		log.Fatalln(err)
	}
}
