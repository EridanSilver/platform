package app

import (
	"context"
	"github.com/EridanSilver/platform/internal/pkg/handlers"
	"github.com/EridanSilver/platform/internal/pkg/router"
)

type Application struct {
	Router         *router.Router
	HandlerService *handlers.SomeHandlerService
}

func NewApp() *Application {
	return &Application{
		Router:         router.NewRouter(),
		HandlerService: handlers.NewHandlerService(),
	}
}

func (a *Application) Run(ctx context.Context) error {
	a.Router.Endpoints = []router.Endpoint{
		{
			Path:        "/hand",
			Method:      "POST",
			Controller:  a.HandlerService.Hand,
			Description: "desc",
			Request:     nil,
			Response:    nil,
		},
		{
			Path:        "/api",
			Method:      "GET",
			Controller:  a.HandlerService.Hand2,
			Description: "",
			Request:     nil,
			Response:    nil,
		}, {
			Path:        "/test",
			Method:      "POST",
			Controller:  a.HandlerService.Hand2,
			Description: "",
			Request:     nil,
			Response:    nil,
		},
		{
			Path:        "/get-by-id",
			Method:      "POST",
			Controller:  a.HandlerService.GetByID,
			Description: "",
			Request:     nil,
			Response:    nil,
		},
	}
	return a.Router.ListenAndServe()
}
