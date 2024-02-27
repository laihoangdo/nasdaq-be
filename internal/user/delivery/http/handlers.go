package http

import (
	"net/http"
	"strconv"

	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/response"
	"nasdaqvfs/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Create a user
func (h userHandlers) Create(c *gin.Context) {
	ctx := c.Request.Context()

	n := userCreateRequest{}
	if err := c.Bind(&n); err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(errors.WithMessage(err, "Invalid request body")))
		return
	}

	input, err := n.toModel()
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	err = h.userUC.Create(ctx, input)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, nil)
}

// Login  admin user to the system
func (h *userHandlers) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewError(http.StatusBadRequest, "request body is invalid", nil))
		return
	}

	if err := req.validate(); err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	resp, err := h.userUC.Login(ctx, req.Username, req.Password)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithErrorAndMap(c, err, mapError)
		return
	}

	response.WithOK(c, toUserResponse(resp))
}

// UpdateUserByID update user
func (h userHandlers) UpdateUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(errors.WithMessage(err, "Invalid request ID")))
		return
	}

	n := userRequest{}
	if err = c.Bind(&n); err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(errors.WithMessage(err, "Invalid request body")))
		return
	}

	input, err := n.toModel()
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	err = h.userUC.UpdateUserByID(ctx, id, input)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, nil)
}

// GetUserByID get user by id
func (h userHandlers) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(errors.WithMessage(err, "Invalid request ID")))
		return
	}

	user, err := h.userUC.GetUserByID(ctx, id)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, toUserResponse(user))
}

// List users
func (h userHandlers) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	pq, err := utils.GetPaginationFromCtx(c)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(err))
		return
	}

	//filter := models.UserFilter{
	//	ID:       strings.TrimSpace(c.Query("id")),
	//	Phone:    strings.TrimSpace(c.Query("phone")),
	//	Status:   strings.TrimSpace(c.Query("status")),
	//	FromDate: strings.TrimSpace(c.Query("from_date")),
	//	ToDate:   strings.TrimSpace(c.Query("to_date")),
	//}
	//pq.SetFilter(filter)

	userList, err := h.userUC.GetUsers(ctx, &pq)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewInternalServerError(err))
		return
	}

	response.WithOK(c, toUserListResp(pq, userList))
}
