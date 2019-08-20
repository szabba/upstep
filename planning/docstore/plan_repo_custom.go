package docstore

import (
	"context"

	"github.com/szabba/upstep/planning"
)

func (repo *PlanRepository) GetMany(ctx context.Context, id planning.PlannerID, plans []*planning.Plan) error {
	return nil
}
