package main

import (
	"fmt"
	"os"

	"github.com/kori/surt/pkg"
)

func main() {
	switch os.Args[1] {
	default:
		fmt.Println("no operation specified")
	case "build":
		file := os.Args[2]
		pkg.Build(file)
	case "add":
		fmt.Println("not implemented yet!")
	}
}
