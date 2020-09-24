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

func GetByID(context.Context, interface{}) (interface{}, error) {
	return "data", nil
}

var getByIDEndpoint = Endpoint{
	Path:        "/get-by-id",
	Method:      "POST",
	Handler:     nil,
	Controller:  GetByID,
	Description: "",
	Request:     nil,
	Response:    nil,
}

var hand2 http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	e := Endpoint{
		Path:        "",
		Method:      "",
		Handler:     nil,
		Controller:  controller,
		Description: "",
		Request:     nil,
		Response:    nil,
	}

	f := e.handleRequest()
	f(w, r)
}

type Controller func(context.Context, interface{}) (interface{}, error)

var controller Controller = func(context.Context, interface{}) (interface{}, error) {
	return "test2", nil
}

type Endpoint struct {
	Path        string
	Method      string
	Handler     http.HandlerFunc
	Controller  Controller
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
				Method:      "POST",
				Handler:     hand,
				Description: "desc",
				Request:     nil,
				Response:    nil,
			},
			{
				Path:        "/api",
				Method:      "GET",
				Handler:     hand,
				Description: "",
				Request:     nil,
				Response:    nil,
			}, {
				Path:        "/test",
				Method:      "POST",
				Handler:     hand2,
				Controller:  nil,
				Description: "",
				Request:     nil,
				Response:    nil,
			}, getByIDEndpoint,
		},
	}
}

func (a *Application) Run(ctx context.Context) error {
	r := chi.NewRouter()
	for _, endpoint := range a.endpoints {
		if endpoint.Method == "POST" {
			r.Post(endpoint.Path, endpoint.handleRequest())
		}
		if endpoint.Method == "GET" {
			r.Get(endpoint.Path, endpoint.Handler)
		}
	}
	return http.ListenAndServe(":8080", r)
}

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
