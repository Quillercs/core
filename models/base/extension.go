package base

type Extension struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Active     bool   `json:"active"`
	CallerId   string `json:"caller_id"`
	DTMFSource int    `json:"dtmf_source"`
	Tech       `json:"tech"`
}

type Extensions []*Extension

func NewExtension(name, address, string, active bool, techType TechType) *Extension {
	return &Extension{name, address, true, techType}
}
