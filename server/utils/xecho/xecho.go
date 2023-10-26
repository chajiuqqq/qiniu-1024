package xecho

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"net/http"
	"qiniu-1024-server/utils/xerr"
	"qiniu-1024-server/utils/xlog"
	"strings"
)

// NewErrorHandler return a customize echo's HTTP error handler.
func NewErrorHandler(logger *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		// the final response body
		var res *xerr.XError

		// check err
		if he, ok := err.(*xerr.XError); ok {
			// custom error by this package
			res = he
		} else if ve, ok := err.(validator.ValidationErrors); ok {
			// validator errors
			res = xerr.New(400, "BadRequest", ve.Error())
		} else if ee, ok := err.(*echo.HTTPError); ok {
			// echo errors
			res = xerr.New(ee.Code, strings.ReplaceAll(http.StatusText(ee.Code), " ", ""), ee.Error())
		} else if errors.Is(err, mongo.ErrNoDocuments) {
			// gorm not found
			res = xerr.New(404, "NotFound", "record not found")
		} else if errors.Is(err, context.Canceled) {
			res = xerr.New(400, "ClientCanceled", "client has closed the request")
		} else {
			// other server errors
			res = xerr.New(500, "ServerError", err.Error())
		}

		// log server error, and hide the message to client
		if res.StatusCode == 500 {
			logger.Error(res.Message, xlog.FieldsFromContext(c)...)
			if !c.Echo().Debug {
				res = xerr.ServerError
			}
		} else if res.StatusCode >= 400 {
			logger.Info(res.Message, xlog.FieldsFromContext(c)...)
		}

		// echo need this
		if !c.Response().Committed {
			if c.Request().Method == echo.HEAD {
				err = c.NoContent(res.StatusCode)
			} else {
				err = c.JSON(res.StatusCode, res)
			}
			if err != nil {
				// log the resp sent error
				logger.Error(err.Error())
			}
		}
	}
}
