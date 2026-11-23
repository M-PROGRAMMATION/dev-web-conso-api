package main

import (
	"dev-web-conso-apii/src/banner"
	"dev-web-conso-apii/src/routes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"
	var startTime = time.Now()

	routes.DefaultRoute()
	routes.HealthRoute(startTime)

	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path != "/artists" {
			http.NotFound(w, r)
			return
		}

		response := struct {
			ID           int      `json:"id"`
			Name         string   `json:"name"`
			CreationDate int      `json:"creationDate"`
			Members      []string `json:"members"`
			FirstAlbum   string   `json:"firstAlbum"`
		}{
			ID:           1,
			Name:         "Aurore Nomade",
			CreationDate: 2017,
			Members:      []string{"Lina", "Nora", "Sami"},
			FirstAlbum:   "12-05-2019",
		}

		json.NewEncoder(w).Encode(response)
	})

	banner.PrintBanner(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
