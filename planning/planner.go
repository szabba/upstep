package planning

import "context"

// A Planner is a person who can create plans and complete them.
type Planner struct{}

// A PlannerID identifies a planner.
type PlannerID struct{ value int64 }

type PlannerRepository interface {
	Get(context.Context, PlannerID) (*Planner, error)
	Save(context.Context, *Planner) error
}

// CreatePlan creates a plan associated with the planner.
func (planner *Planner) CreatePlan(ctx context.Context, svc Service, goal Goal, constraints []Constraint) *Plan {
	return nil
}

// Take ...
//
// TODO: Needs better documentation.
func (planner *Planner) Take(step *Step) error {
	return nil
}

// A TakenStep describes a step that was already undertaken by a planner.
type TakenStep struct {
	planner PlannerID
	step    StepID
}
