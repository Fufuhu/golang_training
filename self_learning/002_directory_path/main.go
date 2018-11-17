package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	exe, err := os.Executable()

	if err != nil {
		panic(err)
	}

	directory := filepath.Dir(exe)

	fmt.Printf("フルパス:%s\n", exe)
	fmt.Printf("ディレクトリ:%s\n", directory)
}
