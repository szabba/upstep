package docstore

import (
	"context"

	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type StepRepository struct {
	coll *docstore.Collection
}

func NewStepRepository(coll *docstore.Collection) *StepRepository {
	return &StepRepository{coll}
}

var _ planning.StepRepository = new(StepRepository)

func (repo *StepRepository) Get(ctx context.Context, id planning.StepID) (*planning.Step, error) {
	rawID := id.Value()
	doc := _StepDoc{ID: rawID}

	err := repo.coll.Get(ctx, &doc)
	if err != nil {
		return nil, translateGetError(err)
	}

	return doc.ToDomain(repo.coll)
}

func (repo *StepRepository) Save(ctx context.Context, step *planning.Step) error {
	doc := &_StepDoc{}
	err := doc.FromDomain(step, repo.coll)
	if err != nil {
		return translateSaveError(err)
	}

	err = repo.coll.Put(ctx, doc)
	return translateSaveError(err)
}
