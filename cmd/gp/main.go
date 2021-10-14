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
		WD/cmd/NAME/main.go			<- overide with -file
		WD/cmd/NAME/main_test.go	<- overide with -file

*/

var debug bool

func main() {
	var namePack string
	var nameFile string
	var enableTop bool
	var enableCmd bool
	var enableTst bool
	var enableTyp bool
	var useForce bool
	var onlyVersion bool

	flag.StringVar(&namePack, "name", "", "package name")
	flag.StringVar(&nameFile, "file", "", "filename to use")

	flag.BoolVar(&enableTop, "top", false, "top-level package => no subdirectory created")
	flag.BoolVar(&enableCmd, "cmd", false, "command package => special structure created")
	flag.BoolVar(&enableTst, "test", true, "generate package test")
	flag.BoolVar(&enableTyp, "data", true, "generate type and func")
	flag.BoolVar(&useForce, "force", false, "overwrite existing file(s)")

	flag.BoolVar(&onlyVersion, "version", false, "show version information and exit")
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	if onlyVersion {
		fmt.Printf("gp version: %s\n", gp.Version())
		return
	}

	if namePack == "" {
		panic(errors.New("no package name specified"))
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var baseDir string
	var baseName string
	var filenameCode string
	var filenameTest string
	var packageName string

	switch nameFile {
	case "":
		switch enableCmd {
		case true:
			baseName = "main"
		default:
			baseName = namePack
		}
	default:
		baseName = nameFile
	}

	switch {
	case enableTop:
		baseDir = wd
		packageName = namePack
	case enableCmd:
		baseDir = filepath.Join(wd, "cmd", namePack)
		packageName = "main"
	default:
		baseDir = filepath.Join(wd, namePack)
		packageName = namePack
	}

	filenameCode = fmt.Sprintf("%s.go", baseName)
	filenameTest = fmt.Sprintf("%s_test.go", baseName)

	if debug {
		fmt.Fprintf(os.Stderr, "     baseDir: %q\n", baseDir)
		fmt.Fprintf(os.Stderr, "filenameCode: %q\n", filenameCode)
		fmt.Fprintf(os.Stderr, "filenameTest: %q\n", filenameTest)
		fmt.Fprintf(os.Stderr, " packageName: %q\n", packageName)
	}

	info := gp.New(packageName, namePack, enableCmd, enableTst, enableTyp)
	textCode, err := info.CreatePackageCode()
	if err != nil {
		panic(err)
	}

	if err := write(baseDir, filenameCode, textCode, useForce); err != nil {
		panic(err)
	}

	if !enableTst {
		return
	}

	textTest, err := info.CreatePackageTest()
	if err != nil {
		panic(err)
	}

	if err := write(baseDir, filenameTest, textTest, useForce); err != nil {
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
