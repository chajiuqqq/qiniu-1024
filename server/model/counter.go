package model

type Counter struct {
	ID  string `bson:"_id" json:"_id"` // collection name
	Seq int64  `bson:"seq" json:"seq"` // max used number
}

func (c *Counter) Collection() string {
	return "counters"
}
