package templates

import (
	_ "embed"
)

var (
	//go:embed resources/code.go.template
	CodeTemplate string

	//go:embed resources/test.go.template
	TestTemplate string
)
