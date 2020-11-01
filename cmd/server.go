package cmd

import (
	"arka/cmd/config"
	dashboard "arka/cmd/delivery/http/v1/dashboard/auth"
	"arka/cmd/repositories"
	"arka/cmd/repositories/user/mysql"
	"arka/cmd/repositories/user/redis"
	service "arka/cmd/service/auth"
	"arka/pkg/auth"
	"arka/pkg/cache"
	"arka/pkg/casbin"
	"arka/pkg/database"
	"arka/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

func Run() {
	formatter := runtime.Formatter{ChildFormatter: &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "01/02/06 15:04:05",
	}}
	formatter.Line = true
	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.StandardLogger()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := config.LoadConfig()
	logrus.Info("Config Loaded")

	// Database
	db, err := database.InitDB(config.DBConfig)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info("Success Initialize Database")

	// Redis
	redisConn := database.InitRedis(config.RedisConfig)
	defer redisConn.Close()

	redisCommand := cache.Init(redisConn)
	logrus.Info("Success Initialize Redis")

	// Auth
	authModule := auth.New(redisCommand)
	tokenModule := auth.NewToken()

	// Casbin
	casbin := casbin.New(db)
	casbin.Register()

	// Repository
	var authRepository repositories.UserRepository = mysql.NewUser(db)
	authRepository = redis.NewUser(redisCommand, authRepository, config.RedisConfig.Prefix)

	// Service
	var authService = service.NewAuthService(authRepository, authModule, tokenModule)

	// Server
	server := server.New(&config.ServerConfig)

	// router
	router := server.Router()
	v1 := router.Group("/v1")
	dashboardRouter := v1.Group("/dashboard")
	dashboard.NewAuthDashboard(authService).Register(dashboardRouter)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serveChan := server.Run()

	select {
	case err := <-serveChan:
		if err != nil {
			panic(err)
		}
	case <-sigChan:
	}

	server.Stop()
}
