package httpglue

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/szabba/upstep/planning"
)

func NewMux(plans *PlanHandlers) *mux.Router {
	r := mux.NewRouter()
	r.Path("/plan/{id}").Methods(http.MethodGet).HandlerFunc(plans.GetOne)
	return r
}

type PlanHandlers struct {
	repo planning.PlanRepository
}

func NewPlans(repo planning.PlanRepository) *PlanHandlers {
	return &PlanHandlers{repo}
}

func (h *PlanHandlers) GetOne(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := planning.PlanIDOf(rawID)
	plan, _ := h.repo.Get(r.Context(), id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}
