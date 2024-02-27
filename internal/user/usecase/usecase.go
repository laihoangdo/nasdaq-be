package usecase

import (
	"context"

	"nasdaqvfs/internal/models"
	user2 "nasdaqvfs/internal/user"
	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/utils"
)

var (
	hashPasswordFunc = utils.HashPassword
)

// Create a user
func (u userUC) Create(ctx context.Context, user models.User) error {

	//Hash password
	if user.Password != "" {
		hashedPassword, err := hashPasswordFunc(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	//update user
	err := u.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// Login admin user to the system and return credentials
func (u *userUC) Login(ctx context.Context, username string, password string) (models.User, error) {
	funcName := "userUC.Login"

	user, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		u.logger.Errorf(ctx, "%s.userRepo.GetByUsername username=%s error=%s", funcName, username, err)
		if err == errors.NotFound {
			return models.User{}, user2.ErrUserNotFound
		}
		return models.User{}, user2.ErrUserNotFound
	}

	//if len(admin.Roles) > 0 {
	//	permissions, err := u.permissionUC.ListAllAdminPermission(ctx, admin.ID, admin.GetRole().ID, admin.IsAdministrator(), nil)
	//	if err != nil {
	//		u.logger.Errorf(ctx, "%s.permissionUC.ListAllAdminPermission username=%s error=%s", funcName, username, err)
	//	}
	//	admin.Permissions = permissions
	//}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		u.logger.Errorf(ctx, "%s.utils.CheckPassword username=%s error=%s", funcName, username, err)
		return models.User{}, user2.ErrUserWrongPassword
	}

	//accessToken, expAccessToken, err := utils.GenerateJWTAdminUser(u.cfg, admin, utils.JWTTypeAccessToken)
	//if err != nil {
	//	u.logger.Errorf(ctx, "%s.utils.GenerateJWTAdminUser username=%s error=%s", funcName, username, err)
	//	return models.AdminUserCredentials{}, user.ErrAdminUserGenJWT
	//}
	//
	//refreshToken, expRefreshToken, err := utils.GenerateJWTAdminUser(u.cfg, admin, utils.JWTTypeRefreshToken)
	//if err != nil {
	//	u.logger.Errorf(ctx, "%s.utils.GenerateJWTAdminUser username=%s error=%s", funcName, username, err)
	//	return models.AdminUserCredentials{}, user.ErrAdminUserGenJWT
	//}

	return user, nil
}

// GetUserByID get user by id
func (u userUC) GetUserByID(ctx context.Context, userID int64) (models.User, error) {
	user, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUserByID update user by id
func (u userUC) UpdateUserByID(ctx context.Context, userID int64, user models.User) error {
	//check user exist by id
	_, err := u.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	//Hash password
	if user.Password != "" {
		hashedPassword, err := hashPasswordFunc(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	//update user
	err = u.userRepo.UpdateUserByID(ctx, userID, user)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers list user
func (u userUC) GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]models.User, error) {
	results, err := u.userRepo.GetUsers(ctx, pq)
	if err != nil {
		return []models.User{}, errors.InternalServerError
	}

	return results, err
}
