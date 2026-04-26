package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	// Execute "ls -l" (on Unix) or "dir" (on Windows)
	out, err := exec.Command("sha256sum", "blazingly_fast.go").Output()
	if err != nil {
		log.Fatal(err)
	}

	split := string(out)
	fmt.Println(strings.Split(split, " ")[0])
}
