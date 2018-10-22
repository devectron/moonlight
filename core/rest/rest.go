package rest

type Rest struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Image   string  `json:"image"`
	Quality Quality `json:"quality"`
}

type Quality struct {
	Nourriture int `json:"nourriture"`
	Salle      int `json:"salle"`
	Service    int `json:"service"`
}
