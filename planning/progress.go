package planning

// A Progress measures how close to reaching the plan goal the planner is.
type Progress struct{}

// Less compared two progresses.
// It returns true only when the left one is definitely less than the right one.
// In ambigous cases it returns false.
func (progress Progress) Less(other Progress) bool {
	return false
}
