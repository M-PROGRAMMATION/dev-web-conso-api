package utils

import (
	"dev-web-conso-apii/src/objects"
	"sort"
	"strings"
	"time"
)

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

func SortEventsByDate(events []objects.Event) {
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date.Before(events[j].Date)
	})
}

func FilterEventsByArtist(events []objects.Event, artistID int) []objects.Event {
	filtered := make([]objects.Event, 0)

	for _, event := range events {
		if event.ArtistID == artistID {
			filtered = append(filtered, event)
		}
	}

	return filtered
}

func FilterEventsByCity(events []objects.Event, city string) []objects.Event {
	if city == "" {
		return events
	}

	filtered := make([]objects.Event, 0)

	for _, event := range events {
		if strings.EqualFold(event.City, city) {
			filtered = append(filtered, event)
		}
	}

	return filtered
}

func FilterEventsByDateRange(events []objects.Event, from, to time.Time) []objects.Event {
	if from.IsZero() && to.IsZero() {
		return events
	}

	filtered := make([]objects.Event, 0)

	for _, event := range events {
		if !from.IsZero() && event.Date.Before(from) {
			continue
		}
		if !to.IsZero() && event.Date.After(to) {
			continue
		}
		filtered = append(filtered, event)
	}

	return filtered
}
