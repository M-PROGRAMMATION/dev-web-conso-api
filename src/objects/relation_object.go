package objects

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

var relations = RelationIndex{
	Index: []Relation{
		{
			ID: 1, // même ID que l'artiste concerné
			DatesLocations: map[string][]string{
				"toulouse-france": {"14-03-2025", "15-03-2025"},
				"lyon-france":     {"02-04-2025"},
				"paris-france":    {"20-05-2025"},
			},
		},
	},
}
