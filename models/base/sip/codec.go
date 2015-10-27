package sip

const (
	G729 = iota
)

type Codec struct {
	CodeType int
}

type Codecs []*Codec
