package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

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

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
