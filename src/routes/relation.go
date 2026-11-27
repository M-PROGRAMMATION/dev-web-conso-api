package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"errors"
	"fmt"
	"net/http"
)

// GetAllRelation godoc
//
//	@Summary	Liste de toutes les relations (dates/lieux)
//	@Tags		relation
//	@Produce	json
//	@Success	200	{object}	objects.RelationIndex
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/relation [get]
func GetAllRelation() {
	http.HandleFunc("GET /relation", allRelationHandler)
}

func allRelationHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := utils.RelationStore.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, objects.RelationIndex{Index: relations})
}

// GetRelationById godoc
//
//	@Summary	Récupère une relation par son id
//	@Tags		relation
//	@Produce	json
//	@Param		id	path		int	true	"ID de la relation"
//	@Success	200	{object}	objects.Relation
//	@Failure	400	{string}	string	"id invalide"
//	@Failure	404	{string}	string	"relation introuvable"
//	@Failure	500	{string}	string	"erreur serveur"
//	@Router		/relation/{id} [get]
func GetRelationById() {
	http.HandleFunc("GET /relation/{id}", relationByIdHandler)
}

func relationByIdHandler(w http.ResponseWriter, r *http.Request) {
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
}
