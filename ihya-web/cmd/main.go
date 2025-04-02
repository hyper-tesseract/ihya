package main

import (
	"log"
	"net/http"

	handle "ihya-web/app/adapter/routing/handler"
	"ihya-web/app/adapter/sqldb"
	"ihya-web/app/adapter/templating"
	"ihya-web/app/usecase/ssr"
	"ihya-web/app/usecase/user"
	"ihya-web/app/usecase/userland"
)

func main() {
	fileServer := http.FileServer(http.Dir("../public"))

	personaRepo := sqldb.NewFakePersonaRepo()
	personaProvider := userland.NewPersonaRenderer(&personaRepo)

	themeRepo := sqldb.NewFakeThemeRepo()
	themeProvider := ssr.NewThemeProvider(&themeRepo)

	renderer := templating.NewTemplateRenderer()

	userRepo := sqldb.NewFakeUserRetrieverRepo()
	userRetriever := user.NewUserRetriever(&userRepo)

	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.HandleFunc("/", handle.ServePersonaHome(personaProvider, renderer, themeProvider, userRetriever))

	log.Fatal(http.ListenAndServe(":8090", nil))
	log.Println("Starting server on :8090")
}
