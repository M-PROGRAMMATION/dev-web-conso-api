package routes

import (
	"dev-web-conso-apii/src/utils"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func loadConcertsFixtureTemplate(t *testing.T) *utils.Templates {
	t.Helper()
	path := filepath.Join(t.TempDir(), "concerts.html")
	content := `{{define "concerts.html"}}{{range .Events}}<p>{{.ArtistName}} {{.City}}</p>{{end}}{{end}}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template fixture: %v", err)
	}

	templates, err := utils.LoadTemplates(path)
	if err != nil {
		t.Fatalf("LoadTemplates() returned an error: %v", err)
	}
	return templates
}

func TestConcertsHandlerRendersAllEvents(t *testing.T) {
	seedEventFixtures(t)

	handler := concertsHandler(loadConcertsFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/concerts", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
	if strings.Count(recorder.Body.String(), "<p>") != 3 {
		t.Fatalf("expected 3 events, got %q", recorder.Body.String())
	}
}

func TestConcertsHandlerFiltersByCity(t *testing.T) {
	seedEventFixtures(t)

	handler := concertsHandler(loadConcertsFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/concerts?city=paris", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	body := recorder.Body.String()
	if strings.Count(body, "<p>") != 2 {
		t.Fatalf("expected 2 Paris events, got %q", body)
	}
}

func TestConcertsHandlerInvalidDateReturns400(t *testing.T) {
	seedEventFixtures(t)

	handler := concertsHandler(loadConcertsFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/concerts?from=not-a-date", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 400 {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}
}
