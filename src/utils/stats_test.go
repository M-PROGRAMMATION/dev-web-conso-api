package utils

import (
	"dev-web-conso-apii/src/objects"
	"testing"
)

func TestBuildStats(t *testing.T) {
	withTempStores(t)

	if err := ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := ArtistStore.Create(objects.Artist{ID: 2, Name: "Muse"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := RelationStore.Create(objects.Relation{
		ID: 1,
		DatesLocations: map[string][]string{
			"paris-france": {"14-03-2025", "15-03-2025"},
			"lyon-france":  {"20-03-2025"},
		},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	stats, err := BuildStats()
	if err != nil {
		t.Fatalf("BuildStats() returned an error: %v", err)
	}

	if stats.TotalArtists != 2 {
		t.Fatalf("expected 2 artists, got %d", stats.TotalArtists)
	}
	if stats.TotalEvents != 3 {
		t.Fatalf("expected 3 events, got %d", stats.TotalEvents)
	}
	if stats.EventsByCity["Paris"] != 2 {
		t.Fatalf("expected 2 events in Paris, got %d", stats.EventsByCity["Paris"])
	}
	if stats.EventsByCity["Lyon"] != 1 {
		t.Fatalf("expected 1 event in Lyon, got %d", stats.EventsByCity["Lyon"])
	}
}
