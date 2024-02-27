package middleware

import (
	"context"
	"net/http"

	"nasdaqvfs/internal/models"
	pkgErrors "nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/token"
	"nasdaqvfs/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	authorizationHeader = "Authorization"
	userUUIDHeader      = "x-usertest-uuid"
)

func (mw *MiddlewareManager) checkHeaderAndGetPayload(c *gin.Context) (token.Payload, *pkgErrors.Error) {
	tokenString := c.GetHeader(authorizationHeader)
	userUUID := c.GetHeader(userUUIDHeader)
	if tokenString == "" && userUUID == "" {
		return token.Payload{}, pkgErrors.NewUnauthorizedError(pkgErrors.Unauthorized)
	}

	ctx := c.Request.Context()
	mw.logger.Infof(ctx, "auth middleware header %s", tokenString)
	payload, err := mw.tokenMaker.VerifyToken(tokenString)
	if err != nil {
		mw.logger.Error(ctx, "middleware validateJWTToken", zap.String("headerJWT", err.Error()))
		return token.Payload{}, pkgErrors.NewUnauthorizedError(pkgErrors.Unauthorized)
	}

	//if err := mw.checkAuthBlackList(ctx, tokenString, payload); err != nil {
	//	return token.Payload{}, err
	//}

	return payload, nil
}

// JWT way of auth using cookie or Authorization header
func (mw *MiddlewareManager) AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := mw.checkHeaderAndGetPayload(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		ctx := c.Request.Context()

		ctx, pkErr := mw.setSessionToCtx(ctx, payload)
		if pkErr != nil {
			c.JSON(http.StatusUnauthorized, pkgErrors.NewUnauthorizedError(pkgErrors.InvalidJWTToken))
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, utils.UserUUIDCtxKey{}, payload.UserUUID)
		req := c.Request.WithContext(ctx)
		c.Request = req
		c.Next()
	}
}

func (mw *MiddlewareManager) AuthProfileJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := mw.checkHeaderAndGetPayload(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx, pkErr := mw.setSessionToCtx(ctx, payload)
		if pkErr != nil {
			c.JSON(http.StatusUnauthorized, pkgErrors.NewUnauthorizedError(pkgErrors.InvalidJWTToken))
			c.Abort()
			return
		}

		ctx, pkErr = mw.checkProfile(ctx, payload)
		if pkErr != nil {
			c.JSON(http.StatusUnauthorized, pkgErrors.NewUnauthorizedError(pkgErrors.InvalidJWTToken))
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, utils.UserUUIDCtxKey{}, payload.UserUUID)

		req := c.Request.WithContext(ctx)
		c.Request = req
		c.Next()
	}
}

func (mw *MiddlewareManager) setSessionToCtx(ctx context.Context, payload token.Payload) (context.Context, *pkgErrors.Error) {
	session := models.Session{
		UUID:      payload.ID,
		UserUUID:  payload.UserUUID,
		ExpiredAt: payload.ExpiredAt,
		CreatedAt: payload.IssueAt,
	}
	if !session.IsValid() {
		return ctx, pkgErrors.NewUnauthorizedError(pkgErrors.InvalidJWTToken)
	}
	ctx = utils.SetSessionToCtx(ctx, session)

	return ctx, nil
}

func (mw *MiddlewareManager) checkProfile(ctx context.Context, payload token.Payload) (context.Context, *pkgErrors.Error) {
	if payload.Profile.UUID == "" || payload.Profile.RoleSlug == "" {
		return ctx, pkgErrors.NewUnauthorizedError(pkgErrors.InvalidJWTToken)
	}

	profile := models.User{
		//UUID:     payload.Profile.UUID,
		//RoleSlug: models.RoleSlug(payload.Profile.RoleSlug),
	}
	ctx = utils.SetProfileToCtx(ctx, profile)
	return ctx, nil
}
