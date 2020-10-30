package main

import (
	"fmt"
	"os"
	"flag"
	)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		fmt.Println(args)
		name := flag.String("app")
		fmt.Println(name)
	} else {
		fmt.Println("Enter Args ..!")
		os.Exit(0)
	}
}
