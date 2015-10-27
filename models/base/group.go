package models

type Group struct {
	Name       string `json:"name"`
	Active     bool   `json:"active"`
	Extensions `json:"extensions"`
}

type Groups []*Group

func NewGroup(name string) *Group {
	return &Group{Name: name, Active: true}
}

func NewGroupWithExtensions(name string, extensions Extensions) *Group {
	return &Group{name, true, extensions}
}
