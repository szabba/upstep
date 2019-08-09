package feature_test

import (
	"context"
	"testing"

	"github.com/szabba/assert"

	. "github.com/szabba/upstep/feature"
)

func Test_MapSet_ZeroValueReportsAllFeaturesToBeOff(t *testing.T) {
	// given
	var ctx context.Context
	name := Named("a-feature")
	set := MapSet{}

	// when
	isOn := set.IsOn(ctx, name)

	// then
	assert.That(
		!isOn,
		t.Errorf, "every feature should be off given a zero value set")
}

func Test_MapSet_BuiltFromNilMapReportsAllFeaturesToBeOff(t *testing.T) {
	// given
	var ctx context.Context
	name := Named("a-feature")
	set := SetFromMap(nil)

	// when
	isOn := set.IsOn(ctx, name)

	// then
	assert.That(
		!isOn,
		t.Errorf, "every feature should be off when the map was nil")
}

func Test_MapSet_BuiltFromEmptyMapReportsAllFeaturesToBeOff(t *testing.T) {
	// given
	var ctx context.Context
	name := Named("a-feature")
	set := SetFromMap(map[Name]bool{})

	// when
	isOn := set.IsOn(ctx, name)

	// then
	assert.That(
		!isOn,
		t.Errorf, "every feature should be off when the map was empty")
}

func Test_MapSet_BuiltFromMapReportsFeatureOnWhenSetSoInTheMap(t *testing.T) {
	// given
	var ctx context.Context
	name := Named("a-feature")
	set := SetFromMap(map[Name]bool{
		name: true,
	})

	// when
	isOn := set.IsOn(ctx, name)

	// then
	assert.That(
		isOn,
		t.Errorf, "feature should be on when set true in the map")
}

func Test_MapSet_BuiltFromMapReportsFeatureOffWhenSetSoInTheMap(t *testing.T) {
	// given
	var ctx context.Context
	name := Named("a-feature")
	set := SetFromMap(map[Name]bool{
		name: false,
	})

	// when
	isOn := set.IsOn(ctx, name)

	// then
	assert.That(
		!isOn,
		t.Errorf, "feature should be off when set false in the map")
}
