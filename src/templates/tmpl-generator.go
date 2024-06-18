package main

import (
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
    return string(rand.Intn(max - min))
}

func currentYear() string {
    return string(time.Now().Year())
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
    Content * template.Template
    Keywords []string
}

var tmpl * template.Template

func main() {
    funcMap := createFuncMap()

    tmpl = template.New("").Funcs(funcMap)

    const templateDir = "./src/templates/"
    viewDir := filepath.Join(templateDir, "views")
    
    homePageBodyContent, err := tmpl.Parse(viewDir + "home-page.html")
    if err != nil {
        log.Fatal(err)
    }

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
        Content: homePageBodyContent,
        Keywords: []string{"websites", "user experience", "conversions"},
    }

    err = tmpl.ExecuteTemplate(homePageFile, "home-page.html", homePageParams)
    if err != nil {
        log.Fatalf("Could not render template to dist/index.html: %v", err)
    }
}
