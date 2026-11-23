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
