package main

import (
	"Gproject/judge"
	"fmt"
	"os"
)

func main() {
	f := judge.NewFlag()
	f.Parse(os.Args[1:])
	code, err := judge.New(f)
	if err != nil {
		os.Exit(1)
	}
	err = code.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
