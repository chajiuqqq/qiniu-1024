package main

import (
	"context"
	"flag"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zeromicro/go-zero/core/conf"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"qiniu-1024-server/service"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/shared"
	"qiniu-1024-server/utils/xecho"
	"qiniu-1024-server/utils/xlog"
	"qiniu-1024-server/utils/xmongo"
	"syscall"
	"time"
)

var configPath = flag.String("f", "config.yaml", "config path")

func main() {
	flag.Parse()
	var config = new(types.Config)
	conf.MustLoad(*configPath, config)

	logger := xlog.New("")
	srv := service.NewService(config, logger)

	// Echo instance
	e := echo.New()
	if config.Debug {
		e.Debug = true
	}
	// Real IP
	_, ipV4, _ := net.ParseCIDR("0.0.0.0/0")
	_, ipV6, _ := net.ParseCIDR("0:0:0:0:0:0:0:0/0")
	e.IPExtractor = echo.ExtractIPFromXFFHeader(echo.TrustIPRange(ipV4), echo.TrustIPRange(ipV6))
	// Error handler
	e.HTTPErrorHandler = xecho.NewErrorHandler(logger)
	// Disable echo logs, error handler above will log the error
	e.Logger.SetOutput(io.Discard)

	// Auto recover
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORS())

	u := e.Group("/v1", echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(config.JWTSecret),
		SuccessHandler: xecho.ParseUserJWT,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(shared.JWTUserClaims)
		},
	}))
	u.GET("/user", func(c echo.Context) error {
		return c.String(200, "user")
	})

	pub := e.Group("/v1")
	pub.GET("/public", func(c echo.Context) error {
		return c.String(200, "public")
	})
	// routes

	// Start server
	go func() {
		log.Infof("API Start at %s", config.ListenAddr)
		err := e.Start(config.ListenAddr)
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("API Force Shutting down: %s", err)
			log.Fatal("Force shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// shutdown echo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := e.Shutdown(ctx)
	if err != nil {
		log.Errorf("Universal API graceful shutdown failed: %s", err)
		log.Fatal(err)
	}
	// release global mongo connection
	xmongo.Close(ctx)
	log.Info("API graceful shutdown")
}
