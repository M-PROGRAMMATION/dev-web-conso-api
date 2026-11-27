package utils

import (
	"testing"
	"time"
)

func TestNormalizeDateParsesStandardFormat(t *testing.T) {
	date, err := NormalizeDate("14-03-2025")
	if err != nil {
		t.Fatalf("NormalizeDate() returned an error: %v", err)
	}

	want := time.Date(2025, time.March, 14, 0, 0, 0, 0, time.UTC)
	if !date.Equal(want) {
		t.Fatalf("expected %v, got %v", want, date)
	}
}

func TestNormalizeDateStripsStarPrefix(t *testing.T) {
	date, err := NormalizeDate("*14-03-2025")
	if err != nil {
		t.Fatalf("NormalizeDate() returned an error: %v", err)
	}

	want := time.Date(2025, time.March, 14, 0, 0, 0, 0, time.UTC)
	if !date.Equal(want) {
		t.Fatalf("expected %v, got %v", want, date)
	}
}

func TestNormalizeDateInvalidFormatReturnsError(t *testing.T) {
	if _, err := NormalizeDate("2025-03-14"); err == nil {
		t.Fatal("expected an error for a non dd-mm-yyyy date")
	}
}

func TestNormalizeLocationSplitsCityAndCountry(t *testing.T) {
	city, country := NormalizeLocation("new-york-usa")
	if city != "New York" {
		t.Fatalf("expected city %q, got %q", "New York", city)
	}
	if country != "Usa" {
		t.Fatalf("expected country %q, got %q", "Usa", country)
	}
}

func TestNormalizeLocationSingleWord(t *testing.T) {
	city, country := NormalizeLocation("tokyo")
	if city != "Tokyo" {
		t.Fatalf("expected city %q, got %q", "Tokyo", city)
	}
	if country != "" {
		t.Fatalf("expected empty country, got %q", country)
	}
}
