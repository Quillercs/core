package models

const (
	DTMF_SOURCE_UNKNOWN = iota
	DTMF_SOURCE_INBAND_AUDIO
	DTMF_SOURCE_RTP
	DTMF_SOURCE_ENDPOINT
	DTMF_SOURCE_APP
)

type DTMF struct {
	Digit    byte `json:"digit"`
	Duration int  `json:"duration"`
	Source   int  `json:"source"`
}

func NewDTMF(digit byte, duration, source int) *DTMF {
	return &DTMF{digit, duration, source}
}
