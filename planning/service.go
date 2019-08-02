package planning

import "context"

type Service struct {
	repo StepRepository
}

func (svc *Service) Find(ctx context.Context, goal Goal, steps []*Step) error {
	return nil
}
