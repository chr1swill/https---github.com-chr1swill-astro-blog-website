package main

import (
	"context"
	"log"
    "github.com/a-h/templ"
	"mrwill/src/templates/views"
	"mrwill/src/templates/layouts"
	"os"
)

type MainLayoutParams struct {
    title string
    description string
    pageUrl string
    content templ.Component 
    keywords []string
}

func main() {
    if _, err := os.Stat("dist"); os.IsNotExist(err) {
        err := os.Mkdir("dist", os.ModePerm)
        if err != nil {
            log.Fatalf("Could not create dist: %v", err)
        }
    }

    homePageFile, err := os.Create("dist/index.html")
    if err != nil {
        log.Fatalf("Could not create dist/index.html: %v", err)
    }
    defer homePageFile.Close()

    hpp := MainLayoutParams{
        title: "Website That feel Native | Christian Williams",
        description: "I create Website that provide a rich experience to your customer that will convert into sales for you buisness",
        pageUrl: "https://mrwill.ca",
        content: views.HomePage(),
        keywords: []string{"websites", "user experience", "conversions"},
    }

    err = layouts.MainLayout(hpp.title, hpp.description, hpp.pageUrl, hpp.content, hpp.keywords).Render(context.Background(), homePageFile)
    if err != nil {
        log.Fatalf("Could not render template to dist/index.html: %v", err)
    }
}
