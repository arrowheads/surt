package main

import (
	"fmt"
	"os"

	"github.com/kori/surt/pkg"
)

func main() {
	file := os.Args[1]
	ay := pkg.Prepare(file)
	fmt.Println(ay.Info.Name)
}
