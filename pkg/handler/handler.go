package handler

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	route := gin.New()

	auth := route.Group("auth")
	{
		auth.POST("/sigh-up", h.sighUp)
		auth.POST("/sigh-in", h.sighIn)
	}

	api := route.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItems)
				items.GET("/", h.getAllItems)
				items.GET("/:items_id", h.getItemsById)
				items.PUT("/:items_id", h.updateItems)
				items.DELETE("/:items_id", h.deleteItems)
			}
		}
	}
	return route
}
