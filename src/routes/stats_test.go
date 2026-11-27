package routes

import (
	"dev-web-conso-apii/src/objects"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestStatsHandlerReturnsStats(t *testing.T) {
	seedEventFixtures(t)

	req := httptest.NewRequest("GET", "/api/stats", nil)
	recorder := httptest.NewRecorder()

	statsHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var stats objects.Stats
	if err := json.Unmarshal(recorder.Body.Bytes(), &stats); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if stats.TotalArtists != 2 {
		t.Fatalf("expected 2 artists, got %d", stats.TotalArtists)
	}
	if stats.TotalEvents != 3 {
		t.Fatalf("expected 3 events, got %d", stats.TotalEvents)
	}
}
