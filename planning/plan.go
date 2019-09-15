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
	name      PlanName
	status    PlanStatus
}

// A PlanName is a name given by the planner to their plan.
type PlanName string

func (n PlanName) String() string { return string(n) }

// NewPlan creeates a new plan.
func NewPlan(id PlanID, plannerID PlannerID, name PlanName, status PlanStatus, rev PlanRevision) *Plan {
	return &Plan{
		id:        id,
		rev:       rev,
		plannerID: plannerID,
		name:      name,
		status:    status,
	}
}

// Name is the name the planner gave to the plan.
func (plan *Plan) Name() PlanName { return plan.name }

// Status is the current status of the plan.
func (plan *Plan) Status() PlanStatus { return plan.status }

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

type PlanStatus string

const (
	PlanActive    = PlanStatus("Active")
	PlanSuspended = PlanStatus("Suspended")
	PlanComplete  = PlanStatus("Complete")
)

func (s PlanStatus) Valid() bool {
	return s == PlanActive || s == PlanSuspended || s == PlanComplete
}

func (s PlanStatus) String() string {
	if !s.Valid() {
		return "UnknwonPlanStatus:" + string(s)
	}
	return string(s)
}
