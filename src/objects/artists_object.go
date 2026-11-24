package objects

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func (a Artist) GetID() int {
	return a.ID
}

var SeedArtists = []Artist{
	{ID: 1, Name: "Aurore Nomade", CreationDate: 2017,
		Members: []string{"Lina", "Nora", "Sami"}, FirstAlbum: "12-05-2019"},
}
