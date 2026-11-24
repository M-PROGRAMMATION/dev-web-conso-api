package routes

import (
	"dev-web-conso-apii/src/utils"
	"time"
)

func InitRoutes(startTime time.Time) {
	utils.SeedStores()

	// Default Routes
	DefaultRoute()
	HealthRoute(startTime)

	// Artists Routes
	GetAllArtists()
	GetArtistById()

	// Relation Routes
	GetAllRelation()
	GetRelationById()
}
