package objects

type Stats struct {
	TotalArtists int            `json:"totalArtists"`
	TotalEvents  int            `json:"totalEvents"`
	EventsByCity map[string]int `json:"eventsByCity"`
}
