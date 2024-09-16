package model

type InsertParking struct {
	Name string `json:"name"`
}

func NewInsertParking(name string) *InsertParking {
	return &InsertParking{
		Name: name,
	}
}
