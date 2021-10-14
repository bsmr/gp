package gp

import (
	"bytes"
	"go/format"
	"html/template"
)

const (
	codeTemplate = "package {{ .Name }}"
	testTemplate = "package {{ .Name }}"
)

type Information struct {
	Name string
}

func New(name string) Information {
	return Information{
		Name: name,
	}
}

func (i Information) CreatePackageCode() (string, error) {
	return i.CreatePackage("code", codeTemplate)
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
