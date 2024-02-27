package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/logger"
	"nasdaqvfs/pkg/sanitize"
)

// Get request id from gin context
func GetRequestID(c *gin.Context) string {
	return requestid.Get(c)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from echo context
func GetCtxWithReqID(c *gin.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// Get context  with request id
func GetRequestCtx(c *gin.Context) context.Context {
	return context.WithValue(c.Request.Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// UserIDCtxKey is a key used for the User object in the context
type UserIDCtxKey struct{}

// UserUUIDCtxKey is a uuid key used for the User object in the context
type UserUUIDCtxKey struct{}

type ProfileCtxKey struct{}

// Session
type SessionCtxKey struct{}

// Get usertest ip address
func GetIPAddress(c *gin.Context) string {
	return c.ClientIP()
}

// Get usertest id from context
func GetUserIDFromCtx(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(UserIDCtxKey{}).(string)
	if !ok {
		return "", errors.Unauthorized
	}

	return userID, nil
}

// Get usertest uuid from context
func GetUserUUIDFromCtx(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(UserUUIDCtxKey{}).(string)
	if !ok {
		return "", errors.Unauthorized
	}

	return userID, nil
}

func GetSessionFromCtx(ctx context.Context) (models.Session, error) {
	session, ok := ctx.Value(SessionCtxKey{}).(models.Session)
	if !ok {
		return models.Session{}, errors.Unauthorized
	}

	return session, nil
}

func SetSessionToCtx(ctx context.Context, session models.Session) context.Context {
	return context.WithValue(ctx, SessionCtxKey{}, session)
}

func SetProfileToCtx(ctx context.Context, profile models.User) context.Context {
	return context.WithValue(ctx, ProfileCtxKey{}, profile)
}

func GetProfileFromCtx(ctx context.Context) (models.User, error) {
	profile, ok := ctx.Value(ProfileCtxKey{}).(models.User)
	if !ok {
		return models.User{}, errors.Unauthorized
	}
	return profile, nil
}

// Error response with logging error for echo context
func LogResponseError(c *gin.Context, logger logger.Logger, err error) {
	logger.Errorf(
		c.Request.Context(),
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(c),
		GetIPAddress(c),
		err,
	)
}

// Read request body and validate
func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}

// Read sanitize and validate request
func SanitizeRequest(ctx *gin.Context, request interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	defer ctx.Request.Body.Close()

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request.Context(), request)
}
