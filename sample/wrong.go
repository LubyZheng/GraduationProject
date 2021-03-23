package main

import "fmt"

func main() {
	var a [1]int
	for i := 1; i >= 0; i-- {
		fmt.Println(a[i])
	}
}
