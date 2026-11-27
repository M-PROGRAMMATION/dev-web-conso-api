package orm

import (
	"errors"
	"path/filepath"
	"testing"
)

type testRecord struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r testRecord) GetID() int {
	return r.ID
}

func newTestStore(t *testing.T) *Store[testRecord] {
	t.Helper()
	path := filepath.Join(t.TempDir(), "records.json")
	return New[testRecord](path)
}

func TestAllOnMissingFileReturnsEmptySlice(t *testing.T) {
	store := newTestStore(t)

	records, err := store.All()
	if err != nil {
		t.Fatalf("All() returned an error: %v", err)
	}
	if len(records) != 0 {
		t.Fatalf("expected no records, got %d", len(records))
	}
}

func TestCreateAndFind(t *testing.T) {
	store := newTestStore(t)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	found, err := store.Find(1)
	if err != nil {
		t.Fatalf("Find() returned an error: %v", err)
	}
	if found.Name != "Alice" {
		t.Fatalf("expected Alice, got %q", found.Name)
	}
}

func TestCreateDuplicateFails(t *testing.T) {
	store := newTestStore(t)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	err := store.Create(testRecord{ID: 1, Name: "Bob"})
	if !errors.Is(err, ErrDuplicate) {
		t.Fatalf("expected ErrDuplicate, got %v", err)
	}
}

func TestFindMissingReturnsErrNotFound(t *testing.T) {
	store := newTestStore(t)

	_, err := store.Find(42)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateExistingRecord(t *testing.T) {
	store := newTestStore(t)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	if err := store.Update(testRecord{ID: 1, Name: "Alicia"}); err != nil {
		t.Fatalf("Update() returned an error: %v", err)
	}

	found, err := store.Find(1)
	if err != nil {
		t.Fatalf("Find() returned an error: %v", err)
	}
	if found.Name != "Alicia" {
		t.Fatalf("expected Alicia, got %q", found.Name)
	}
}

func TestUpdateMissingReturnsErrNotFound(t *testing.T) {
	store := newTestStore(t)

	err := store.Update(testRecord{ID: 1, Name: "Alice"})
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteExistingRecord(t *testing.T) {
	store := newTestStore(t)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	if err := store.Delete(1); err != nil {
		t.Fatalf("Delete() returned an error: %v", err)
	}

	_, err := store.Find(1)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestDeleteMissingReturnsErrNotFound(t *testing.T) {
	store := newTestStore(t)

	err := store.Delete(1)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestWhereFiltersRecords(t *testing.T) {
	store := newTestStore(t)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}
	if err := store.Create(testRecord{ID: 2, Name: "Bob"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	results, err := store.Where(func(r testRecord) bool {
		return r.Name == "Bob"
	})
	if err != nil {
		t.Fatalf("Where() returned an error: %v", err)
	}
	if len(results) != 1 || results[0].ID != 2 {
		t.Fatalf("expected only Bob, got %+v", results)
	}
}

func TestDataPersistsAcrossStoreInstances(t *testing.T) {
	path := filepath.Join(t.TempDir(), "records.json")
	store := New[testRecord](path)

	if err := store.Create(testRecord{ID: 1, Name: "Alice"}); err != nil {
		t.Fatalf("Create() returned an error: %v", err)
	}

	reloaded := New[testRecord](path)
	found, err := reloaded.Find(1)
	if err != nil {
		t.Fatalf("Find() returned an error: %v", err)
	}
	if found.Name != "Alice" {
		t.Fatalf("expected Alice, got %q", found.Name)
	}
}
