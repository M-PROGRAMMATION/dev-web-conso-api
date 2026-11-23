package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/utils"
	"net/http"
	"strconv"
)

func GetAllArtists() {
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, objects.AllArtists)
	})
}

func GetArtistById() {
	http.HandleFunc("GET /artists/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "id invalide", http.StatusBadRequest)
			return
		}
		for _, a := range objects.AllArtists {
			if a.ID == id {
				utils.WriteJSON(w, a)
				return
			}
		}
		http.NotFound(w, r)
	})
}
