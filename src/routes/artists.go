package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/utils"
	"net/http"
)

func GetAllArtists() {
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, objects.AllArtists)
	})
}

func GetAllRelation() {
	http.HandleFunc("GET /relation", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, objects.AllRelations)
	})
}
