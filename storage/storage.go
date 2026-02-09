package storage

import (
	"github.com/SaidovZohid/deposit-project/storage/postgres"
	"github.com/SaidovZohid/deposit-project/storage/repo"
	"gorm.io/gorm"
)

type StorageI interface {
	Category() repo.CategoryI
}

type storage struct {
	categoryRepo repo.CategoryI
}

func New(db *gorm.DB) StorageI {
	return &storage{
		categoryRepo: postgres.NewCategory(db),
	}
}

func (s *storage) Category() repo.CategoryI {
	return s.categoryRepo
}
