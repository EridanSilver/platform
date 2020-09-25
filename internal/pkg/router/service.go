package router

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"
)

type Router struct {
	Endpoints []Endpoint
}

func NewRouter() *Router {
	return &Router{}
}

type Endpoint struct {
	Path        string
	Method      string
	Controller  Controller
	Description string
	Request     interface{}
	Response    interface{}
}

type Controller func(context.Context, interface{}) (interface{}, error)

func (e Endpoint) handleRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := e.getDecoder()
		request, err := decoder(context.Background(), r)
		if err != nil {
			panic(err)
			return
		}
		response, err := e.Controller(context.Background(), request)
		if err != nil {
			panic(err)
			return
		}
		err = EncodeJSONResponse(context.Background(), w, response)
		if err != nil {
			panic(err)
			return
		}
	}
}

func (r *Router) ListenAndServe() error {
	rt := chi.NewRouter()
	for _, endpoint := range r.Endpoints {
		if endpoint.Method == "POST" {
			rt.Post(endpoint.Path, endpoint.handleRequest())
		}
		if endpoint.Method == "GET" {
			rt.Get(endpoint.Path, endpoint.handleRequest())
		}
	}
	return http.ListenAndServe(":8080", rt)
}
