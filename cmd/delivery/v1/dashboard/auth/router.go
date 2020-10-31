package dashboard

import (
	"arka/cmd/service"
	"arka/pkg/router"
	"time"
)

var timeout = 5 * time.Second

type authDashboard struct {
	service service.AuthService
}

func NewAuthDashboard(service service.AuthService) *authDashboard {
	return &authDashboard{service}
}

func (dashboard *authDashboard) Register(router *router.Router) {
	// router.Use(middleware.RequiresAccessToken)
	router.GET("/login", dashboard.Login)
}
