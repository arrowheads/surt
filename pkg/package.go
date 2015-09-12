package pkg

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type Package struct {
	Source struct {
		Source  string
		SHA256  []string
		Patches []string
	}
	Build struct {
		Static    bool
		Parallel  bool
		Configure []string
		Extra     []string
	}
	Info struct {
		Name        string
		Version     string
		Compression string
		Strip       bool
	}
}

func Prepare(recipe string) Package {
	// create work dir
	wd, err := ioutil.TempDir(os.TempDir(), "surt-")
	check(err)
	defer os.Remove(wd)

	r, err := os.Open(recipe)
	check(err)

	buf, err := ioutil.ReadAll(r)
	check(err)
	r.Close()

	var pkg Package
	check(toml.Unmarshal(buf, &pkg))
	return pkg
}

func main() {
	file := os.Args[1]
	ay := Prepare(file)
	fmt.Println(ay.Info.Name)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
