package docstore

import (
	"context"

	"github.com/szabba/upstep/planning"
)

func (repo *PlanRepository) GetMany(_ context.Context, _ planning.PlannerID, _ []*planning.Plan) error {
	// TODO: implement properly.
	return nil
}
