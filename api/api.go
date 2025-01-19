package main

import (
	"fmt"
	"os"
)

func main() {
	if err := mainRun(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mainRun() error {
	fmt.Println("Starting...")
	return nil
}
