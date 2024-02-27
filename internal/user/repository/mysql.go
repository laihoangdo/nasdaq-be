package repository

import (
	"context"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/database/mysql"
	"nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/utils"

	"gorm.io/gorm"
)

// Create a user
func (r userRepo) Create(ctx context.Context, user models.User) error {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// GetByUsername gets  user info using username
func (r *userRepo) GetByUsername(ctx context.Context, username string) (record models.User, err error) {
	if err = r.db.WithContext(ctx).
		//Preload("Roles").
		First(&record, "username = ?", username).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, errors.NotFound
		}
		return models.User{}, err
	}
	return record, nil
}

// GetUserByID get user by id
func (r userRepo) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	var user models.User
	err := r.db.Model(&models.User{}).Where("id = ?", id).Preload("UserProfile.Role").First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// UpdateUserByID update user by ID
func (r userRepo) UpdateUserByID(ctx context.Context, id int64, user models.User) error {
	err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// Get User
func (r userRepo) GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]models.User, error) {
	var err error
	users := []models.User{}

	db := r.db.WithContext(ctx)

	// Filter
	//filter, ok := pq.Filter.(models.UserFilter)
	//if !ok {
	//	return []models.User{}, err
	//}
	//
	//if filter.ID != "" {
	//	db = db.Where("id = ?", filter.ID)
	//}
	//
	//if filter.Phone != "" {
	//	db = db.Where("phone = ?", filter.Phone)
	//}
	//
	//if filter.Status != "" {
	//	db = db.Where("status = ?", filter.Status)
	//}
	//
	//if filter.FromDate != "" && filter.ToDate != "" {
	//	db = db.Scopes(mysql.FilterByCreationDate(filter.FromDate, filter.ToDate))
	//}

	// Count total items
	if _, err := mysql.CountTotal(db, users, pq); err != nil {
		return []models.User{}, err
	}

	// Pagination
	if db, err = mysql.Paginate(db, users, pq); err != nil {
		return []models.User{}, err
	}

	// Query
	if err = db.Find(&users).Error; err != nil {
		return []models.User{}, err
	}

	return users, nil
}
