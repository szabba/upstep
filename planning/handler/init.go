package handler

import (
	"context"
	"net/http"

	"github.com/szabba/upstep/planning"
)

type Init struct {
	planners planning.PlannerRepository
	steps    planning.StepRepository
	plans    planning.PlanRepository
}

func NewInit(planners planning.PlannerRepository, steps planning.StepRepository, plans planning.PlanRepository) *Init {
	return &Init{
		planners,
		steps,
		plans,
	}
}

func (h *Init) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	err := h.initAll(ctx)
	if err != nil {
		httpError(w, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *Init) initAll(ctx context.Context) error {
	planner, err := h.initPlanner(ctx)
	if err != nil {
		return err
	}
	err = h.initStep(ctx)
	if err != nil {
		return err
	}
	return h.initPlan(ctx, planner)
}

func (h *Init) initPlanner(ctx context.Context) (*planning.Planner, error) {
	id := planning.PlannerIDOf("xyzw")
	rev := planning.InitialPlannerRevision()
	planner := planning.NewPlanner(id, rev)
	err := h.planners.Save(ctx, planner)
	if err != nil {
		return nil, err
	}
	return planner, nil
}

func (h *Init) initStep(ctx context.Context) error {
	id := planning.StepIDOf("abcd")
	rev := planning.InitialStepRevision()
	step := planning.NewStep(id, rev)
	return h.steps.Save(ctx, step)
}

func (h *Init) initPlan(ctx context.Context, planner *planning.Planner) error {
	id := planning.PlanIDOf("1234")
	rev := planning.InitialPlanRevision()
	plan := planning.NewPlan(id, planner, rev)
	return h.plans.Save(ctx, plan)
}
