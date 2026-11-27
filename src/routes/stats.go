package routes

import (
	"dev-web-conso-apii/src/utils"
	"net/http"
)

// GetStats godoc
//
//	@Summary	Statistiques globales (nombre d'artistes, de concerts, répartition par ville)
//	@Tags		stats
//	@Produce	json
//	@Success	200	{object}	objects.Stats
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/api/stats [get]
func GetStats() {
	http.HandleFunc("GET /api/stats", statsHandler)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := utils.BuildStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, stats)
}
