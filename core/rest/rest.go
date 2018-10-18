package rest

import "errors"

var (
	Restslist []Rest
)

type Rests struct {
	rest []Rest
}

type Rest struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Quality quality `json:"quality"`
}

type quality struct {
	nourriture int `json:"nourriture"`
	salle      int `json:"salle"`
	service    int `json:"service"`
}

func init() {
	Restslist = append(Restslist, Rest{Name: "restaurant 1", ID: len(Restslist)})
	Restslist = append(Restslist, Rest{Name: "restaurant 2", ID: len(Restslist)})
	Restslist = append(Restslist, Rest{Name: "restaurant 3", ID: len(Restslist)})
}
func GetRests() []Rest {
	return Restslist
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
