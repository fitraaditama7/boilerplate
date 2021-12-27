package userhttp

import (
	"arka/cmd/middleware"
	"arka/cmd/service"
	"arka/pkg/router"
	"time"
)

var timeout = 5 * time.Second

type userDashboard struct {
	service    service.UserService
	middleware middleware.AuthMiddleware
}

func NewUserDashboard(service service.UserService, middleware middleware.AuthMiddleware) *userDashboard {
	return &userDashboard{service, middleware}
}

func (dashboard *userDashboard) Register(router *router.Router) {
	router.Use(dashboard.middleware.RequiresAccessToken)
	userRouter := router.Group("/user")
	userRouter.GET("/", middleware.RequiresAuthorization(dashboard.ListUser))
	userRouter.GET("/:id", dashboard.GetUserByID)
	userRouter.POST("/", middleware.RequiresAuthorization(dashboard.InserUser))
	userRouter.PUT("/:id", dashboard.UpdateUser)
	userRouter.DELETE("/:id", middleware.RequiresAuthorization(dashboard.DeleteUser))
}
