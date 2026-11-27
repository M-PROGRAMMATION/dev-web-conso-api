package routes

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHealthHandlerReturnsOk(t *testing.T) {
	handler := healthHandler(time.Now())

	req := httptest.NewRequest("GET", "/health", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
	if got := recorder.Header().Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Fatalf("expected text/plain content type, got %q", got)
	}
	if !strings.Contains(recorder.Body.String(), "Health: OK") {
		t.Fatalf("expected body to contain Health: OK, got %q", recorder.Body.String())
	}
}

func TestHealthHandlerUnknownPathReturns404(t *testing.T) {
	handler := healthHandler(time.Now())

	req := httptest.NewRequest("GET", "/unknown", nil)
	recorder := httptest.NewRecorder()

	handler(recorder, req)

	if recorder.Code != 404 {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}
}
