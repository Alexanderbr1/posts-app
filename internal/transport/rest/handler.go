package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"posts-app/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.siqnUp)
		auth.POST("/sign-in", h.siqnIn)
		auth.GET("/refresh", h.refresh)
	}

	api := router.Group("/api/v1", h.authMiddleware)
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.create)
			posts.GET("/", h.getAll)
			posts.GET("/:id", h.getByID)
			posts.PUT("/:id", h.update)
			posts.DELETE("/:id", h.delete)
		}
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
