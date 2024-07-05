package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	UserTitle    string
	FooterText   string
	PortfolioURL string
	Author       string
}

func main() {
	pageData := PageData{
		UserTitle:    "Muhammad Azeem",
		FooterText:   "Ihya by",
		PortfolioURL: "https://github.com/tangent-cloud",
		Author:       "TangentLabs",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parsedTemplate, err := template.ParseGlob("templates/*.html")
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
		_, _ = parsedTemplate.ParseGlob("templates/themes/*/*.html")
		err = parsedTemplate.ExecuteTemplate(
			w, "ihya-web", pageData,
		)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	})

	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
	log.Println("Starting server on :8090")
}
