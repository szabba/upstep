package planning

import "errors"

var (
	// ErrNotFound indicates that an object was not found it a repository.
	ErrNotFound = errors.New("not found")
	// ErrConflict indicates that an update to an object failed due to concurrent modification.
	ErrConflict = errors.New("conflict")
)
