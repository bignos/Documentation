package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	var err error
	if cwd, err = os.Getwd(); err != nil {
		log.Fatal("os.Getwd failed: %v", err)
	}
}

func main() {
	fmt.Printf("CWD: %s\n", cwd)
}
