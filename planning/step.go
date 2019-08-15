package planning

import "context"

// A StepRepository persistently stores and later retrieves steps.
type StepRepository interface {

	// Get retrieves the step with the given ID from a backing storage.
	//
	// When a step is not found the error will be ErrNotFound.
	Get(context.Context, StepID) (*Step, error)

	// Save stores the given step.
	//
	// The step must have an assigned ID.
	Save(context.Context, *Step) error
}

// A StepID identifies a step.
type StepID struct{ value string }

// StepIDOf creates a step ID from a raw string value.
func StepIDOf(value string) StepID {
	return StepID{value}
}

// Value is raw string value of the step ID.
func (id StepID) Value() string { return id.value }

// A StepRevision is used to track versions of a step stored in a repository.
// The persistence layer uses these to detect concurrent modifications.
type StepRevision struct{ value string }

// InitialStepRevision returns the step revision for a step that was not persisted yet.
func InitialStepRevision() StepRevision {
	return StepRevision{}
}

// StepRevisionOf turns a string into a StepRevision.
// It is used by the persistence code for concurrency control.
func StepRevisionOf(value string) StepRevision {
	return StepRevision{value}
}

// IsInitial returns true when the containing step was not stored yet.
func (rev StepRevision) IsInitial() bool { return rev.value == "" }

// Value is the raw string value of the step revision.
func (rev StepRevision) Value() string { return rev.value }

// A Step is something that can be completed en route to a goal.
//
// Examples of steps include:
// reading a book,
// taking a course / workshop,
// experience using a technique on a project.
type Step struct {
	id  StepID
	rev StepRevision
}

// NewStep creates a new step.
func NewStep(id StepID, rev StepRevision) *Step {
	return &Step{id: id, rev: rev}
}

// ID returns the identifier of a step.
func (step *Step) ID() StepID { return step.id }

// Revision is the revision of the step.
func (step *Step) Revision() StepRevision { return step.rev }
