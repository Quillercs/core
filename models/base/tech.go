package base

type Tech struct {
	TechType int `json:"tech_type"`
}

func NewTech(techType int) *Tech {
	return &Tech{techType}
}
