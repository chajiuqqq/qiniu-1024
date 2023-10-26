package shared

import "github.com/golang-jwt/jwt/v5"

const (
	JWTClaimUid = "uid"
)

const (
	JWTAudienceUser  = "user"
	JWTAudienceAdmin = "admin"
)

// JWTUserClaims is custom jwt claims
type JWTUserClaims struct {
	UserID    string `json:"uid"`
	ExpiredAt int64  `json:"exp"`
	Audience  string `json:"aud"`
	jwt.RegisteredClaims
}

// JWTAdminClaims is custom jwt claims
type JWTAdminClaims struct {
	UserID    string `json:"uid"`
	ExpiredAt int64  `json:"exp"`
	Audience  string `json:"aud"`
	jwt.RegisteredClaims
}
