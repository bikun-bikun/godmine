package main

import (
	"fmt"
	"os"
)

func main() {
	exe, _ := os.Executable()
	fmt.Println(exe)
}
