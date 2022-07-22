package main

import (
	"context"
	"fmt"
	"github.com/alekseyklimenko/go-proj-bootstrap/config"
	"github.com/alekseyklimenko/go-proj-bootstrap/controllers"
	"github.com/alekseyklimenko/go-proj-bootstrap/database"
	"github.com/alekseyklimenko/go-proj-bootstrap/logger"
	"github.com/alekseyklimenko/go-proj-bootstrap/services/initservices"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf := config.New()
	db := database.New(conf)
	database.RunMigrations(db)
	initservices.Init(db, conf)
	srv := initWebServer(conf)
	deferShutdown(srv)
}

func initWebServer(conf *config.Config) *http.Server {
	router := gin.Default()
	controllers.RegisterHandlers(router)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Web.Port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.NewEntry().WithError(err).Error("ListenAndServe error")
		}
	}()
	return srv
}

func deferShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.NewEntry().Info("Shutting down server...")
	initservices.Shutdown()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.NewEntry().WithError(err).Error("Server forced to shutdown")
	}
	logger.NewEntry().Info("Server exiting")
}
