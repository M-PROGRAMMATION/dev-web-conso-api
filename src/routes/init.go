package routes

import (
	"dev-web-conso-apii/src/utils"
	"net/http"
	"time"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func InitRoutes(startTime time.Time) {
	utils.SeedStores()

	// Default Routes
	DefaultRoute()
	HealthRoute(startTime)
	ArtistPageRoute()
	ConcertsRoute()

	// Swagger UI
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// Artists Routes
	GetAllArtists()
	GetArtistById()

	// Relation Routes
	GetAllRelation()
	GetRelationById()

	// Events & Stats Routes
	GetEvents()
	GetStats()
}
