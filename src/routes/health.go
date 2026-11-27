package routes

import (
	"fmt"
	"net/http"
	"time"
)

// HealthRoute godoc
//
//	@Summary	Statut du serveur
//	@Description	Retourne l'heure de démarrage et l'uptime du serveur
//	@Tags		default
//	@Produce	plain
//	@Success	200	{string}	string	"Health: OK"
//	@Router		/health [get]
func HealthRoute(startTime time.Time) {
	http.HandleFunc("/health", healthHandler(startTime))
}

func healthHandler(startTime time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			http.NotFound(w, r)
			return
		}

		uptime := time.Since(startTime).Round(time.Second)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Health: OK\nDepuis : %s\nUptime : %s\n",
			startTime.Format("02/01/2006 15:04:05"),
			uptime,
		)
	}
}
