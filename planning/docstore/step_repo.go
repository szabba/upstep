package docstore

import (
	"context"
	"fmt"

	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type StepRepository struct {
	coll *docstore.Collection
}

func NewStepRepository() (*StepRepository, error) {
	ctx := context.Background()
	url := envOrElse("STEP_COLLECTION_URL", "mem://ID")
	coll, err := docstore.OpenCollection(ctx, url)
	repo := &StepRepository{coll}
	if err != nil {
		return nil, fmt.Errorf("cannot create %T: %s", repo, err)
	}
	return repo, err
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
