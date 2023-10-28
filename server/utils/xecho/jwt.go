package xecho

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"qiniu-1024-server/utils/shared"
	"qiniu-1024-server/utils/xerr"
)

func ParseUserJWT(c echo.Context) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*shared.JWTCustomClaims)
	c.Set(shared.JWTClaimUid, claims.UserID)
}
func CurUserID(c echo.Context) (int64, error) {
	uid := c.Get(shared.JWTClaimUid)
	if uid == nil {
		return 0, xerr.New(http.StatusUnauthorized, "NeedLogin", "need user id")
	}
	return uid.(int64), nil
}
