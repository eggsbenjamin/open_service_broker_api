package models

type Catalog struct {
	Services []*Service `json:"services"`
}
