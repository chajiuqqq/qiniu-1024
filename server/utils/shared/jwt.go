package shared

import "github.com/golang-jwt/jwt/v5"

const (
	JWTClaimUid = "uid"
)

const (
	JWTAudienceUser  = "user"
	JWTAudienceAdmin = "admin"
)

// JWTCustomClaims is custom jwt claims
type JWTCustomClaims struct {
	UserID   int64  `json:"uid"`
	Audience string `json:"aud"`
	jwt.RegisteredClaims
}
