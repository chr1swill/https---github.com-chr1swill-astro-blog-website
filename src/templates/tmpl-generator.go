package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func randomId() string {
    max := 999999
    min := 100000
    return fmt.Sprintf("%d", rand.Intn(max - min))
}

func currentYear() string {
    return fmt.Sprintf("%d", time.Now().Year())
}

func convertKeywordsSliceToString(keywords []string) string {
    return strings.Join(keywords, ",")
}

func createFuncMap() template.FuncMap {
   return template.FuncMap{
        "randomId": randomId,
        "currentYear": currentYear,
        "convertKeywordSliceToString": convertKeywordsSliceToString,
    }
}

type MainLayoutParams struct {
    Title string
    Description string
    PageUrl string
    Content template.HTML
    Keywords []string
}


var tmpl * template.Template

func loadTemplate(patterns []string) {

	funcMap := createFuncMap()

	tmpl = template.New("").Funcs(funcMap)

	for _, pattern := range patterns {

		matches, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatalf("Failed to glob pattern %s: %v", pattern, err)
		}

		if matches == nil {
			log.Fatalf("No templates found matching pattern: %s", pattern)
		}

		for _, match := range matches {

			_, err := tmpl.ParseFiles(match)
			if err != nil {
				log.Fatalf("Failed to parse template %s: %v", match, err)
			}
		}
	}
}

func main() {
    projectRoot := filepath.Join(".")
    templateDir := filepath.Join(projectRoot, "src", "templates")
    viewDir := filepath.Join(templateDir, "views", "*.html")
    layoutDir := filepath.Join(templateDir, "layouts", "*.html")
    componentsDir := filepath.Join(templateDir, "components", "*.html")

    paths := []string{viewDir, layoutDir, componentsDir}
    loadTemplate(paths)

    var homePageContent bytes.Buffer
    err := tmpl.ExecuteTemplate(&homePageContent, "home-page.html", nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("log.Logger: %v\n", homePageContent)

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

    homePageParams := MainLayoutParams{
        Title: "Website That feel Native | Christian Williams",
        Description: "I create Website that provide a rich experience to your customer that will convert into sales for you buisness",
        PageUrl: "https://mrwill.ca",
        Content: template.HTML(homePageContent.String()),
        Keywords: []string{"websites", "user experience", "conversions"},
    }

    // Check template names
	for _, t := range tmpl.Templates() {
		fmt.Println("Template name:", t.Name())
	}

    err = tmpl.ExecuteTemplate(homePageFile, "main-layout", homePageParams)
    if err != nil {
        log.Fatalf("Could not render template to dist/index.html: %v", err)
    }
}
