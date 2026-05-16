package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func fileCopy(v string) {
	src, err := os.Open(v)
	if err != nil {
		log.Fatal(err)
	}

	defer src.Close()

	dst, err := os.Create(v + ".bak")

	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(src, dst)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	args, _ := script.Args().Slice()
	if len(args) != 1 {
		log.Fatal("1 argument, yes/no")
	}
	replace_str := fmt.Sprintf("%s %s", "PasswordAuthentication", args[0])
	regex := regexp.MustCompile("#PasswordAuthentication.*")

	ssh_files, err := script.FindFiles("/etc/ssh/").Slice()

	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range ssh_files {

		//fileCopy(v)
		auth, err := script.File(v).ReplaceRegexp(regex, replace_str).Stdout()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(auth)
		//fileCopy(v)
	}
}
