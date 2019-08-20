package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/szabba/upstep/planning"
)

const (
	_PlanActive    = "active"
	_PlanSuspended = "suspended"
	_PlanCompleted = "completed"
)

// Plan is a handler for plan-related activities.
type Plan struct {
	plans planning.PlanRepository
}

func NewPlan(plans planning.PlanRepository) *Plan {
	return &Plan{plans}
}

// Get retrieves a plan based on the {id} mux variable.
func (h *Plan) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawID := vars["id"]
	id := planning.PlanIDOf(rawID)
	plan, err := h.plans.Get(r.Context(), id)

	if err == nil {
		h.respondWithPlan(w, plan)
	} else if err == planning.ErrNotFound {
		httpError(w, http.StatusNotFound)
	} else {
		log.Print(err)
		httpError(w, http.StatusInternalServerError)
	}
}

func (h *Plan) respondWithPlan(w io.Writer, plan *planning.Plan) {
	var dto _PlanDto
	dto.From(plan)
	enc := json.NewEncoder(w)
	err := enc.Encode(dto)
	if err != nil {
		log.Print(err)
	}
}

type _PlanDto struct {
	ID     string
	Name   string
	Status string
	Steps  []_PlanStepDto
}

type _PlanStepDto struct {
	ID          string
	Name        string
	Goal        bool
	Taken       bool
	NextStepIDs []string
}

func (dto *_PlanDto) From(plan *planning.Plan) {
	dto.ID = plan.ID().Value()
}
