package sip

type SIPExtension struct {
	URI         string
	BindAddress string
	BindPort    int

	CodecsAllow    Codecs
	CodecsDisallow Codecs
}
