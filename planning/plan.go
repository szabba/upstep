package planning

import "context"

// A Plan keeps track of progress towards a goal.
type Plan struct{}

// A PlanID identifies a plan.
type PlanID struct{ value int64 }

type PlanRepository interface {
	Get(context.Context, PlanID) (*Plan, error)
	// TODO: Needs a better name.
	// Something like this will be needed for updating plans.
	GetMany(context.Context, PlannerID, []*Plan) error
	Add(context.Context, *Plan) error
	Update(context.Context, *Plan) error
}

// Progress measures the progress towards the plan goal.
func (plan *Plan) Progress() Progress {
	return Progress{}
}

// AvailableSteps returns a slice of steps that can be taken to get closer to the plan goal.
func (plan *Plan) AvailableSteps() []*Step {
	return nil
}
