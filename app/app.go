package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/go-chi/chi"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

func getDecoder(e Endpoint) DecodeRequestFunc {
	if e.Request == nil {
		return EmptyRequest
	}

	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		prototype := getCopy(e.Request)
		if len(body) == 0 {
			return prototype, nil
		}

		err = json.Unmarshal(body, &prototype)
		if err != nil {
			panic(err)
		}

		return prototype, nil
	}
}

func getCopy(obj interface{}) interface{} {
	indirect := reflect.New(reflect.TypeOf(obj))
	return indirect.Interface()
}

var hand http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("welcome")); err != nil {
		return
	}
}

type Endpoint struct {
	Path        string
	Method      string
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
			},
		},
	}
}

func (a *Application) Run(ctx context.Context) error {
	r := chi.NewRouter()
	for _, endpoint := range a.endpoints {
		if endpoint.Method == "POST" {
			r.Post(endpoint.Path, endpoint.Handler)
		}
		if endpoint.Method == "GET" {
			r.Get(endpoint.Path, endpoint.Handler)
		}
	}
	return http.ListenAndServe(":8080", r)
}

// EmptyRequest ...
func EmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req interface{}
	return req, nil
}
