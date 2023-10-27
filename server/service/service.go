package service

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xmongo"
)

type Service struct {
	Logger *zap.Logger
	Mongo  *xmongo.Database
	Rdb    *redis.Client
	Conf   *types.Config
}

func NewService(conf *types.Config, logger *zap.Logger) *Service {
	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// mongo
	if conf.Debug {
		xmongo.SetDebug()
	}
	db, err := xmongo.DB(conf.Mongo)
	if err != nil {
		logger.Sugar().Fatalw("connect to mongo db failed", "error", err,
			"uri", conf.Mongo.URI, "db", conf.Mongo.DB)
	}

	if err != nil {
		panic(err)
	}
	logger.Info("connect mongo success")

	return &Service{
		Logger: logger,
		Rdb:    rdb,
		Mongo:  db,
		Conf:   conf,
	}
}
