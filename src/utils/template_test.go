package utils

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadTemplatesMissingFileReturnsError(t *testing.T) {
	_, err := LoadTemplates(filepath.Join(t.TempDir(), "missing.html"))
	if err == nil {
		t.Fatal("expected an error when the template file does not exist")
	}
}

func TestRenderWritesHTMLBody(t *testing.T) {
	path := filepath.Join(t.TempDir(), "index.html")
	content := `{{define "index.html"}}<p>Hello {{.}}</p>{{end}}`
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write template fixture: %v", err)
	}

	templates, err := LoadTemplates(path)
	if err != nil {
		t.Fatalf("LoadTemplates() returned an error: %v", err)
	}

	recorder := httptest.NewRecorder()
	if err := templates.Render(recorder, "index.html", "World"); err != nil {
		t.Fatalf("Render() returned an error: %v", err)
	}

	if got := recorder.Header().Get("Content-Type"); got != "text/html; charset=utf-8" {
		t.Fatalf("expected text/html content type, got %q", got)
	}
	if !strings.Contains(recorder.Body.String(), "Hello World") {
		t.Fatalf("expected rendered body to contain %q, got %q", "Hello World", recorder.Body.String())
	}
}

func TestRenderUnknownTemplateReturnsError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "index.html")
	if err := os.WriteFile(path, []byte(`{{define "index.html"}}ok{{end}}`), 0644); err != nil {
		t.Fatalf("failed to write template fixture: %v", err)
	}

	templates, err := LoadTemplates(path)
	if err != nil {
		t.Fatalf("LoadTemplates() returned an error: %v", err)
	}

	recorder := httptest.NewRecorder()
	if err := templates.Render(recorder, "does-not-exist.html", nil); err == nil {
		t.Fatal("expected an error when rendering an unknown template")
	}
}
