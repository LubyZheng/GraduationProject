package main

import (
	"Gproject/tool"
	"Gproject/tool/flag"
	"os"
)

func main() {
	f := flag.New()
	f.Parse(os.Args[1:])
	code := tool.New(f)
	err := code.Run()
	if err != nil {
		os.Exit(1)
	}
}
