package api

import (
	v1 "travelxona/api/v1"
	"travelxona/config"
	"travelxona/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cfg  *config.Config
	Strg storage.StorageI
}

func New(h *Handler) *gin.Engine {
	engine := gin.Default()

	handlerV1 := v1.New(&v1.HandleV1{
		Cfg:  h.Cfg,
		Strg: h.Strg,
	})

	apiV1 := engine.Group("/v1")
	// apiV1.GET("/user/:id", handlerV1.GetUserById)
	apiV1.POST("/category", handlerV1.CreateCategory)
	// apiV1.PUT("/user", handlerV1.UpdateUser)
	// apiV1.DELETE("/user/:id", handlerV1.DeleteUser)
	apiV1.GET("/categories", handlerV1.GetListCategories)

	apiV1.POST("/users/login", handlerV1.Login)

	return engine
}
