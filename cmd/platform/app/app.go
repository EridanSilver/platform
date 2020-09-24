package app

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
)

type Application struct {
}

func NewApp() *Application {
	return &Application{}
}

func (a *Application) Run(ctx context.Context) error {
	r := chi.NewRouter()
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("welcome")); err != nil {
			return
		}
	})
	return http.ListenAndServe(":80", r)
}
