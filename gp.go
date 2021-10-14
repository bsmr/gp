package gp

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

var (
	//go:embed resources/code.go.template
	codeTemplate string

	//go:embed resources/test.go.template
	testTemplate string
)

type Information struct {
	Name  string
	Title string
	Main  bool
	Test  bool
	Data  bool
}

func New(name, title string, main, test, data bool) Information {
	return Information{
		Name:  name,
		Title: title,
		Main:  main,
		Test:  test,
		Data:  data,
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

	data := map[string]interface{}{
		"Name":  i.Name,
		"Title": i.Title,
		"Main":  i.Main,
		"Test":  i.Test,
		"Data":  i.Data,
		"Type":  firstUpper(i.Name),
	}
	var sb bytes.Buffer
	if err := t.Execute(&sb, data); err != nil {
		return "", err
	}

	buf, err := format.Source(sb.Bytes())
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func firstUpper(text string) string {
	if text == "" {
		return text
	}

	f := strings.ToUpper(text[:1])
	return fmt.Sprintf("%s%s", f, text[1:])
}
