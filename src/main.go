package main

import (
	"dev-web-conso-apii/src/routes"
	"dev-web-conso-apii/src/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8080"
	var startTime = time.Now()

	// Default Routes
	routes.DefaultRoute()
	routes.HealthRoute(startTime)

	// Artists Routes
	routes.GetAllArtists()
	routes.GetArtistById()

	// Relation Routes
	routes.GetAllRelation()
	routes.GetRelationById()

	utils.PrintBanner(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
