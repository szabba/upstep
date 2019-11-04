package docstore

import (
	"github.com/szabba/upstep/planning"
	"gocloud.dev/docstore"
)

type _PlanDoc struct {
	ID               string
	DocumentRevision interface{}
	PlannerID        string
	Name             string
	Status           string
}

func (doc *_PlanDoc) FromDomain(plan *planning.Plan, coll *docstore.Collection) error {
	return nil
}

func (doc *_PlanDoc) ToDomain(coll *docstore.Collection) (*planning.Plan, error) {
	id := planning.PlanIDOf(doc.ID)
	rawRev, err := coll.RevisionToString(doc.DocumentRevision)
	if err != nil {
		return nil, err
	}
	rev := planning.PlanRevisionOf(rawRev)
	plannerID := planning.PlannerIDOf(doc.PlannerID)
	name := planning.PlanName(doc.Name)
	status := planning.PlanStatus(doc.Status)
	plan := planning.NewPlan(id, plannerID, name, status, rev)
	return plan, nil
}
