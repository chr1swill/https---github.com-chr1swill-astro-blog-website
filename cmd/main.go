package main

import (
	"context"
	"log"
	"mrwill/src/templates/views"
	"os"
)

func main() {
    if _, err := os.Stat("dist"); os.IsNotExist(err) {
        err := os.Mkdir("dist", os.ModePerm)
        if err != nil {
            log.Fatalf("Could not create dist: %v", err)
        }
    }

    homePage, err := os.Create("dist/index.html")
    if err != nil {
        log.Fatalf("Could not create dist/index.html: %v", err)
    }
    
    err = views.HomePage("blac", "blag", "blag", "blag").Render(context.Background(), homePage)
    if err != nil {
        log.Fatalf("Could not render template to dist/index.html: %v", err)
    }
}
