package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/video-hosting-platform/subscription-service/internal/service"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Init() *echo.Echo {

	router := echo.New()
	h.initRoutes(router)

	return router
}

func (h *Handler) initRoutes(router *echo.Echo) {
	api := router.Group("/api")
	api.POST("/subscribe/:id", greetings)   // subscribe to user
	api.POST("/unsubscribe/:id", greetings) // unsubscribe to user

	api.GET("/subscribe", greetings)            // get list of subscriptions for this user
	api.GET("/subscribe_status/:id", greetings) // get status of subscription to user
	api.GET("/subscribe_count/:id", greetings)  // get count of subscribers of user
	api.GET("/subscribe_count", greetings)      // get count of subscribers for this user
}

func greetings(c echo.Context) error {
	return c.JSON(200, "hello")
}
