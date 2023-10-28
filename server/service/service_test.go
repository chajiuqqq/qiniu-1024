package service

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
	"qiniu-1024-server/model"
	"qiniu-1024-server/types"
	"qiniu-1024-server/utils/xlog"
	"strings"
	"testing"
)

var srv *Service
var ctx = context.Background()

func TestMain(m *testing.M) {
	configPath := "../cmd/ut.yaml"
	var config = new(types.Config)
	conf.MustLoad(configPath, config)

	logger := xlog.New("")
	srv = NewService(config, logger)

	if !strings.HasPrefix(srv.Conf.Mongo.DB, "ut-") {
		panic("mongo db name must start with ut-")
	}

	colls := []string{
		model.User{}.Collection(),
		model.Video{}.Collection(),
		model.Counter{}.Collection(),
		model.CommentLog{}.Collection(),
		model.VideoLog{}.Collection(),
	}
	for _, coll := range colls {
		err := srv.Mongo.Collection(coll).Drop(ctx)
		if err != nil {
			panic(err)
		}
	}

	exitVal := m.Run()
	os.Exit(exitVal)
}
