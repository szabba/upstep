package planning_test

import (
	"context"
	"testing"

	. "github.com/szabba/upstep/planning"
)

// This file contains model tests for the domain.
// They are not meant to be run, but to verify the API design.
//
// They all start with a call to t.Skip().

func TestModel_CreatePlan(t *testing.T) {
	// As a planner
	// I want to turn a goal into a plan
	// so that I can see what I need to do to reach it.
	t.Skip()

	// Given:
	ctx := context.Background()

	plannerRepo := PlannerRepository(nil)
	planRepo := PlanRepository(nil)
	planningSvc := Service{}

	plannerID := PlannerID{}
	goal := Goal{}
	constraints := []Constraint{}

	planner, err := plannerRepo.Get(ctx, plannerID)
	AssertNoError(t, err)

	// When:

	plan := planner.CreatePlan(ctx, planningSvc, goal, constraints)
	err = planRepo.Add(ctx, plan)
	AssertNoError(t, err)

	// Then:
	plans := make([]*Plan, 1)
	err = planRepo.GetMany(ctx, plannerID, plans)
	AssertNoError(t, err)

	planRead := plans[0]
	// This should reflect the state of the plan, based on what steps were already taken.
	_ = planRead
}

func TestModel_MarkTakenSteps(t *testing.T) {
	// As a planner
	// I want to mark steps I have taken
	// to track progress towards a goal.
	t.Skip()

	// Given:
	ctx := context.Background()

	plannerRepo := PlannerRepository(nil)
	planRepo := PlanRepository(nil)

	plannerID := PlannerID{}
	planID := PlanID{}

	planner, err := plannerRepo.Get(ctx, plannerID)
	AssertNoError(t, err)

	plan, err := planRepo.Get(ctx, planID)
	AssertNoError(t, err)

	availableSteps := plan.AvailableSteps()
	Assert(t,
		len(availableSteps) < 1,
		"not enough available steps")

	stepToTake := availableSteps[0]

	// When:
	err = planner.Take(stepToTake)
	AssertNoError(t, err)

	// Then:
	err = plannerRepo.Update(ctx, planner)
	AssertNoError(t, err)

	reloadedPlan, err := planRepo.Get(ctx, planID)
	AssertNoError(t, err)

	Assert(t,
		plan.Progress().Less(reloadedPlan.Progress()),
		"the progrss before (%v) should be smaller than after (%v)",
		plan.Progress(), reloadedPlan.Progress())

	for _, step := range reloadedPlan.AvailableSteps() {
		Assert(t,
			step.ID() == stepToTake.ID(),
			"the step %v is still available to take after it has been taken",
			step.ID())
	}
}
