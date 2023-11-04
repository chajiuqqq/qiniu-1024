package main

import (
	"context"
	"flag"
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

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zeromicro/go-zero/core/conf"
)

var configPath = flag.String("f", "../local.yaml", "config path")

func main() {
	flag.Parse()
	var config = new(types.Config)
	conf.MustLoad(*configPath, config)

	logger := xlog.New("")
	sugar := logger.Sugar()
	sugar.Info("API Starting...")

	srv := service.NewService(config, logger)
	actionSrv := service.NewDefaultActionService(srv)
	srv.SetActionService(actionSrv)

	h := NewHandler(srv)

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
	// validator
	e.Validator = &xecho.CustomValidator{Validator: validator.New()}
	// Auto recover
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORS())

	u := e.Group("/v1", echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(config.JWTSecret),
		SuccessHandler: xecho.ParseUserJWT,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(shared.JWTCustomClaims)
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

	// oss
	u.POST("/upload", h.UploadFile)
	pub.POST("/oss/task/callback", h.OssTaskCallback)

	// user
	pub.POST("/user/register", h.PostRegister)
	pub.POST("/user/login", h.PostLogin)
	pub.GET("/user/:id", h.GetUser)              //带上作品
	u.POST("/user/:id/action", h.PostUserAction) // 关注/取消关注/喜欢/取消喜欢某人
	u.GET("/current-user", h.GetCurUser)         // 关注/取消关注/喜欢/取消喜欢某人

	// category
	pub.GET("/categories", h.GetMainCategories)
	u.POST("/categories", h.PostMainCategories)
	u.PUT("/category/:id", h.PutMainCategory)

	// video
	pub.GET("/videos", h.GetMainVideos)
	pub.GET("/video/:id", h.GetVideo)
	u.POST("/video", h.PostMainVideo)

	// video action
	pub.POST("/action/play/video/:id", h.PostPlayVideo)
	u.POST("/action/like/video/:id", h.PostLikeVideo)
	u.POST("/action/collect/video/:id", h.PostCollectVideo)
	u.DELETE("/action/like/video/:id", h.DeleteLikeVideo)
	u.DELETE("/action/collect/video/:id", h.DeleteCollectVideo)
	u.POST("/action/comment/video/:id", h.PostCommentVideo)

	// Start server
	go func() {
		sugar.Infof("API Start at %s", config.ListenAddr)
		err := e.Start(config.ListenAddr)
		if err != nil && err != http.ErrServerClosed {
			sugar.Errorf("API Force Shutting down: %s", err)
			sugar.Fatal("Force shutting down the server")
		}
	}()
	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	sugar.Infof("API now Shutdown...")
	// shutdown echo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := e.Shutdown(ctx)
	if err != nil {
		sugar.Errorf("API graceful shutdown failed: %s", err)
		sugar.Fatal(err)
	}
	// release global mongo connection
	xmongo.Close(ctx)
	sugar.Info("API graceful shutdown")
}
