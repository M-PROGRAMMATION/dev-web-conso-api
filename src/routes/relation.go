package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/utils"
	"fmt"
	"net/http"
)

func GetAllRelation() {
	http.HandleFunc("GET /relation", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSON(w, objects.AllRelations)
	})
}

func GetRelationById() {
	http.HandleFunc("GET /relation/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id int
		if _, err := fmt.Sscan(r.PathValue("id"), &id); err != nil {
			http.Error(w, "id invalide", http.StatusBadRequest)
			return
		}
		for _, rel := range objects.AllRelations.Index {
			if rel.ID == id {
				utils.WriteJSON(w, rel)
				return
			}
		}
		http.NotFound(w, r)
	})
}
