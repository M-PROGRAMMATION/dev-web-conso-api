package routes

import "time"

func InitRoutes(startTime time.Time) {
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
