package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/szabba/upstep/planning"
)

const (
	_IDParam = "id"
)

type StepHandler struct {
	repo planning.StepRepository
}

func (h *StepHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := planning.StepIDOf(r.Form.Get(_IDParam))
		step, err := h.repo.Get(r.Context(), id)
		if err == planning.ErrNotFound {
			httpError(w, http.StatusNotFound)
		} else {
			_ = step
			var dto struct {
			}
			err = json.NewEncoder(w).Encode(&dto)
			if err != nil {
				log.Print(err)
			}
		}
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func httpError(w http.ResponseWriter, code int) {
	msg := http.StatusText(code)
	http.Error(w, msg, code)
}
