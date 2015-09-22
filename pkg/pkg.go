package pkg

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

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
		Prebuild  []string
		Features  [][]string
		Configure []string
		Extra     []string
		Postbuild []string
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
	// get necessary info
	p := prepare(recipe)
	// destination file
	df := p.Info.Name + "#" + p.Info.Version + ".pkg.tgz"

	// create work dir
	wd, err := ioutil.TempDir(os.TempDir(), "surt-")
	check(err)
	defer os.Remove(wd)

	// open work dir
	dir, err := os.Open(wd)
	check(err)
	defer dir.Close()

	// TODO: add actual building logic here (make, cmake, configure array
	// parsing, features array parsing, prebuild commands, post build commands,
	// downloading, etc. (probably in another function)

	// read files in dir
	fl, err := dir.Readdir(0)
	check(err)

	// create destinaton tarball file
	tf, err := os.Create(df)
	check(err)
	defer tf.Close()

	// create file writers and closers in preparation for compression
	var fw io.WriteCloser = tf

	// create gzip writer, all writes here are now compressed
	fw = gzip.NewWriter(tf)
	defer fw.Close()

	// create tarball file writer
	tfw := tar.NewWriter(fw)
	defer tfw.Close()

	// add files to tarball
	for _, fi := range fl {
		if fi.IsDir() {
			continue
		}
		// open file to be written to tarball
		file, err := os.Open(dir.Name() + string(filepath.Separator) + fi.Name())
		check(err)
		defer file.Close()

		// create header info
		th := new(tar.Header)
		th.Name = file.Name()
		th.Size = fi.Size()
		th.Mode = int64(fi.Mode())
		th.ModTime = fi.ModTime()
		// write header
		check(tfw.WriteHeader(th))
		// write file to tarball
		_, err = io.Copy(tfw, file)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
