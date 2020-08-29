package mgmt

type Application struct {
	Sukuk Service
}

func (app *Application) SukukManager() Service {
	return app.Sukuk
}
