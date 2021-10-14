package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bsmr/gp"
)

/*
	TOP:
		WD/NAME.go
		WD/NAME_test.go
	NORMAL:
		WD/NAME/NAME.go
		WD/NAME/NAME_test.go
	COMMAND:
		WD/cmd/NAME/main.go
		WD/cmd/NAME/main_test.go

*/

var debug bool

func main() {
	var top bool
	var cmd bool
	var name string
	var test bool
	var force bool

	flag.BoolVar(&top, "top", false, "top-level package => no subdirectory created")
	flag.BoolVar(&cmd, "cmd", false, "command package => special structure created")

	flag.StringVar(&name, "name", "", "package name")
	flag.BoolVar(&test, "test", true, "generate package test")

	flag.BoolVar(&force, "force", false, "overwrite existing file(s)")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	if name == "" {
		panic(errors.New("no package name specified"))
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var baseDir string
	var filenameCode string
	var filenameTest string
	var packageName string

	switch {
	case top:
		baseDir = wd
		filenameCode = fmt.Sprintf("%s.go", name)
		filenameTest = fmt.Sprintf("%s_test.go", name)
		packageName = name
	case cmd:
		baseDir = filepath.Join(wd, "cmd", name)
		filenameCode = "main.go"
		filenameTest = "main_test.go"
		packageName = "main"
	default:
		baseDir = filepath.Join(wd, name)
		filenameCode = fmt.Sprintf("%s.go", name)
		filenameTest = fmt.Sprintf("%s_test.go", name)
		packageName = name
	}

	if debug {
		fmt.Fprintf(os.Stderr, "     baseDir: %q\n", baseDir)
		fmt.Fprintf(os.Stderr, "filenameCode: %q\n", filenameCode)
		fmt.Fprintf(os.Stderr, "filenameTest: %q\n", filenameTest)
		fmt.Fprintf(os.Stderr, " packageName: %q\n", packageName)
	}

	info := gp.New(packageName, name, cmd)
	textCode, err := info.CreatePackageCode()
	if err != nil {
		panic(err)
	}

	if err := write(baseDir, filenameCode, textCode, force); err != nil {
		panic(err)
	}

	if !test {
		return
	}

	textTest, err := info.CreatePackageTest()
	if err != nil {
		panic(err)
	}

	if err := write(baseDir, filenameTest, textTest, force); err != nil {
		panic(err)
	}
}

func write(dir, name, text string, force bool) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fp := filepath.Join(dir, name)
	fi, _ := os.Stat(fp)

	switch {
	case (fi != nil) && (fi.IsDir()):
		return fmt.Errorf("%q is a directory", fp)
	case (fi != nil) && !force:
		return fmt.Errorf("%q exists", fp)
	default:
		return os.WriteFile(fp, []byte(text), 0644)
	}
}
