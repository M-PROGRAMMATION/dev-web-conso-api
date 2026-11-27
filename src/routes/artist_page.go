package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type artistPageData struct {
	Artist objects.Artist
	Events []objects.Event
}

// ArtistPageRoute godoc
//
//	@Summary	Page détaillée d'un artiste
//	@Description	Affiche les informations d'un artiste et la liste de ses concerts
//	@Tags		default
//	@Produce	html
//	@Param		id	path		int	true	"ID de l'artiste"
//	@Success	200	{string}	string	"page HTML"
//	@Failure	404	{string}	string	"artiste introuvable"
//	@Router		/artist/{id} [get]
func ArtistPageRoute() {
	templates, err := utils.LoadTemplates("resources/artist.html")
	if err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}

	http.HandleFunc("GET /artist/{id}", artistPageHandler(templates))
}

func artistPageHandler(templates *utils.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "id invalide", http.StatusBadRequest)
			return
		}

		artist, err := utils.ArtistStore.Find(id)
		if errors.Is(err, orm.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		events, err := utils.BuildEvents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		events = utils.FilterEventsByArtist(events, id)
		utils.SortEventsByDate(events)

		data := artistPageData{Artist: artist, Events: events}
		if err := templates.Render(w, "artist.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
