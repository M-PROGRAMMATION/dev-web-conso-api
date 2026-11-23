package routes

import (
	"fmt"
	"net/http"
	"time"
)

func HealthRoute(startTime time.Time) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
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
	})
}
