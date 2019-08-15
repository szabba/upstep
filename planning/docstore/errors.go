package docstore

import (
	"gocloud.dev/gcerrors"

	"github.com/szabba/upstep/planning"
)

func translateGetError(err error) error {
	if gcerrors.Code(err) == gcerrors.NotFound {
		return planning.ErrNotFound
	}
	return err
}

func translateSaveError(err error) error {
	if gcerrors.Code(err) == gcerrors.FailedPrecondition {
		return planning.ErrConflict
	}
	return err
}
