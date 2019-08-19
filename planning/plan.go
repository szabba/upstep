package planning

import "context"

// A PlanRepository persistently stores and later retrieves plans.
type PlanRepository interface {

	// Get retrieves the plan with the given ID from a backing storage.
	//
	// When a plan is not found the error will be ErrNotFound.
	Get(context.Context, PlanID) (*Plan, error)

	// GetMany retrieves the plans owned by a planner with some ID.
	GetMany(context.Context, PlannerID, []*Plan) error

	// Save stores the given plan.
	//
	// The plan must have an assigned ID.
	Save(context.Context, *Plan) error
}

// A Plan keeps track of progress towards a goal.
type Plan struct {
	id        PlanID
	rev       PlanRevision
	plannerID PlannerID
}

// NewPlan creeates a new plan.
func NewPlan(id PlanID, planner *Planner, rev PlanRevision) *Plan {
	return &Plan{
		id:        id,
		rev:       rev,
		plannerID: planner.ID(),
	}
}

// Progress measures the progress towards the plan goal.
//
// TODO: Implement.
func (plan *Plan) Progress() Progress {
	return Progress{}
}

// AvailableSteps returns a slice of steps that can be taken to get closer to the plan goal.
func (plan *Plan) AvailableSteps() []*Step {
	return nil
}
