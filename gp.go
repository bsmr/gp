package gp

import (
	"bytes"
	"go/format"
	"html/template"
)

const (
	codeTemplate = "package {{ .Name }}"
	mainTemplate = `package {{ .Name }}

	import "fmt"

	func main() {
		fmt.Println("Hello, {{ .Title }}!")
	}`
	testTemplate = "package {{ .Name }}"
)

type Information struct {
	Name  string
	Title string
	Main  bool
}

func New(name, title string, main bool) Information {
	return Information{
		Name:  name,
		Title: title,
		Main:  main,
	}
}

func (i Information) CreatePackageCode() (string, error) {
	switch i.Main {
	case true:
		return i.CreatePackage("code", mainTemplate)
	default:
		return i.CreatePackage("code", codeTemplate)
	}
}

func (i Information) CreatePackageTest() (string, error) {
	return i.CreatePackage("test", testTemplate)
}

func (i Information) CreatePackage(templateName, templateText string) (string, error) {
	t, err := template.New(templateName).Parse(templateText)
	if err != nil {
		return "", err
	}

	var sb bytes.Buffer
	if err := t.Execute(&sb, &i); err != nil {
		return "", err
	}

	buf, err := format.Source(sb.Bytes())
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
