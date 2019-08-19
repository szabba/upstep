package docstore

import (
	"context"
	"fmt"

	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type PlannerRepository struct {
	coll *docstore.Collection
}

func NewPlannerRepository() (*PlannerRepository, error) {
	ctx := context.Background()
	url := envOrElse("PLANNER_COLLECTION_URL", "mem://planners/ID")
	coll, err := docstore.OpenCollection(ctx, url)
	repo := &PlannerRepository{coll}
	if err != nil {
		return nil, fmt.Errorf("cannot create %T: %s", repo, err)
	}
	return repo, err
}

var _ planning.PlannerRepository = new(PlannerRepository)

func (repo *PlannerRepository) Get(ctx context.Context, id planning.PlannerID) (*planning.Planner, error) {
	rawID := id.Value()
	doc := _PlannerDoc{ID: rawID}

	err := repo.coll.Get(ctx, &doc)
	if err != nil {
		return nil, translateGetError(err)
	}

	return doc.ToDomain(repo.coll)
}

func (repo *PlannerRepository) Save(ctx context.Context, planner *planning.Planner) error {
	doc := &_PlannerDoc{}
	err := doc.FromDomain(planner, repo.coll)
	if err != nil {
		return translateSaveError(err)
	}

	err = repo.coll.Put(ctx, doc)
	return translateSaveError(err)
}
