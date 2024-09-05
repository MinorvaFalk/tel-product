package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"tel/product/config"
	"tel/product/internal/api"
	"tel/product/pkg/exception"
	"tel/product/pkg/validation"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	e *echo.Echo
	h api.Handler
}

func New(h api.Handler) *Router {
	e := echo.New()
	e.Validator = validation.NewValidator()
	e.HTTPErrorHandler = exception.EchoErrorHandler

	e.Use(middleware.Recover())

	return &Router{
		e: e,
		h: h,
	}
}

func (r *Router) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := r.e.Start(":" + config.ReadConfig().Port); err != nil && err != http.ErrServerClosed {
			r.e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.e.Shutdown(ctx); err != nil {
		r.e.Logger.Fatal(err)
	}
}
