package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/utils"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func seedEventFixtures(t *testing.T) {
	t.Helper()
	withTempArtistStore(t)
	withTempRelationStore(t)

	if err := utils.ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := utils.ArtistStore.Create(objects.Artist{ID: 2, Name: "Muse"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}
	if err := utils.RelationStore.Create(objects.Relation{
		ID: 1,
		DatesLocations: map[string][]string{
			"paris-france": {"20-05-2025"},
			"lyon-france":  {"14-03-2025"},
		},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}
	if err := utils.RelationStore.Create(objects.Relation{
		ID:             2,
		DatesLocations: map[string][]string{"paris-france": {"01-01-2026"}},
	}); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}
}

func TestEventsHandlerReturnsSortedEvents(t *testing.T) {
	seedEventFixtures(t)

	req := httptest.NewRequest("GET", "/api/events", nil)
	recorder := httptest.NewRecorder()

	eventsHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var events []objects.Event
	if err := json.Unmarshal(recorder.Body.Bytes(), &events); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("expected 3 events, got %d", len(events))
	}
	for i := 1; i < len(events); i++ {
		if events[i].Date.Before(events[i-1].Date) {
			t.Fatalf("expected events sorted chronologically, got %+v", events)
		}
	}
}

func TestEventsHandlerFiltersByCity(t *testing.T) {
	seedEventFixtures(t)

	req := httptest.NewRequest("GET", "/api/events?city=lyon", nil)
	recorder := httptest.NewRecorder()

	eventsHandler(recorder, req)

	var events []objects.Event
	if err := json.Unmarshal(recorder.Body.Bytes(), &events); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if len(events) != 1 || events[0].City != "Lyon" {
		t.Fatalf("expected only Lyon, got %+v", events)
	}
}

func TestEventsHandlerFiltersByDateRange(t *testing.T) {
	seedEventFixtures(t)

	req := httptest.NewRequest("GET", "/api/events?from=2025-01-01&to=2025-12-31", nil)
	recorder := httptest.NewRecorder()

	eventsHandler(recorder, req)

	var events []objects.Event
	if err := json.Unmarshal(recorder.Body.Bytes(), &events); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("expected 2 events in 2025, got %+v", events)
	}
}

func TestEventsHandlerInvalidDateReturns400(t *testing.T) {
	seedEventFixtures(t)

	req := httptest.NewRequest("GET", "/api/events?from=not-a-date", nil)
	recorder := httptest.NewRecorder()

	eventsHandler(recorder, req)

	if recorder.Code != 400 {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}
}
