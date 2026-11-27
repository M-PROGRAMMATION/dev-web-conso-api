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

func withTempRelationStore(t *testing.T) {
	t.Helper()
	original := utils.RelationStore
	path := filepath.Join(t.TempDir(), "relations.json")
	utils.RelationStore = orm.New[objects.Relation](path)
	t.Cleanup(func() { utils.RelationStore = original })
}

func TestAllRelationHandlerReturnsIndex(t *testing.T) {
	withTempRelationStore(t)
	relation := objects.Relation{
		ID:             1,
		DatesLocations: map[string][]string{"paris-france": {"20-05-2025"}},
	}
	if err := utils.RelationStore.Create(relation); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	req := httptest.NewRequest("GET", "/relation", nil)
	recorder := httptest.NewRecorder()

	allRelationHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}

	var index objects.RelationIndex
	if err := json.Unmarshal(recorder.Body.Bytes(), &index); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if len(index.Index) != 1 || index.Index[0].ID != 1 {
		t.Fatalf("expected one relation with id 1, got %+v", index.Index)
	}
}

func TestRelationByIdHandlerFound(t *testing.T) {
	withTempRelationStore(t)
	relation := objects.Relation{
		ID:             1,
		DatesLocations: map[string][]string{"paris-france": {"20-05-2025"}},
	}
	if err := utils.RelationStore.Create(relation); err != nil {
		t.Fatalf("failed to seed relation: %v", err)
	}

	req := httptest.NewRequest("GET", "/relation/1", nil)
	req.SetPathValue("id", "1")
	recorder := httptest.NewRecorder()

	relationByIdHandler(recorder, req)

	if recorder.Code != 200 {
		t.Fatalf("expected status 200, got %d", recorder.Code)
	}
}

func TestRelationByIdHandlerNotFound(t *testing.T) {
	withTempRelationStore(t)

	req := httptest.NewRequest("GET", "/relation/99", nil)
	req.SetPathValue("id", "99")
	recorder := httptest.NewRecorder()

	relationByIdHandler(recorder, req)

	if recorder.Code != 404 {
		t.Fatalf("expected status 404, got %d", recorder.Code)
	}
}

func TestRelationByIdHandlerInvalidId(t *testing.T) {
	withTempRelationStore(t)

	req := httptest.NewRequest("GET", "/relation/abc", nil)
	req.SetPathValue("id", "abc")
	recorder := httptest.NewRecorder()

	relationByIdHandler(recorder, req)

	if recorder.Code != 400 {
		t.Fatalf("expected status 400, got %d", recorder.Code)
	}
}
