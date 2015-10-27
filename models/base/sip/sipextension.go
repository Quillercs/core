package sip

type SIPExtension struct {
	base.Extension
	URI         string
	BindAddress string
	BindPort    int

	CodecsAllow    Codecs
	CodecsDisallow Codecs
}
