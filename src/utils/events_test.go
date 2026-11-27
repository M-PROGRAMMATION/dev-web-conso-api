package utils

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"path/filepath"
	"testing"
	"time"
)

func withTempStores(t *testing.T) {
	t.Helper()
	originalArtists := ArtistStore
	originalRelations := RelationStore

	ArtistStore = orm.New[objects.Artist](filepath.Join(t.TempDir(), "artists.json"))
	RelationStore = orm.New[objects.Relation](filepath.Join(t.TempDir(), "relations.json"))

	t.Cleanup(func() {
		ArtistStore = originalArtists
		RelationStore = originalRelations
	})
}

func TestBuildEventsJoinsArtistsAndRelations(t *testing.T) {
	withTempStores(t)

	if err := ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := RelationStore.Create(objects.Relation{
		ID: 1,
		DatesLocations: map[string][]string{
			"paris-france": {"14-03-2025", "15-03-2025"},
		},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	events, err := BuildEvents()
	if err != nil {
		t.Fatalf("BuildEvents() returned an error: %v", err)
	}

	if len(events) != 2 {
		t.Fatalf("expected 2 events, got %d", len(events))
	}

	for _, event := range events {
		if event.ArtistName != "Queen" {
			t.Fatalf("expected artist name Queen, got %q", event.ArtistName)
		}
		if event.City != "Paris" || event.Country != "France" {
			t.Fatalf("expected Paris/France, got %s/%s", event.City, event.Country)
		}
	}
}

func TestBuildEventsSkipsRelationsWithoutMatchingArtist(t *testing.T) {
	withTempStores(t)

	if err := RelationStore.Create(objects.Relation{
		ID:             99,
		DatesLocations: map[string][]string{"paris-france": {"14-03-2025"}},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	events, err := BuildEvents()
	if err != nil {
		t.Fatalf("BuildEvents() returned an error: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("expected no events, got %d", len(events))
	}
}

func TestBuildEventsReturnsErrorOnInvalidDate(t *testing.T) {
	withTempStores(t)

	if err := ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := RelationStore.Create(objects.Relation{
		ID:             1,
		DatesLocations: map[string][]string{"paris-france": {"not-a-date"}},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	if _, err := BuildEvents(); err == nil {
		t.Fatal("expected an error for an invalid date")
	}
}

func TestBuildEventsDateIsNormalized(t *testing.T) {
	withTempStores(t)

	if err := ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := RelationStore.Create(objects.Relation{
		ID:             1,
		DatesLocations: map[string][]string{"paris-france": {"14-03-2025"}},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	events, err := BuildEvents()
	if err != nil {
		t.Fatalf("BuildEvents() returned an error: %v", err)
	}

	want := time.Date(2025, time.March, 14, 0, 0, 0, 0, time.UTC)
	if !events[0].Date.Equal(want) {
		t.Fatalf("expected date %v, got %v", want, events[0].Date)
	}
}
