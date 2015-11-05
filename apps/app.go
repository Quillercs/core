package apps

type Args map[string]string

type AppInterface interface {
	Load() error
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

func NewApp(uid, name, mjVersion, mnVersion string) *App {
	return &App{uid, name, mjVersion, mnVersion, false}
}
