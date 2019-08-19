package docstore

import (
	"gocloud.dev/docstore"

	"github.com/szabba/upstep/planning"
)

type _StepDoc struct {
	ID               string
	DocstoreRevision interface{}
}

func (doc _StepDoc) ToDomain(coll *docstore.Collection) (*planning.Step, error) {
	id := planning.StepIDOf(doc.ID)
	revString, err := coll.RevisionToString(doc.DocstoreRevision)
	if err != nil {
		return nil, err
	}
	rev := planning.StepRevisionOf(revString)
	step := planning.NewStep(id, rev)
	return step, nil
}

func (doc *_StepDoc) FromDomain(step *planning.Step, coll *docstore.Collection) error {
	doc.ID = step.ID().Value()
	if !step.Revision().IsInitial() {
		revString := step.Revision().Value()
		rev, err := coll.StringToRevision(revString)
		doc.DocstoreRevision = rev
		return err
	}
	return nil
}
