package app

func New() *App {
	return &App{}
}

type App struct{}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	return nil
}
