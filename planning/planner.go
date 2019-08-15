package planning

import "context"

// A PlannerRepository persistently stores and later retrieves planners.
type PlannerRepository interface {

	// Get retrieves the planner with the given ID from a backing storage.
	//
	// When a planner is not found the error will be ErrNotFound.
	Get(context.Context, PlannerID) (*Planner, error)

	// Save stores the given planner.
	//
	// The planner must have an assigned ID.
	Save(context.Context, *Planner) error
}

// A Planner is a person who can create plans and complete them.
type Planner struct {
	id  PlannerID
	rev PlannerRevision
}

// NewPlanner creates a new planner.
func NewPlanner(id PlannerID, rev PlannerRevision) *Planner {
	return &Planner{id: id, rev: rev}
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
