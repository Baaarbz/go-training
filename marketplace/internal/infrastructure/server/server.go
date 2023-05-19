package server

import (
	adHandler "barbz.dev/marketplace/internal/infrastructure/server/handler/ad"
	"barbz.dev/marketplace/internal/infrastructure/server/handler/health"
	"barbz.dev/marketplace/internal/infrastructure/server/middleware/logging"
	"barbz.dev/marketplace/internal/infrastructure/server/middleware/recovery"
	adService "barbz.dev/marketplace/internal/pkg/application/ad"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	shutdownTimeout time.Duration
	// Dependencies
	services map[string]interface{}
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, services map[string]interface{}) (context.Context, Server) {
	engine := gin.New()
	// Register middlewares
	engine.Use(recovery.Middleware(), logging.Middleware())

	srv := Server{
		engine:          engine,
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		services:        services,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.APIStatus())

	s.engine.POST("/v1/ads", adHandler.SaveAd(s.services["ad.SaveAd"].(adService.SaveAd)))
	s.engine.GET("/v1/ads", adHandler.FindAllAds(s.services["ad.FindAllAds"].(adService.FindAllAds)))
	s.engine.GET("/v1/ads/:id", adHandler.FindAdById(s.services["ad.FindAdById"].(adService.FindAdById)))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
