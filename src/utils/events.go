package utils

import "dev-web-conso-apii/src/objects"

// BuildEvents joint les artistes et leurs relations (dates/lieux) pour
// produire la liste des concerts, avec dates et lieux normalisés.
func BuildEvents() ([]objects.Event, error) {
	artists, err := ArtistStore.All()
	if err != nil {
		return nil, err
	}

	relations, err := RelationStore.All()
	if err != nil {
		return nil, err
	}

	artistsByID := make(map[int]objects.Artist, len(artists))
	for _, artist := range artists {
		artistsByID[artist.ID] = artist
	}

	events := make([]objects.Event, 0)

	for _, relation := range relations {
		artist, ok := artistsByID[relation.ID]
		if !ok {
			continue
		}

		for location, dates := range relation.DatesLocations {
			city, country := NormalizeLocation(location)

			for _, raw := range dates {
				date, err := NormalizeDate(raw)
				if err != nil {
					return nil, err
				}

				events = append(events, objects.Event{
					ArtistID:   artist.ID,
					ArtistName: artist.Name,
					City:       city,
					Country:    country,
					Date:       date,
				})
			}
		}
	}

	return events, nil
}
