package main

import (
	"dev-web-conso-apii/src/banner"
	"dev-web-conso-apii/src/routes"
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
	routes.GetAllRelation()

	banner.PrintBanner(addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
