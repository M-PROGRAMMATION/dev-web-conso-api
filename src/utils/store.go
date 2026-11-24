package utils

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"log"
)

var ArtistStore = orm.New[objects.Artist]("data/artists.json")
var RelationStore = orm.New[objects.Relation]("data/relations.json")

func SeedStores() {
	seedArtists, err := ArtistStore.All()
	if err != nil {
		log.Fatalf("failed to load artists store: %v", err)
	}
	if len(seedArtists) == 0 {
		for _, artist := range objects.SeedArtists {
			if err := ArtistStore.Create(artist); err != nil {
				log.Fatalf("failed to seed artists store: %v", err)
			}
		}
	}

	seedRelations, err := RelationStore.All()
	if err != nil {
		log.Fatalf("failed to load relations store: %v", err)
	}
	if len(seedRelations) == 0 {
		for _, relation := range objects.SeedRelations {
			if err := RelationStore.Create(relation); err != nil {
				log.Fatalf("failed to seed relations store: %v", err)
			}
		}
	}
}
