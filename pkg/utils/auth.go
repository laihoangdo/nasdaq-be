package utils

import (
	"context"

	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/logger"
)

// Validate is usertest from owner of content
func ValidateIsOwner(ctx context.Context, creatorID string, logger logger.Logger) error {
	userId, err := GetUserIDFromCtx(ctx)
	if err != nil {
		return err
	}

	if userId != creatorID {
		logger.Errorf(
			ctx,
			"ValidateIsOwner, userID: %v, creatorID: %v",
			userId,
			creatorID,
		)
		return errors.Forbidden
	}

	return nil
}
