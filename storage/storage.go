package storage

import (
	"travelxona/storage/postgres"
	"travelxona/storage/repo"

	"gorm.io/gorm"
)

type StorageI interface {
	Category() repo.CategoryI
	User() repo.UserI
}

type storage struct {
	categoryRepo repo.CategoryI
	userRepo     repo.UserI
}

func New(db *gorm.DB) StorageI {
	return &storage{
		categoryRepo: postgres.NewCategory(db),
		userRepo:     postgres.NewUser(db),
	}
}

func (s *storage) Category() repo.CategoryI {
	return s.categoryRepo
}

func (s *storage) User() repo.UserI {
	return s.userRepo
}
