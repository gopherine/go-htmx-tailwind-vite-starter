package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang_starter/internal/assets"
	"github.com/golang_starter/internal/logger"
	"github.com/golang_starter/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	_ = godotenv.Load()

	// Initialize logger
	logger.Init(logger.Config{
		Level:  getEnv("LOG_LEVEL", "info"),
		Pretty: getEnv("LOG_PRETTY", "true") == "true",
	})

	ctx := logger.WithTrace(context.Background(), "SERVER")
	logger.Info(ctx, "Starting server...")

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	assets.SetFS(PublicFS)

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(loggingMiddleware())

	// Routes
	e.GET("/public/*", echo.WrapHandler(http.StripPrefix("/public/", http.FileServer(http.FS(PublicFS)))))
	routes.RegisterAll(e)

	// Start server
	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		port := getEnv("LISTEN_ADDR", "1323")
		logger.Info(ctx, "Server listening on :"+port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			logger.Fatal(ctx, "Server failed: "+err.Error())
		}
	}()

	// Wait for interrupt
	<-signalCtx.Done()
	logger.Info(ctx, "Shutting down...")

	shutdownTimeout := getShutdownTimeout()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		logger.Error(ctx, "Shutdown error: "+err.Error())
	}

	logger.Info(ctx, "Server stopped")
}

func loggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := logger.WithTrace(c.Request().Context(), "HTTP")
			ctx = context.WithValue(ctx, c, c) // Store echo context for later

			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path

			fields := logger.F{
				"method":  method,
				"path":    path,
				"status":  status,
				"latency": latency.String(),
			}

			if status >= 500 {
				logger.Log(ctx, "error", "Request failed", fields)
			} else if status >= 400 {
				logger.Log(ctx, "warn", "Request error", fields)
			} else {
				logger.Log(ctx, "debug", "Request", fields)
			}

			return err
		}
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getShutdownTimeout() time.Duration {
	if val := getEnv("SHUTDOWN_TIMEOUT", ""); val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			return duration
		}
	}
	return 10 * time.Second
}
