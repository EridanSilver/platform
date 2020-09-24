package app

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

var hand http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("welcome")); err != nil {
		return
	}
}

type Endpoint struct {
	Path        string
	Handler     http.HandlerFunc
	Description string
	Request     interface{}
	Response    interface{}
}

type Application struct {
	endpoints []Endpoint
}

func NewApp() *Application {
	return &Application{
		endpoints: []Endpoint{
			{
				Path:        "/hand",
				Handler:     hand,
				Description: "desc",
				Request:     nil,
				Response:    nil,
			},
		},
	}
}

func (a *Application) Run(ctx context.Context) error {
	r := chi.NewRouter()
	r.Get("/api", hand)
	return http.ListenAndServe(":80", r)
}
