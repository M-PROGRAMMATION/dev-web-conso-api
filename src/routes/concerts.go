package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/utils"
	"log"
	"net/http"
)

type concertsPageData struct {
	Events []objects.Event
	City   string
	From   string
	To     string
}

// ConcertsRoute godoc
//
//	@Summary	Liste des concerts
//	@Description	Affiche tous les concerts, triés par date, filtrables par ville et intervalle de dates
//	@Tags		default
//	@Produce	html
//	@Param		city	query		string	false	"Filtre par ville"
//	@Param		from	query		string	false	"Date de début (AAAA-MM-JJ)"
//	@Param		to		query		string	false	"Date de fin (AAAA-MM-JJ)"
//	@Success	200	{string}	string	"page HTML"
//	@Router		/concerts [get]
func ConcertsRoute() {
	templates, err := utils.LoadTemplates("resources/concerts.html")
	if err != nil {
		log.Fatalf("failed to load templates: %v", err)
	}

	http.HandleFunc("GET /concerts", concertsHandler(templates))
}

func concertsHandler(templates *utils.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := utils.BuildEvents()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		query := r.URL.Query()
		city := query.Get("city")
		fromRaw := query.Get("from")
		toRaw := query.Get("to")

		from, to, err := parseEventDateRange(fromRaw, toRaw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		events = utils.FilterEventsByCity(events, city)
		events = utils.FilterEventsByDateRange(events, from, to)
		utils.SortEventsByDate(events)

		data := concertsPageData{Events: events, City: city, From: fromRaw, To: toRaw}
		if err := templates.Render(w, "concerts.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
