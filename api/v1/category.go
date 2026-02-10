package v1

import (
	"net/http"
	"strconv"

	"travelxona/api/models"
	"travelxona/storage/repo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreateCategory(ctx *gin.Context) {
	var req models.Category
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "error while binding data",
		})
		return
	}

	data, err := h.strg.Category().Create(ctx, repo.Category{
		Id:   uuid.New(),
		Name: req.Name,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error while creating user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, parseCategoryRepoToApi(data))
}

// func (h *handlerV1) UpdateUser(ctx *gin.Context) {
// 	var req models.UpdateUserReq
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   err.Error(),
// 			"message": "error while binding data",
// 		})
// 		return
// 	}

// 	data, err := h.strg.User().Update(ctx, &repo.UpdateUserReq{
// 		Id:          req.Id,
// 		PhoneNumber: req.PhoneNumber,
// 		FullName:    req.FullName,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"message": "error while updating user",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, parseUserRepoToApi(data))
// }

// func (h *handlerV1) GetUserById(ctx *gin.Context) {
// 	id := ctx.Param("id")

// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"message": "strconv atoi not worked",
// 		})
// 		return
// 	}
// 	data, err := h.strg.User().GetById(ctx, int64(idInt))
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error":   err.Error(),
// 				"message": "user not found",
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"message": "user get by id not worked",
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, parseUserRepoToApi(data))
// }

// func (h *handlerV1) DeleteUser(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	idInt, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"message": "strconv atoi not worked",
// 		})
// 		return
// 	}

// 	err = h.strg.User().Delete(ctx, idInt)
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			ctx.JSON(http.StatusNotFound, gin.H{
// 				"error":   err.Error(),
// 				"message": "user not found",
// 			})
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   err.Error(),
// 			"message": "user get by id not worked",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"ok": true,
// 	})
// }

func (h *handlerV1) GetListCategories(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "")
	offset := ctx.DefaultQuery("page", "")
	query := ctx.DefaultQuery("query", "")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "strconv atoi not worked",
		})
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "strconv atoi not worked",
		})
		return
	}

	data, err := h.strg.Category().GetListCategories(ctx, repo.GetAllCategoriesReq{
		Limit: int32(limitInt),
		Page:  int32(offsetInt),
		Query: query,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error while getting all users",
		})
		return
	}

	resp := models.GetCategoriesListResp{
		Count: data.Count,
	}

	for _, elem := range data.Categories {
		category := parseCategoryRepoToApi(&elem)
		resp.Categories = append(resp.Categories, &category)
	}

	ctx.JSON(http.StatusOK, resp)
}

func parseCategoryRepoToApi(c *repo.Category) models.CategoryModelResp {
	return models.CategoryModelResp{
		Id:   c.Id,
		Name: c.Name,
	}
}
