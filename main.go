package main

import (
	"context"
	"flag"
	stdlog "log"
	"os"
	"os/signal"
	"time"

	"github.com/checkaayush/gofrog/config"
	"github.com/checkaayush/gofrog/internal"
	"github.com/checkaayush/gofrog/pkg/artifactory"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func getEnvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}

	return val
}

func main() {
	confFile := flag.String("config", "config.toml", "Path to config file")
	flag.Parse()

	// Load service configuration
	cfg, err := config.Load(*confFile)
	if err != nil {
		stdlog.Fatalf("config.Load(%s) failed: %s", *confFile, err.Error())
	}
	stdlog.Printf("loaded config from %s", *confFile)

	// Initialize artifactory client
	rtClient, err := artifactory.NewClient(cfg.Artifactory)
	if err != nil {
		stdlog.Fatalf("artifactory.NewClient() failed: %s", err.Error())
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Add middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add Basic Auth
	e.Use(middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
		authUser := getEnvWithDefault("AUTH_USERNAME", "admin")
		authPass := getEnvWithDefault("AUTH_PASSWORD", "admin")
		if user == authUser && pass == authPass {
			return true, nil
		}
		return false, nil
	}))

	// Initialize handler
	h := internal.NewHandler(rtClient)

	// Add routes
	v1 := e.Group("/v1")
	v1.GET("/health", h.Health)
	v1.GET("/artifacts/mostPopular", h.GetMostPopularArtifacts)

	// Start server
	addr := ":" + getEnvWithDefault("PORT", "5000")
	go func() {
		if err := e.Start(addr); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
