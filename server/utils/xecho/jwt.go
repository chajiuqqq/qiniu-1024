package xecho

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"qiniu-1024-server/utils/shared"
)

func ParseUserJWT(c echo.Context) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*shared.JWTCustomClaims)
	c.Set(shared.JWTClaimUid, claims.UserID)
}
