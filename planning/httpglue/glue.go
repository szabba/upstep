package httpglue

import (
	"encoding/json"
	"fmt"
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

	plan, err := h.repo.Get(r.Context(), id)
	if err == nil {
		h.writePlan(w, plan)
	} else {
		fmt.Print(err)
	}
}

func (h *PlanHandlers) writePlan(w http.ResponseWriter, plan *planning.Plan) {
	w.Header().Add("Content-Type", "application/json")
	var dto struct {
		Name string
	}
	dto.Name = string(plan.Name())
	json.NewEncoder(w).Encode(dto)
}
