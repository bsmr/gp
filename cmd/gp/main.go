package main

import (
	"fmt"
	"os"

	"github.com/bsmr/gp"
)

func main() {
	svc := gp.New()
	if err := svc.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s!\n", err)
	}
}
