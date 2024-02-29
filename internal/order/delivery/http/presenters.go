package http

import (
	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"

	"time"
)

var (
	minLengthPassword = 6
	maxLengthPassword = 20
)

type createRequest struct {
	UserID         int    `json:"username"`
	CoinCode       string `json:"password"`
	OrderPackageID int    `json:"order_package_id"`
	Blance         int    `json:"balance"`
}

type orderResponse struct {
	CoinCode     string    `json:"coin_code"`
	Time         int       `json:"time"`
	Date         time.Time `json:"date"`
	SKU          int       `json:"sku"`
	IsInprogress bool      `json:"is_inprogress"`
	Status       bool      `json:"status"`
	Balance      int       `json:"balance"`
}

type orderListResponse struct {
	TotalCount int64           `json:"total_count"`
	TotalPages int             `json:"total_pages"`
	Page       int             `json:"page"`
	Size       int             `json:"size"`
	HasMore    bool            `json:"has_more"`
	Users      []orderResponse `json:"users"`
}

// ToModel convert request to model
func (req createRequest) toModel() (models.Order, error) {
	// validErrCollector := errors.NewValidationErrorsCollector()

	// if req.Password != req.RetypePassword {
	// 	validErrCollector.Add(errors.NewValidationError("retype password not match", errMsgPasswordNotMatch))
	// }

	// if validErrCollector.HasError() {
	// 	return models.Order{}, validErrCollector
	// }

	return models.Order{
		UserID:         req.UserID,
		Balance:        req.Blance,
		OrderPackageID: req.OrderPackageID,
		CoinCode:       req.CoinCode,
	}, nil
}

func toOrderResponse(model models.Order) orderResponse {
	return orderResponse{
		CoinCode:     model.CoinCode,
		Time:         model.Time,
		SKU:          model.SKU,
		Balance:      model.Balance,
		Date:         model.Date,
		Status:       model.Status,
		IsInprogress: model.IsInprogress,
	}
}

func toUserListResp(pq utils.PaginationQuery, orders []models.Order) orderListResponse {
	orderResponse := make([]orderResponse, len(orders))
	for idx := range orders {
		orderResponse[idx] = toOrderResponse(orders[idx])
	}

	return orderListResponse{
		TotalCount: pq.TotalCount,
		TotalPages: pq.GetTotalPages(),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    pq.GetHasMore(),
		Users:      orderResponse,
	}
}
