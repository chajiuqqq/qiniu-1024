package xmongo

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	*mongo.Database
}

func DB(config Config) (*Database, error) {
	c, err := GetDefaultClient(config)
	if err != nil {
		return nil, err
	}
	logger.Infof("connect to mongo db %s successful", config.DB)
	return c.Database(config.DB), nil
}
