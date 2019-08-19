package docstore

import (
	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type _PlannerDoc struct {
	ID               string
	DocstoreRevision interface{}
}

func (doc *_PlannerDoc) SetID(id planning.PlannerID) {
	doc.ID = id.Value()
}

func (doc *_PlannerDoc) FromDomain(planner *planning.Planner, coll *docstore.Collection) error {
	doc.SetID(planner.ID())
	if !planner.Revision().IsInitial() {
		revString := planner.Revision().Value()
		rev, err := coll.StringToRevision(revString)
		doc.DocstoreRevision = rev
		return err
	}
	return nil
}

func (doc _PlannerDoc) ToDomain(coll *docstore.Collection) (*planning.Planner, error) {
	id := planning.PlannerIDOf(doc.ID)
	revString, err := coll.RevisionToString(doc.DocstoreRevision)
	if err != nil {
		return nil, err
	}
	rev := planning.PlannerRevisionOf(revString)
	planner := planning.NewPlanner(id, rev)
	return planner, nil
}
