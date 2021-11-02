package userhttp

import (
	"arka/cmd/service"
	"arka/pkg/router"
	"time"
)

var timeout = 5 * time.Second

type userDashboard struct {
	service service.UserService
}

func NewUserDashboard(service service.UserService) *userDashboard {
	return &userDashboard{service}
}

func (dashboard *userDashboard) Register(router *router.Router) {
	router.GET("/", dashboard.ListUser)
	router.GET("/:id", dashboard.GetUserByID)
	router.POST("/", dashboard.InserUser)
	router.PUT("/:id", dashboard.UpdateUser)
	router.DELETE("/:id", dashboard.DeleteUser)
}
