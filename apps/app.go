package apps

import "github.com/quillercs/core/event"

type Args map[string]string

type AppInterface interface {
	Load(*event.EventMachine) error
	Execute(args Args) error
	Unload()
}

type App struct {
	UID          string
	Name         string
	MajorVersion string
	MinorVersion string
	Loaded       bool
	*event.EventMachine
}

func NewAppBase(uid, name, mjVersion, mnVersion string) *AppBase {
	return &AppBase{uid, name, mjVersion, mnVersion, false}
}
