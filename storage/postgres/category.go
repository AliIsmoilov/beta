package postgres

import (
	"context"

	"travelxona/storage/repo"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) repo.CategoryI {
	return &categoryRepo{
		db: db,
	}
}

func (u *categoryRepo) Create(ctx context.Context, category repo.Category) (*repo.Category, error) {

	if err := u.db.WithContext(ctx).
		Table("category").
		Create(&category).
		Error; err != nil {
		return nil, err
	}

	return &category, nil
}

// // Update user info
// func (u *userRepo) Update(ctx context.Context, req *repo.UpdateUserReq) (*repo.UserModelResp, error) {
// 	var user repo.UserModelResp

// 	if err := u.db.WithContext(ctx).
// 		Table("users").
// 		First(&user, req.Id).Error; err != nil {
// 		return nil, err
// 	}

// 	user.FirstName = req.FirstName
// 	user.LastName = req.LastName
// 	user.PhoneNumber = *req.PhoneNumber
// 	user.Email = req.Email

// 	if err := u.db.WithContext(ctx).
// 		Table("users").
// 		Save(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (u *userRepo) UpdateUserPhoto(ctx context.Context, req *repo.UpdateUserPhotoReq) error {
// 	var user repo.UserModelResp
// 	if err := u.db.WithContext(ctx).
// 		Table("users").
// 		First(&user, req.Id).
// 		Error; err != nil {
// 		return err
// 	}

// 	user.FirstName = req.ProfilePhoto
// 	if err := u.db.WithContext(ctx).
// 		Table("users").
// 		Where("id = ?", req.Id).
// 		Update("profile_photo", req.ProfilePhoto).
// 		Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *userRepo) GetById(ctx context.Context, id uuid.UUID) (*repo.UserModelResp, error) {
// 	var user repo.UserModelResp
// 	if err := u.db.WithContext(ctx).
// 		Table("users").
// 		Preload("DriverProfile").
// 		// Preload("RiderProfile").
// 		First(&user, "id = ?", id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

func (u *categoryRepo) GetListCategories(ctx context.Context, req repo.GetAllCategoriesReq) (repo.GetAllCategoriesResp, error) {
	var categories []repo.Category

	req.Query = "%" + req.Query + "%"
	var count int64

	tx := u.db.WithContext(ctx).
		Table("category").
		Where("deleted_at IS NULL").
		Where("name ILIKE ?", "%"+req.Query+"%")

	err := tx.Count(&count).Error
	if err != nil {
		return repo.GetAllCategoriesResp{}, err
	}

	if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		tx = tx.Offset(int(offset)).
			Limit(int(req.Limit))
	} else if req.Limit > 0 {
		tx = tx.Limit(int(req.Limit))
	}

	if err := tx.Find(&categories).Error; err != nil {
		return repo.GetAllCategoriesResp{}, err
	}

	return repo.GetAllCategoriesResp{
		Categories: categories,
		Count:      count,
	}, nil
}
