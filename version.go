package gp

var (
	// set version during build
	// -ldflags "-X github.com/bsmr/gp.VersionText=v1.2.3"
	// go build -ldflags "-X github.com/bsmr/gp.VersionText=$(git branch --show-current)" -o bin/gp cmd/gp/main.go
	VersionText = "*unspecified version*"
)

func Version() string {
	return VersionText
}
