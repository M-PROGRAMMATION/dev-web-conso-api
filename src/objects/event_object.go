package objects

import "time"

type Event struct {
	ArtistID   int       `json:"artistId"`
	ArtistName string    `json:"artistName"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	Date       time.Time `json:"date"`
}
