package utils

import "dev-web-conso-apii/src/objects"

// BuildStats calcule des statistiques globales à partir des artistes et des
// concerts (nombre total, répartition par ville).
func BuildStats() (objects.Stats, error) {
	artists, err := ArtistStore.All()
	if err != nil {
		return objects.Stats{}, err
	}

	events, err := BuildEvents()
	if err != nil {
		return objects.Stats{}, err
	}

	eventsByCity := make(map[string]int)
	for _, event := range events {
		eventsByCity[event.City]++
	}

	return objects.Stats{
		TotalArtists: len(artists),
		TotalEvents:  len(events),
		EventsByCity: eventsByCity,
	}, nil
}
