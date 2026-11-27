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

func mustDate(t *testing.T, raw string) time.Time {
	t.Helper()
	date, err := time.Parse("02-01-2006", raw)
	if err != nil {
		t.Fatalf("invalid fixture date %q: %v", raw, err)
	}
	return date
}

func TestSortEventsByDate(t *testing.T) {
	events := []objects.Event{
		{ArtistName: "B", Date: mustDate(t, "20-05-2025")},
		{ArtistName: "A", Date: mustDate(t, "14-03-2025")},
		{ArtistName: "C", Date: mustDate(t, "01-01-2026")},
	}

	SortEventsByDate(events)

	if events[0].ArtistName != "A" || events[1].ArtistName != "B" || events[2].ArtistName != "C" {
		t.Fatalf("expected events sorted A, B, C, got %v", events)
	}
}

func TestFilterEventsByArtist(t *testing.T) {
	events := []objects.Event{
		{ArtistID: 1, ArtistName: "Queen"},
		{ArtistID: 2, ArtistName: "Muse"},
	}

	filtered := FilterEventsByArtist(events, 2)

	if len(filtered) != 1 || filtered[0].ArtistName != "Muse" {
		t.Fatalf("expected only Muse, got %+v", filtered)
	}
}

func TestFilterEventsByCityIsCaseInsensitive(t *testing.T) {
	events := []objects.Event{
		{City: "Paris"},
		{City: "Lyon"},
	}

	filtered := FilterEventsByCity(events, "paris")

	if len(filtered) != 1 || filtered[0].City != "Paris" {
		t.Fatalf("expected only Paris, got %+v", filtered)
	}
}

func TestFilterEventsByCityEmptyReturnsAll(t *testing.T) {
	events := []objects.Event{{City: "Paris"}, {City: "Lyon"}}

	filtered := FilterEventsByCity(events, "")

	if len(filtered) != 2 {
		t.Fatalf("expected all events, got %+v", filtered)
	}
}

func TestFilterEventsByDateRange(t *testing.T) {
	events := []objects.Event{
		{ArtistName: "too early", Date: mustDate(t, "01-01-2025")},
		{ArtistName: "in range", Date: mustDate(t, "15-03-2025")},
		{ArtistName: "too late", Date: mustDate(t, "01-01-2026")},
	}

	from := mustDate(t, "01-03-2025")
	to := mustDate(t, "31-03-2025")

	filtered := FilterEventsByDateRange(events, from, to)

	if len(filtered) != 1 || filtered[0].ArtistName != "in range" {
		t.Fatalf("expected only the in-range event, got %+v", filtered)
	}
}

func TestFilterEventsByDateRangeZeroBoundsReturnsAll(t *testing.T) {
	events := []objects.Event{
		{Date: mustDate(t, "01-01-2025")},
		{Date: mustDate(t, "01-01-2026")},
	}

	filtered := FilterEventsByDateRange(events, time.Time{}, time.Time{})

	if len(filtered) != 2 {
		t.Fatalf("expected all events, got %+v", filtered)
	}
}
