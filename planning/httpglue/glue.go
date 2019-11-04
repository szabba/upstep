package httpglue

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/szabba/upstep/planning"
)

const (
	_ContentType     = "Content-Type"
	_ApplicationJSON = "application/json"
)

// NewRouter builds a router for all the methods.
func NewRouter(plans *PlanHandler) *mux.Router {
	r := mux.NewRouter()
	r.Path("/plan/{id}").Methods(http.MethodGet).HandlerFunc(plans.GetOne)
	return r
}

// A PlanHandler collects http.Handler methods for dealing with plans.
type PlanHandler struct {
	repo planning.PlanRepository
}

// NewPlanHandler creates a new plan handler.
func NewPlanHandler(repo planning.PlanRepository) *PlanHandler {
	return &PlanHandler{repo}
}

// GetOne implements retrieving a single plan by it's ID.
func (h *PlanHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id := planning.PlanIDOf(rawID)

	plan, err := h.repo.Get(r.Context(), id)

	if err == nil {
		log.Info().
			Str("plan-id", plan.ID().Value()).
			Msg("found plan")
		h.writePlan(w, plan)

	} else if errors.Is(err, planning.ErrNotFound) {
		writeError(w, http.StatusNotFound)

	} else {
		log.Error().Err(err)
		writeError(w, http.StatusInternalServerError)
	}
}

func (h *PlanHandler) writePlan(w http.ResponseWriter, plan *planning.Plan) {
	w.Header().Add(_ContentType, _ApplicationJSON)
	var dto struct {
		ID     string
		Name   string
		Status string
	}
	dto.ID = plan.ID().Value()
	dto.Name = plan.Name().String()
	dto.Status = plan.Status().String()
	json.NewEncoder(w).Encode(dto)
}

func writeError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
