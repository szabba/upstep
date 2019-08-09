package feature

import "context"

// A Name of a feature.
type Name struct{ asString string }

// Named creates a feature with the given name.
func Named(asString string) Name {
	return Name{asString}
}

func (name Name) String() string {
	return name.asString
}

// A ToggleSet says whether a feature is on or off based on it's name.
type ToggleSet interface {
	// IsOn tells whether a feature with the given name is on.
	IsOn(context.Context, Name) bool
}

// A MapSet is a toggle set based on a map.
//
// The zero value is useable, but considers all features to be off.
type MapSet struct {
	knownFlags map[Name]bool
}

// SetFromMap creates a toggle set from a map.
// If a name is present in the map, the feature toggle state is read from it.
// If a name is not present, the feature is considerer off.
//
// After the MapSet is created, modifications to the map are not reflected.
func SetFromMap(knownFlags map[Name]bool) MapSet {
	set := MapSet{}
	set.knownFlags = make(map[Name]bool, len(knownFlags))
	for name, on := range knownFlags {
		set.knownFlags[name] = on
	}
	return set
}

var _ ToggleSet = MapSet{}

// IsOn implements the ToggleSet interface.
func (set MapSet) IsOn(_ context.Context, name Name) bool {
	if set.knownFlags == nil {
		return false
	}
	return set.knownFlags[name]
}
