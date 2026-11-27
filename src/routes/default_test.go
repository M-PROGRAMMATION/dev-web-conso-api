package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func loadFixtureTemplates(t *testing.T) *utils.Templates {
	t.Helper()
	path := filepath.Join(t.TempDir(), "index.html")
	content := `{{define "index.html"}}{{range .}}<p>{{.Name}}</p>{{end}}{{end}}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template fixture: %v", err)
	}

	templates, err := utils.LoadTemplates(path)
	if err != nil {
		t.Fatalf("LoadTemplates() returned an error: %v", err)
	}
	return templates
}

func TestIndexHandlerRendersArtists(t *testing.T) {
	withTempArtistStore(t)
	if err := utils.ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}

	handler := indexHandler(loadFixtureTemplates(t))

	req := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
	if !strings.Contains(recorder.Body.String(), "Queen") {
		t.Fatalf("expected body to contain Queen, got %q", recorder.Body.String())
	}
}

func TestIndexHandlerUnknownPathReturns404(t *testing.T) {
	withTempArtistStore(t)
	handler := indexHandler(loadFixtureTemplates(t))

	req := httptest.NewRequest("GET", "/unknown", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 404 {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}
}

func TestIndexHandlerStoreErrorReturns500(t *testing.T) {
	original := utils.ArtistStore
	t.Cleanup(func() { utils.ArtistStore = original })

	// Point the store at a path nested under a file instead of a directory
	// so that reading it fails with something other than os.ErrNotExist.
	badPath := filepath.Join(t.TempDir(), "not-a-dir")
	if err := os.WriteFile(badPath, []byte("x"), 0644); err != nil {
		t.Fatalf("failed to create fixture file: %v", err)
	}
	utils.ArtistStore = orm.New[objects.Artist](filepath.Join(badPath, "artists.json"))

	handler := indexHandler(loadFixtureTemplates(t))

	req := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 500 {
		t.Fatalf("expected status 500, got %d", recorder.Code)
	}
}
