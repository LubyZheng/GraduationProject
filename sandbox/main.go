package main

import (
	"Gproject/sandbox/judger"
	"fmt"
	"os"
)

func main() {
	l, _ := os.Getwd()
	fmt.Println(l)
	judger.Judge()
}
