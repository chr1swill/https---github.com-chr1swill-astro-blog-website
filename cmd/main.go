package main

import (
	"context"
	"os"
    "mrwill/src/templates/views"
)

func main() {
    hp := views.HomePage("blac", "blag", "blag", "blag")
    hp.Render(context.Background(), os.Stdout)
}
