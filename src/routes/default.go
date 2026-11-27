package routes

import (
	"dev-web-conso-apii/src/utils"
	"log"
	"net/http"
)

// DefaultRoute godoc
//
//	@Summary	Page d'accueil
//	@Description	Affiche la liste des artistes avec une barre de recherche
//	@Tags		default
//	@Produce	html
//	@Success	200	{string}	string	"page HTML"
//	@Router		/ [get]
func DefaultRoute() {
	templates, err := utils.LoadTemplates("resources/index.html")
	if err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}

	http.HandleFunc("/", indexHandler(templates))
}

func indexHandler(templates *utils.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		artists, err := utils.ArtistStore.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := templates.Render(w, "index.html", artists); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
