package docstore_test

import (
	"context"
	"testing"

	"gocloud.dev/docstore/memdocstore"

	"github.com/szabba/assert"

	"github.com/szabba/upstep/planning"
	. "github.com/szabba/upstep/planning/docstore"
)

func Test_StepRepository_NotFoundInEmptyRepository(t *testing.T) {
	// given
	repo := newStepRepo(t)

	ctx := context.Background()
	id := planning.StepIDOf("an-id")

	// when
	step, err := repo.Get(ctx, id)

	// then
	assert.That(err == planning.ErrNotFound, t.Errorf, "unexpected error: %s", err)
	assert.That(step == nil, t.Errorf, "unexpected non-nil step: %#v", step)
}

func Test_StepRepository_FoundAfterSave(t *testing.T) {
	// given
	repo := newStepRepo(t)

	ctx := context.Background()
	id := planning.StepIDOf("an-id")
	initRev := planning.InitialStepRevision()
	stepToSave := planning.NewStep(id, initRev)

	// when
	err := repo.Save(ctx, stepToSave)

	// then
	assert.That(err == nil, t.Fatalf, "unexpected error saving step: %s", err)

	stepFound, err := repo.Get(ctx, id)

	assert.That(err == nil, t.Errorf, "unexpected error getting saved step: %s", err)
	assert.That(stepFound != nil, t.Fatalf, "found step is nil")
	assert.That(
		stepFound.ID() == id,
		t.Errorf, "step found has ID %s, want %s", stepFound.ID(), id)
}

func Test_StepRepository_ConflictAfterConcurrentModification(t *testing.T) {
	// given
	repo := newStepRepo(t)

	ctx := context.Background()
	id := planning.StepIDOf("an-id")

	initRev := planning.InitialStepRevision()
	originalStep := planning.NewStep(id, initRev)

	err := repo.Save(ctx, originalStep)
	assert.That(err == nil, t.Fatalf, "could not save step: %s", err)

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

func newStepRepo(t *testing.T) *StepRepository {
	coll, err := memdocstore.OpenCollection("ID", nil)
	assert.That(
		err == nil, t.Fatalf,
		"unexpected error creating step repository: %s", err)

	return NewStepRepository(coll)
}
