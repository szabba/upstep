package docstore_test

import (
	"context"
	"testing"

	"gocloud.dev/docstore/memdocstore"

	"github.com/szabba/assert"

	"github.com/szabba/upstep/planning"
	. "github.com/szabba/upstep/planning/docstore"
)

func Test_PlannerRepository_NotFoundInEmptyRepository(t *testing.T) {
	// given
	repo := newPlannerRepo(t)

	ctx := context.Background()
	id := planning.PlannerIDOf("an-id")

	// when
	planner, err := repo.Get(ctx, id)

	// then
	assert.That(err == planning.ErrNotFound, t.Errorf, "unexpected error: %s", err)
	assert.That(planner == nil, t.Errorf, "unexpected non-nil planner: %#v", planner)
}

func Test_PlannerRepository_FoundAfterSave(t *testing.T) {
	// given
	repo := newPlannerRepo(t)

	ctx := context.Background()
	id := planning.PlannerIDOf("an-id")
	initRev := planning.InitialPlannerRevision()
	plannerToSave := planning.NewPlanner(id, initRev)

	// when
	err := repo.Save(ctx, plannerToSave)

	// then
	assert.That(err == nil, t.Fatalf, "unexpected error saving planner: %s", err)

	plannerFound, err := repo.Get(ctx, id)

	assert.That(err == nil, t.Errorf, "unexpected error getting saved planner: %s", err)
	assert.That(plannerFound != nil, t.Fatalf, "found planner is nil")
	assert.That(
		plannerFound.ID() == id,
		t.Errorf, "planner found has ID %s, want %s", plannerFound.ID(), id)
}

func Test_PlannerRepository_ConflictAfterConcurrentModification(t *testing.T) {
	// given
	repo := newPlannerRepo(t)

	ctx := context.Background()
	id := planning.PlannerIDOf("an-id")

	initRev := planning.InitialPlannerRevision()
	originalPlanner := planning.NewPlanner(id, initRev)

	err := repo.Save(ctx, originalPlanner)
	assert.That(err == nil, t.Fatalf, "could not save planner: %s", err)

	firstCopy, err := repo.Get(ctx, id)
	assert.That(err == nil, t.Fatalf, "could not find first copy: %s", err)

	secondCopy, err := repo.Get(ctx, id)
	assert.That(err == nil, t.Fatalf, "could not find second copy: %s", err)

	err = repo.Save(ctx, firstCopy)
	assert.That(err == nil, t.Fatalf, "could not save first copy: %s", err)

	// when
	err = repo.Save(ctx, secondCopy)

	// then
	assert.That(err == planning.ErrConflict, t.Errorf, "got %q, wanted %q", err, planning.ErrConflict)
}

func newPlannerRepo(t *testing.T) *PlannerRepository {
	coll, err := memdocstore.OpenCollection("ID", nil)
	assert.That(
		err == nil, t.Fatalf,
		"unexpected error creating planner repository: %s", err)

	return NewPlannerRepository(coll)
}
