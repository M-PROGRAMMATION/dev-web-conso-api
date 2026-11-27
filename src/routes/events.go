package routes

import (
	"dev-web-conso-apii/src/utils"
	"fmt"
	"net/http"
	"time"
)

const eventsQueryDateLayout = "2006-01-02"

// GetEvents godoc
//
//	@Summary	Liste des concerts, triés par date et filtrables par ville / intervalle de dates
//	@Tags		events
//	@Produce	json
//	@Param		city	query		string	false	"Filtre par ville"
//	@Param		from	query		string	false	"Date de début (AAAA-MM-JJ)"
//	@Param		to		query		string	false	"Date de fin (AAAA-MM-JJ)"
//	@Success	200	{array}		objects.Event
//	@Failure	400	{string}	string	"paramètre invalide"
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/api/events [get]
func GetEvents() {
	http.HandleFunc("GET /api/events", eventsHandler)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := utils.BuildEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.URL.Query()

	from, to, err := parseEventDateRange(query.Get("from"), query.Get("to"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events = utils.FilterEventsByCity(events, query.Get("city"))
	events = utils.FilterEventsByDateRange(events, from, to)
	utils.SortEventsByDate(events)

	utils.WriteJSON(w, events)
}

func parseEventDateRange(fromRaw, toRaw string) (from, to time.Time, err error) {
	if fromRaw != "" {
		if from, err = time.Parse(eventsQueryDateLayout, fromRaw); err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("paramètre from invalide (attendu AAAA-MM-JJ)")
		}
	}

	if toRaw != "" {
		if to, err = time.Parse(eventsQueryDateLayout, toRaw); err != nil {
			return time.Time{}, time.Time{}, fmt.Errorf("paramètre to invalide (attendu AAAA-MM-JJ)")
		}
	}

	return from, to, nil
}
