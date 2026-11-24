package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"errors"
	"fmt"
	"net/http"
)

func GetAllRelation() {
	http.HandleFunc("GET /relation", func(w http.ResponseWriter, r *http.Request) {
		relations, err := utils.RelationStore.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, objects.RelationIndex{Index: relations})
	})
}

func GetRelationById() {
	http.HandleFunc("GET /relation/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id int
		if _, err := fmt.Sscan(r.PathValue("id"), &id); err != nil {
			http.Error(w, "id invalide", http.StatusBadRequest)
			return
		}

		relation, err := utils.RelationStore.Find(id)
		if errors.Is(err, orm.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.WriteJSON(w, relation)
	})
}
