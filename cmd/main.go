package main

import (
    "context"
    "os"
    "mrwill/src/templates" 
)

func main() {
    comp := templates.Hello("Chris")
    comp.Render(context.Background(), os.Stdout)
}
