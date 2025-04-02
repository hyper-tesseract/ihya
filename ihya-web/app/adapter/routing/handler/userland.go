package handler

import (
	"fmt"
	"ihya-web/app/adapter/templating"
	"ihya-web/app/hyper"
	"ihya-web/app/usecase/ssr"
	"ihya-web/app/usecase/user"
	"ihya-web/app/usecase/userland"
	"log"
	"net/http"
)

func ServePersonaHome(
	personaProvider userland.PersonaProvider,
	renderer *templating.TemplateRenderer,
	themerProvider ssr.ThemeProvider,
	userRetriever user.UserRetriever,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			RenderIndex(w, renderer)
			return
		}

		invalidPage := fmt.Sprintf("Invalid page path. %s", path)

		pagePath := hyper.SplitString(path, "/")
		if len(pagePath) != 2 {
			log.Println(invalidPage)
			RenderIndex(w, renderer)
			return
		}

		personaName := pagePath[1]
		if len(personaName) == 0 || len(personaName) > 50 {
			log.Println(invalidPage)
			RenderIndex(w, renderer)
			return
		}

		persona, err := personaProvider.RetrievePersona(personaName)
		if err != nil {
			log.Printf("Persona doesn't exist. %s", personaName)
			RenderIndex(w, renderer)
			return
		}

		theme, err := themerProvider.RetrieveTheme(persona.ThemeId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		user, err := userRetriever.RetrieveUser(persona.UserId)
		if err != nil {
			log.Printf("Persona doesn't exist. %s", personaName)
			RenderIndex(w, renderer)
			return
		}

		pageData := ssr.PreparePersonaTemplateMapping(persona, theme, user)
		err = renderer.RenderPage(w, pageData)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
