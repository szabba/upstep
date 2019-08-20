package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

const IDVariable = "id"

func NewMux(init *Init, plan *Plan) http.Handler {
	r := mux.NewRouter()
	r.Path("/init").Methods(http.MethodPost).Handler(init)
	r.Path("/plan/{id}").Methods(http.MethodGet).HandlerFunc(plan.Get)
	return r
}
