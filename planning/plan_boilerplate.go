// Code generated by upstep/codegen; DO NOT EDIT.

package planning

// ID is the unique ID of the Plan.
func (agg *Plan) ID() PlanID { return agg.id }

// Revision is the revision of the Plan.
func (agg *Plan) Revision() PlanRevision { return agg.rev }

// SetRevision sets the Plan revision.
func (agg *Plan) SetRevision(rev PlanRevision) {
	agg.rev = rev
}

// A Plan identifies a Plan.
type PlanID struct{ value string }

// PlanIDOf creates a PlanID from a raw string value.
func PlanIDOf(value string) PlanID {
	return PlanID{value}
}

// Value is the raw string value of the PlanID.
func (id PlanID) Value() string { return id.value }

// A PlanRevision is used to track versions of a Plan stored in a repository.
// The persistence layer uses these to detect concurrent modifications.
type PlanRevision struct{ value string }

// InitialPlanRevision returns the revision for a Plan that was not persisted yet.
func InitialPlanRevision() PlanRevision {
	return PlanRevision{}
}

// PlanRevisionOf turns a string into a Plan revision.
// It is used by the persistence code for concurrency control.
func PlanRevisionOf(value string) PlanRevision {
	return PlanRevision{value}
}

// IsInitial returns true when the containing Plan was not stored yet.
func (rev PlanRevision) IsInitial() bool { return rev.value == "" }

// Value is the raw string value of the Plan revision.
func (rev PlanRevision) Value() string { return rev.value }