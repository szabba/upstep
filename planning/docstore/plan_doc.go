package docstore

import (
	"github.com/szabba/upstep/planning"
	"gocloud.dev/docstore"
)

type _PlanDoc struct {
	ID string
}

func (doc *_PlanDoc) ToDomain(coll *docstore.Collection) (*planning.Plan, error) {
	return nil, nil
}

func (doc *_PlanDoc) FromDomain(plan *planning.Plan, coll *docstore.Collection) error {
	return nil
}
