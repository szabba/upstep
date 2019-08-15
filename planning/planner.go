package planning

import "context"

// A PlannerRepository persistently stores and later retrieves planners.
type PlannerRepository interface {
	Get(context.Context, PlannerID) (*Planner, error)
	Save(context.Context, *Planner) error
}

// A PlannerID identifies a planner.
type PlannerID struct{ value string }

// PlannerIDOf creates a planner ID from an opaque value.
func PlannerIDOf(value string) PlannerID {
	return PlannerID{value}
}

// Value is raw string value of the planner ID.
func (id PlannerID) Value() string { return id.value }

// A PlannerRevision is used to track versions of a planner stored in a repository.
// The persistence layer uses these to detect concurrent modifications.
type PlannerRevision struct{ value string }

// InitialPlannerRevision returns the planner revision for a planner that was not persisted yet.
func InitialPlannerRevision() PlannerRevision {
	return PlannerRevision{}
}

// PlannerRevisionOf turns a string into a PlannerRevision.
// It is used by the persistence code while loading a stored planner.
func PlannerRevisionOf(value string) PlannerRevision {
	return PlannerRevision{value}
}

// IsInitial tells whether the planner was stored yet.
func (rev PlannerRevision) IsInitial() bool { return rev.value == "" }

// Value is the raw string value of the planner revision.
func (rev PlannerRevision) Value() string { return rev.value }

// A Planner is a person who can create plans and complete them.
type Planner struct {
	id  PlannerID
	rev PlannerRevision
}

// NewPlanner creates a new planner.
func NewPlanner(id PlannerID, rev PlannerRevision) *Planner {
	return &Planner{id: id, rev: rev}
}

// ID is the unique ID of a planner.
func (planner *Planner) ID() PlannerID { return planner.id }

// Revision is the revision of the planner.
func (planner *Planner) Revision() PlannerRevision { return planner.rev }

// SetRevision sets the planner revision.
func (planner *Planner) SetRevision(rev PlannerRevision) {
	planner.rev = rev
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
