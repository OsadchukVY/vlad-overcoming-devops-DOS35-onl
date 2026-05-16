package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func proc_file(file string) {
	f_stat, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	out_string := fmt.Sprintf("Size: %d Path: %s", f_stat.Size(), file)

	fmt.Println(out_string)
}

func main() {
	args, _ := script.Args().Slice()
	if len(args) != 2 {
		log.Fatal("Accepting only 2 argument")
	}

	regex := regexp.MustCompile(args[0])
	files, err := script.FindFiles(args[1]).Slice()

	if err != nil {
		log.Fatal(err)
	}
	for _, v := range files {

		str, err := script.File(v).MatchRegexp(regex).String()
		if err != nil {
			log.Fatal(err)
		}

		if len(str) != 0 {
			proc_file(v)
		}
	}
}
