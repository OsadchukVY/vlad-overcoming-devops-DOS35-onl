package main

import (
	"github.com/bitfield/script"
)

func main() {
	pipe := script.Args()
	pipe.AppendFile("args.txt")
}
