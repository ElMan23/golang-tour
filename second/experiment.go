package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Printf("OS: %s.\n", os)
}
