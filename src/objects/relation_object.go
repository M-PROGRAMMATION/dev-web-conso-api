package objects

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
	Index []Relation `json:"index"`
}

func (r Relation) GetID() int {
	return r.ID
}

var SeedRelations = []Relation{
	{
		ID: 1,
		DatesLocations: map[string][]string{
			"toulouse-france": {"14-03-2025", "15-03-2025"},
			"lyon-france":     {"02-04-2025"},
			"paris-france":    {"20-05-2025"},
		},
	},
	{
		ID: 2,
		DatesLocations: map[string][]string{
			"sqdqsd-qsdd": {"14-03-2125", "15-03-2025"},
			"sd-sd":       {"02-04-2525"},
			"dd-ddd":      {"20-05-2225"},
		},
	},
}
