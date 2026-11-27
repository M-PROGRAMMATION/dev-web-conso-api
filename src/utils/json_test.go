package utils

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONSetsContentTypeAndBody(t *testing.T) {
	recorder := httptest.NewRecorder()

	WriteJSON(recorder, map[string]string{"hello": "world"})

	if got := recorder.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("expected application/json content type, got %q", got)
	}

	var body map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if body["hello"] != "world" {
		t.Fatalf("expected hello=world, got %+v", body)
	}
}
