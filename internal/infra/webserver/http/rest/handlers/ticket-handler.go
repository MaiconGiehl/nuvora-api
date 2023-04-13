package handlers

import (
	di "github.com/maicongiehl/nuvora-api/configs/di"
)

type TicketHandler struct {
	app *di.App
}

func NewTicketHandler(
	app *di.App,
) *TicketHandler {
	return &TicketHandler{
		app: app,
	}
}
