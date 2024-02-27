package http

import (
	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/utils"

	"time"
)

var (
	minLengthPassword = 6
	maxLengthPassword = 20
)

type userCreateRequest struct {
	UserName       string `json:"username"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	RetypePassword string `json:"retype_password"`
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

type userResponse struct {
	ID        int       `json:"id"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type userListResponse struct {
	TotalCount int64          `json:"total_count"`
	TotalPages int            `json:"total_pages"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	HasMore    bool           `json:"has_more"`
	Users      []userResponse `json:"users"`
}

func (req loginReq) validate() error {
	errsCollector := errors.NewValidationErrorsCollector()
	if req.Username == "" {
		errsCollector.Add(errors.NewValidationError("username", "username is required"))
	}

	if req.Password == "" {
		errsCollector.Add(errors.NewValidationError("password", "password is required"))
	}

	if errsCollector.HasError() {
		return errsCollector
	}

	return nil
}

// ToModel convert request to model
func (req userCreateRequest) toModel() (models.User, error) {
	validErrCollector := errors.NewValidationErrorsCollector()
	if req.Password != "" && !utils.ValidatePassword(req.Password) {
		validErrCollector.Add(errors.NewValidationError("password", errMsgInvalidPassword))
	}
	if len(req.Password) < minLengthPassword {
		validErrCollector.Add(errors.NewValidationError("password", errMsgPasswordMinLength))
	}
	if len(req.Password) > maxLengthPassword {
		validErrCollector.Add(errors.NewValidationError("password", errMsgPasswordMaxLength))
	}
	if req.UserName == "" {
		validErrCollector.Add(errors.NewValidationError("username", errMsgRequiredUserName))
	}
	if req.Password == "" {
		validErrCollector.Add(errors.NewValidationError("username", errMsgRequiredPassword))
	}
	if req.Phone == "" {
		validErrCollector.Add(errors.NewValidationError("username", errMsgInvalidPhone))
	}
	if req.Email == "" {
		validErrCollector.Add(errors.NewValidationError("username", errMsgRequiredEmail))
	}

	if req.RetypePassword == "" {
		validErrCollector.Add(errors.NewValidationError("retype-password", errMsgRequireRetypePassword))
	}

	if req.Password != req.RetypePassword {
		validErrCollector.Add(errors.NewValidationError("retype password not match", errMsgPasswordNotMatch))
	}

	if validErrCollector.HasError() {
		return models.User{}, validErrCollector
	}

	return models.User{
		UserName: req.UserName,
		Password: req.Password,
		Phone:    req.Phone,
		Email:    req.Email,
	}, nil
}

// ToModel convert request to model
func (req userRequest) toModel() (models.User, error) {
	validErrCollector := errors.NewValidationErrorsCollector()
	if req.Phone != "" && utils.Phone(req.Phone) != nil {
		validErrCollector.Add(errors.NewValidationError("phone", errMsgInvalidPhone))
	}
	if req.Password != "" && !utils.ValidatePassword(req.Password) {
		validErrCollector.Add(errors.NewValidationError("password", errMsgInvalidPassword))
	}
	if len(req.Password) < minLengthPassword {
		validErrCollector.Add(errors.NewValidationError("password", errMsgPasswordMinLength))
	}
	if len(req.Password) > maxLengthPassword {
		validErrCollector.Add(errors.NewValidationError("password", errMsgPasswordMaxLength))
	}

	if validErrCollector.HasError() {
		return models.User{}, validErrCollector
	}

	return models.User{
		Phone:    req.Phone,
		Password: req.Password,
		//Status:   models.UserStatus(req.Status),
	}, nil
}

func toUserResponse(model models.User) userResponse {
	//var userProfileResp []userProfileResponse
	//countUserProfile := 0
	//for _, s := range model.UserProfile {
	//	userProfileResp = append(userProfileResp, userProfileResponse{
	//		ID:        s.ID,
	//		Name:      s.FullName,
	//		RoleName:  s.Role.Name,
	//		Email:     s.Email,
	//		CreatedAt: s.CreatedAt,
	//	})
	//
	//	countUserProfile += 1
	//}

	return userResponse{
		ID:    model.ID,
		Phone: model.Phone,
		Email: model.Email,
		//Status:    model.Status.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toUserListResp(pq utils.PaginationQuery, users []models.User) userListResponse {
	userResponse := make([]userResponse, len(users))
	for idx := range users {
		userResponse[idx] = toUserResponse(users[idx])
	}

	return userListResponse{
		TotalCount: pq.TotalCount,
		TotalPages: pq.GetTotalPages(),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    pq.GetHasMore(),
		Users:      userResponse,
	}
}
