package api

import "github.com/uiansol/task-follow-up/internal/api/handlers"

func configHandlers() *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	return &AppHandlers{
		pingHandler: pingHandler,
	}
}
