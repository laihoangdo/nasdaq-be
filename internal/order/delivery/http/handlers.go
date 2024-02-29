package http

import (
	"strconv"

	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/response"
	"nasdaqvfs/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Create a order
func (h orderHandlers) Create(c *gin.Context) {
	ctx := c.Request.Context()

	n := createRequest{}
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

	err = h.orderUC.Create(ctx, input)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, nil)
}

// UpdateUserByID update user
func (h orderHandlers) UpdateByID(c *gin.Context) {
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

	err = h.orderUC.UpdateByID(ctx, id, input)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, nil)
}

// GetUserByID get user by id
func (h orderHandlers) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewBadRequestError(errors.WithMessage(err, "Invalid request ID")))
		return
	}

	user, err := h.orderUC.GetByID(ctx, id)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, err)
		return
	}

	response.WithOK(c, toUserResponse(user))
}

// List users
func (h orderHandlers) GetOrders(c *gin.Context) {
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

	userList, err := h.orderUC.GetOrders(ctx, &pq)
	if err != nil {
		utils.LogResponseError(c, h.logger, err)
		response.WithError(c, errors.NewInternalServerError(err))
		return
	}

	response.WithOK(c, toUserListResp(pq, userList))
}
