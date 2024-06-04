package main

import (
	"context"
	"os"
	"mrwill/src/templates/components"
)

func main() {
    comp := templates.Hello("name")
    comp.Render(context.Background(), os.Stdout)
}
