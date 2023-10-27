package xmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"qiniu-1024-server/utils/xlog"
	"time"
)

var logger = xlog.New("").Sugar()
var defaultClient *Client

var debug bool

// Config used for autoload in shared.Config
type Config struct {
	URI string
	DB  string
	// seconds, 0 will keep mongo client default
	ConnectTimeout         int `json:",default=0"`
	Timeout                int `json:",default=0"`
	ServerSelectionTimeout int `json:",default=0"`
}
type Client struct {
	*mongo.Client
}

// SetDebug can open debug mode, output db query to log. you must set it before client is created.
func SetDebug() {
	debug = true
}

func NewClient(config Config) (*Client, error) {
	client, err := NewMongoClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewMongoClient(config Config) (*mongo.Client, error) {
	// create new client
	clientOpts := options.Client().ApplyURI(config.URI)
	if config.ConnectTimeout > 0 {
		clientOpts = clientOpts.SetConnectTimeout(time.Second * time.Duration(config.ConnectTimeout))
	}
	if config.Timeout > 0 {
		clientOpts = clientOpts.SetTimeout(time.Second * time.Duration(config.Timeout))
	}
	if config.ServerSelectionTimeout > 0 {
		clientOpts = clientOpts.SetServerSelectionTimeout(time.Second * time.Duration(config.ServerSelectionTimeout))
	}
	if debug {
		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				logger.Debugf(evt.Command.String())
			},
		}
		clientOpts = clientOpts.SetMonitor(monitor)
	}
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, err
	}

	// wait about 1 min, for ci process db starting
	for i := 0; i < 60; i += 1 {
		err = client.Ping(context.TODO(), readpref.Primary())
		if err != nil {
			logger.Infof("connect to db failed, retrying... error: %s", err)
			time.Sleep(time.Second)
			continue
		} else {
			return client, nil
		}
	}
	return nil, fmt.Errorf("connect to mongo db failed: %w", err)
}

func GetDefaultClient(config Config) (*Client, error) {
	if defaultClient != nil {
		return defaultClient, nil
	}
	// create new client
	var err error
	defaultClient, err = NewClient(config)
	if err != nil {
		return nil, err
	}
	return defaultClient, nil
}

func (c *Client) Database(name string, opts ...*options.DatabaseOptions) *Database {
	return &Database{c.Client.Database(name, opts...)}
}

// Close the default client
func Close(ctx context.Context) {
	err := defaultClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("disconnect from mongo db failed: %s\n", err)
	}
	defaultClient = nil
}
