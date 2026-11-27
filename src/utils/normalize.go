package utils

import (
	"fmt"
	"strings"
	"time"
)

// DateLayout est le format "jj-mm-aaaa" utilisé par les données locales et
// par l'API Groupie Trackers (qui préfixe parfois la date d'un "*").
const DateLayout = "02-01-2006"

func NormalizeDate(raw string) (time.Time, error) {
	cleaned := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(raw), "*"))

	date, err := time.Parse(DateLayout, cleaned)
	if err != nil {
		return time.Time{}, fmt.Errorf("date invalide %q: %w", raw, err)
	}

	return date, nil
}

// NormalizeLocation transforme une clé "ville-pays" (ex: "new-york-usa") en
// un couple ville/pays lisible ("New York", "Usa").
func NormalizeLocation(raw string) (city, country string) {
	parts := strings.Split(strings.TrimSpace(raw), "-")
	if len(parts) < 2 {
		return titleCase(strings.Join(parts, " ")), ""
	}

	country = titleCase(parts[len(parts)-1])
	city = titleCase(strings.Join(parts[:len(parts)-1], " "))

	return city, country
}

func titleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}
