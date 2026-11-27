package routes

import (
	"dev-web-conso-apii/src/objects"
	"dev-web-conso-apii/src/orm"
	"dev-web-conso-apii/src/utils"
	"encoding/json"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func withTempArtistStore(t *testing.T) {
	t.Helper()
	original := utils.ArtistStore
	path := filepath.Join(t.TempDir(), "artists.json")
	utils.ArtistStore = orm.New[objects.Artist](path)
	t.Cleanup(func() { utils.ArtistStore = original })
}

func TestAllArtistsHandlerReturnsArtists(t *testing.T) {
	withTempArtistStore(t)
	if err := utils.ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}

	req := httptest.NewRequest("GET", "/artists", nil)
	recorder := httptest.NewRecorder()

	allArtistsHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var artists []objects.Artist
	if err := json.Unmarshal(recorder.Body.Bytes(), &artists); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if len(artists) != 1 || artists[0].Name != "Queen" {
		t.Fatalf("expected one artist named Queen, got %+v", artists)
	}
}

func TestArtistByIdHandlerFound(t *testing.T) {
	withTempArtistStore(t)
	if err := utils.ArtistStore.Create(objects.Artist{ID: 1, Name: "Queen"}); err != nil {
		t.Fatalf("failed to seed artist: %v", err)
	}

	req := httptest.NewRequest("GET", "/artists/1", nil)
	req.SetPathValue("id", "1")
	recorder := httptest.NewRecorder()

	artistByIdHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var artist objects.Artist
	if err := json.Unmarshal(recorder.Body.Bytes(), &artist); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if artist.Name != "Queen" {
		t.Fatalf("expected Queen, got %q", artist.Name)
	}
}

func TestArtistByIdHandlerNotFound(t *testing.T) {
	withTempArtistStore(t)

	req := httptest.NewRequest("GET", "/artists/99", nil)
	req.SetPathValue("id", "99")
	recorder := httptest.NewRecorder()

	artistByIdHandler(recorder, req)

	if recorder.Code != 404 {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}
}

func TestArtistByIdHandlerInvalidId(t *testing.T) {
	withTempArtistStore(t)

	req := httptest.NewRequest("GET", "/artists/abc", nil)
	req.SetPathValue("id", "abc")
	recorder := httptest.NewRecorder()

	artistByIdHandler(recorder, req)

	if recorder.Code != 400 {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}
}
