package base

type Args map[string]string

type AppInterface interface {
	Load(args Args) error
	Execute(args Args) error
	Unload()
}

type App struct {
	UID          string
	Name         string
	MajorVersion string
	MinorVersion string
	Loaded       bool
}

func NewAppBase(uid, name, mjVersion, mnVersion string) *AppBase {
	return &AppBase{uid, name, mjVersion, mnVersion, false}
}
