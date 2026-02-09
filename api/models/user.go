package models

import "github.com/google/uuid"

type CategoryModelResp struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// type UpdateCategoryReq struct {
// 	Id   int64  `json:"id"`
// 	Name string `json:"name"`
// }

type GetCategoriesListResp struct {
	Categories []*CategoryModelResp `json:"categories"`
	Count      int64                `json:"count"`
}

type Category struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
}
