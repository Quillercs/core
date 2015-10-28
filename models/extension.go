package models

type Extension struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	CallerId   string `json:"caller_id"`
	Active     bool   `json:"active"`
	DTMFSource int    `json:"dtmf_source"`
	Tech       `json:"tech"`
}

type Extensions []*Extension

func NewExtension(name, address string, active bool, techType Tech) *Extension {
	return &Extension{
		name,
		address,
		address,
		true,
		DTMF_SOURCE_INBAND_AUDIO,
		techType,
	}
}
