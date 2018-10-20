package rest

import "errors"

type Rests struct {
	rest []Rest
}

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

func NewRest(name string) *Rest {
	return &Rest{Name: name}
}

func (r *Rests) GetRest(id int) (Rest, error) {
	for _, v := range r.rest {
		if v.ID == id {
			return v, nil
		}
	}
	return Rest{}, errors.New("No id found")
}

func (r *Rests) DeleteRest(id int) []Rest {
	rest := make([]Rest, len(r.rest))
	for _, v := range r.rest {
		if v.ID != id {
			rest = append(rest, v)
		}
	}
	return rest
}

func (r *Rests) AddRest(name string) {
	r.rest = append(r.rest, *NewRest(name))
}

func (r *Rests) UpdateRest(id int, rest Rest) {
	r.rest[id] = rest
}
