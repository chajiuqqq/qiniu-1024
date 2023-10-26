package types

import (
	"qiniu-1024-server/utils/oss"
	"qiniu-1024-server/utils/xmongo"
)

const (
	DevEnv   = "Dev"
	DebugEnv = "Debug"
	ProdEnv  = "Prod"
)

type Config struct {
	ListenAddr string `json:",default=0.0.0.0:9133"`
	Debug      bool
	Redis      struct {
		Addr string
	}
	Mongo     xmongo.Config
	JWTSecret string
	Oss       oss.Config
}
