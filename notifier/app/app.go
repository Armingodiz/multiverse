package app

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start(addr string) error {
	return nil
}
