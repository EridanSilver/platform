package app

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error

func (e Endpoint) getDecoder() DecodeRequestFunc {
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

// EmptyRequest ...
func EmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req interface{}
	return req, nil
}

type StatusCoder interface {
	StatusCode() int
}

type Headerer interface {
	Headers() http.Header
}

func EncodeJSONResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if headerer, ok := response.(Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusOK
	if sc, ok := response.(StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func getCopy(obj interface{}) interface{} {
	indirect := reflect.New(reflect.TypeOf(obj))
	return indirect.Interface()
}
