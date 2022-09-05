package main

import (
	"net/http"
	"sum/internal/api/controllers"
	"sum/internal/api/middlewares"
	"sum/internal/app"
	"sum/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

func main() {
	var conf config.App
	err := envconfig.Process("api", &conf)
	if err != nil {
		log.Panic(err.Error())
	}

	l := app.ParseLogLevel(conf.LogLevel)
	log.SetLevel(l)

	router := gin.Default()
	router.POST("/auth", controllers.Auth(conf.JWT.TokenTTLHours, conf.JWT.SecretKey))
	router.POST("/sum", middlewares.JwtAuthMiddleware(conf.JWT.SecretKey), controllers.Sum)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	app.GracefullyShutdown(srv)
}
