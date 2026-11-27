package routes

import (
	"dev-web-conso-apii/src/utils"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func loadArtistFixtureTemplate(t *testing.T) *utils.Templates {
	t.Helper()
	path := filepath.Join(t.TempDir(), "artist.html")
	content := `{{define "artist.html"}}<h1>{{.Artist.Name}}</h1>{{range .Events}}<p>{{.City}} {{.Date.Format "02/01/2006"}}</p>{{end}}{{end}}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template fixture: %v", err)
	}

	templates, err := utils.LoadTemplates(path)
	if err != nil {
		t.Fatalf("LoadTemplates() returned an error: %v", err)
	}
	return templates
}

func TestArtistPageHandlerRendersArtistAndEvents(t *testing.T) {
	seedEventFixtures(t)

	handler := artistPageHandler(loadArtistFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/artist/1", nil)
	req.SetPathValue("id", "1")
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "Queen") {
		t.Fatalf("expected body to contain Queen, got %q", body)
	}
	if !strings.Contains(body, "Paris") || !strings.Contains(body, "Lyon") {
		t.Fatalf("expected body to contain Queen's concerts, got %q", body)
	}
	if strings.Count(body, "<p>") != 2 {
		t.Fatalf("expected only Queen's 2 events, got %q", body)
	}
}

func TestArtistPageHandlerNotFound(t *testing.T) {
	seedEventFixtures(t)

	handler := artistPageHandler(loadArtistFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/artist/99", nil)
	req.SetPathValue("id", "99")
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 404 {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}
}

func TestArtistPageHandlerInvalidId(t *testing.T) {
	seedEventFixtures(t)

	handler := artistPageHandler(loadArtistFixtureTemplate(t))

	req := httptest.NewRequest("GET", "/artist/abc", nil)
	req.SetPathValue("id", "abc")
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 400 {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}
}
