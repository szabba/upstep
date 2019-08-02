package planning

import "context"

// A Step is something that can be completed en route to a goal.
//
// Examples of steps include:
// reading a book,
// taking a course / workshop,
// experience using a technique on a project.
type Step struct{}

// A StepID identifies a step.
type StepID struct{ value int64 }

type StepRepository interface {
	Get(context.Context, StepID) (*Step, error)
	Update(context.Context, *Step) error
}

// ID returns the identifier of a step.
func (step *Step) ID() StepID {
	return StepID{}
}
