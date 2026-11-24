package routes

import (
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"errors"
	"net/http"
	"strconv"
)

// GetAllArtists godoc
//
//	@Summary	Liste des artistes
//	@Tags		artists
//	@Produce	json
//	@Success	200	{array}		objects.Artist
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/artists [get]
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

// GetArtistById godoc
//
//	@Summary	Récupère un artiste par son id
//	@Tags		artists
//	@Produce	json
//	@Param		id	path		int	true	"ID de l'artiste"
//	@Success	200	{object}	objects.Artist
//	@Failure	400	{string}	string	"id invalide"
//	@Failure	404	{string}	string	"artiste introuvable"
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/artists/{id} [get]
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
