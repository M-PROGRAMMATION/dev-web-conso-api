package main

import (
	"dev-web-conso-apii/src/banner"
	"dev-web-conso-apii/src/routes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

var artists = []Artist{
	{ID: 1, Name: "Aurore Nomade", CreationDate: 2017,
		Members: []string{"Lina", "Nora", "Sami"}, FirstAlbum: "12-05-2019"},
}

var relations = RelationIndex{
	Index: []Relation{
		{
			ID: 1, // même ID que l'artiste concerné
			DatesLocations: map[string][]string{
				"toulouse-france": {"14-03-2025", "15-03-2025"},
				"lyon-france":     {"02-04-2025"},
				"paris-france":    {"20-05-2025"},
			},
		},
	},
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
