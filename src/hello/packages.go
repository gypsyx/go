package main

import (
	"fmt"
	"testpkg"
)

func main() {
	fmt.Printf("This is package main")
	testpkg.PrintName()
}