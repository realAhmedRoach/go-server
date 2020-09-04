package controller

import "sces/mgmt"

type Application struct {
	Sukuk mgmt.Service
}

func (app *Application) SukukManager() mgmt.Service {
	return app.Sukuk
}
