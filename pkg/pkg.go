package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/naoina/toml"
)

type Package struct {
	Info struct {
		Name         string
		Version      string
		Dependencies []string
	}
	Source struct {
		Source  string
		SHA256  []string
		Patches []string
	}
	Build struct {
		Strip     bool
		Static    bool
		Parallel  bool
		PreBuild  []string
		Features  [][]string
		Configure []string
		Extra     []string
		PostBuild []string
	}
}

func prepare(recipe string) Package {
	r, err := os.Open(recipe)
	check(err)

	buf, err := ioutil.ReadAll(r)
	check(err)
	r.Close()

	var pkg Package
	check(toml.Unmarshal(buf, &pkg))
	return pkg
}

func Build(recipe string) {
	p := prepare(recipe)
	// destination file
	df := p.Info.Name + "#" + p.Info.Version + ".pkg.tar"

	// create work dir
	wd, err := ioutil.TempDir(os.TempDir(), "surt-")
	check(err)
	defer os.Remove(wd)
	fmt.Println(wd)
	time.Sleep(10 * time.Second)

	dir, err := os.Open(wd)
	check(err)
	defer dir.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
