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
		recipe := os.Args[2]
		r := pkg.Prepare(recipe)
		r.Build()
	case "add":
		fmt.Println("not implemented yet!")
	}
}
