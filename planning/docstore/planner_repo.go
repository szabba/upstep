package docstore

import (
	"context"

	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type PlannerRepository struct {
	coll *docstore.Collection
}

func NewPlannerRepository(coll *docstore.Collection) *PlannerRepository {
	return &PlannerRepository{coll}
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
