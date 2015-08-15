package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := ioutil.TempDir(os.TempDir(), "surt-")
	defer os.Remove(file)
	fmt.Println(file)
}
