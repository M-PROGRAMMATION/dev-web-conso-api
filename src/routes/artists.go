package routes

import (
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"errors"
	"net/http"
	"strconv"
)

func GetAllArtists() {
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		artists, err := utils.ArtistStore.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, artists)
	})
}

func GetArtistById() {
	http.HandleFunc("GET /artists/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "id invalide", http.StatusBadRequest)
			return
		}

		artist, err := utils.ArtistStore.Find(id)
		if errors.Is(err, orm.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.WriteJSON(w, artist)
	})
}
