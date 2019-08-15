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
