package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS: %s.\n", getOs())
}

func getOs() string {
	defer fmt.Println("You got it!")
	os := runtime.GOOS
	return os
}
