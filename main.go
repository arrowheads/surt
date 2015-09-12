package main

import (
	"fmt"
	"os"

	"github.com/kori/surt/pkg"
)

func main() {
	switch os.Args[1] {
	case "build":
		file := os.Args[2]
		p := pkg.Prepare(file)
		fmt.Println(p.Info.Name)
	case "add":
		fmt.Println("not implemented yet!")
	}
}
