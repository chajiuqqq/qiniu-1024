package xlog

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"qiniu-1024-server/types"
)

var logger *zap.Logger

func New(env string) *zap.Logger {
	var err error
	if logger != nil {
		return logger
	}
	switch env {
	case types.DebugEnv:
	case types.DevEnv:
	case types.ProdEnv:
	default:
		config := zap.NewDevelopmentConfig()
		logger, err = config.Build()
	}

	if err != nil {
		panic(err)
	}
	return logger
}

func FieldsFromContext(c echo.Context) []zap.Field {
	req := c.Request()
	res := c.Response()

	fields := []zapcore.Field{
		zap.String("path", c.Path()),
		zap.String("remote_ip", c.RealIP()),
		zap.String("host", req.Host),
		zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
		zap.String("user_agent", req.UserAgent()),
	}

	if res.Committed {
		fields = append(fields,
			zap.Int("status", res.Status),
			zap.Int64("size", res.Size),
		)
	}

	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	fields = append(fields, zap.String("request_id", id))

	return fields
}
