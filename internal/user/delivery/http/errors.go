package http

import (
	"net/http"

	"nasdaqvfs/internal/user"
	"nasdaqvfs/pkg/response"
)

var mapError = response.MapResponse{
	// Other
	user.ErrUserNotFound:      response.Response{Code: 404, StatusCode: http.StatusOK, Message: "user not found"},
	user.ErrUserWrongPassword: response.Response{Code: 400, StatusCode: http.StatusOK, Message: "wrong password"},
}

const (
	errMsgInvalidPhone          = "Invalid phone"
	errMsgInvalidStatus         = "Invalid status"
	errMsgInvalidPassword       = "Invalid password"
	errMsgRequiredPassword      = "password is required"
	errMsgRequiredUserName      = "username is required"
	errMsgRequiredPhone         = "phone is required"
	errMsgRequiredEmail         = "email is required"
	errMsgRequireRetypePassword = "re-type pasword is required"
	errMsgInvalidUserName       = "Invalid username"
	errMsgInvalidReTypePassword = "Invalid retype password"
	errMsgPasswordNotMatch      = "password not match"
	errMsgInvalidEmail          = "Invalid email"
	errMsgPasswordMinLength     = "Password min length is 6"
	errMsgPasswordMaxLength     = "Password max length is 20"
)
