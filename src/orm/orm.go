package orm

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	ErrNotFound  = errors.New("record not found")
	ErrDuplicate = errors.New("record already exists")
)

type Model interface {
	GetID() int
}

type Store[T Model] struct {
	path string
	mu   sync.RWMutex
}

func New[T Model](path string) *Store[T] {
	return &Store[T]{path: path}
}

func (s *Store[T]) load() ([]T, error) {
	records := make([]T, 0)

	data, err := os.ReadFile(s.path)

	if errors.Is(err, os.ErrNotExist) {
		return records, nil
	}

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (s *Store[T]) save(records []T) error {
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.CreateTemp(dir, ".orm-*.json")
	if err != nil {
		return err
	}

	defer os.Remove(file.Name())
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	if err := file.Sync(); err != nil {
		return err
	}

	if err := file.Close(); err != nil {
		return err
	}

	return os.Rename(file.Name(), s.path)
}

func (s *Store[T]) All() ([]T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.load()
}

func (s *Store[T]) Find(id int) (T, error) {
	var zero T

	records, err := s.All()
	if err != nil {
		return zero, err
	}

	for _, record := range records {
		if record.GetID() == id {
			return record, nil
		}
	}

	return zero, ErrNotFound
}

func (s *Store[T]) Create(record T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.load()
	if err != nil {
		return err
	}

	for _, existing := range records {
		if existing.GetID() == record.GetID() {
			return ErrDuplicate
		}
	}

	records = append(records, record)

	return s.save(records)
}

func (s *Store[T]) Update(record T) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.load()
	if err != nil {
		return err
	}

	for i, existing := range records {
		if existing.GetID() == record.GetID() {
			records[i] = record
			return s.save(records)
		}
	}

	return ErrNotFound
}

func (s *Store[T]) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.load()
	if err != nil {
		return err
	}

	for i, record := range records {
		if record.GetID() == id {
			records = append(records[:i], records[i+1:]...)
			return s.save(records)
		}
	}

	return ErrNotFound
}

func (s *Store[T]) Where(condition func(T) bool) ([]T, error) {
	records, err := s.All()
	if err != nil {
		return nil, err
	}

	results := make([]T, 0)

	for _, record := range records {
		if condition(record) {
			results = append(results, record)
		}
	}

	return results, nil
}
