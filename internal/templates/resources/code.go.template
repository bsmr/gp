{{if .Name}}
package {{ .Name }}
{{if .Data}}

type {{ .Type }} struct {
}

func New() {{ .Type }} {
    return {{ .Type }}{}
}

func (s {{ .Type }}) String() string {
    return "<<Type is {{ .Type }}>>"
}
{{end}}
{{if .Main}}
{{if .Title}}

import "fmt"

func main() {
    fmt.Println("Hello, {{ .Title }}!")
}
{{end}}
{{end}}
{{end}}
