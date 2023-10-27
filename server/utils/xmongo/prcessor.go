package xmongo

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func processInsertData(data interface{}) interface{} {
	v := reflect.Indirect(reflect.ValueOf(data))
	// only process struct pointer
	if !v.CanAddr() || v.Kind() != reflect.Struct {
		return data
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		var ft = t.Field(i)
		var fv = v.Field(i)
		if !fv.CanAddr() || !fv.CanInterface() {
			continue
		}
		tag, _ := bsoncodec.DefaultStructTagParser(ft) //nolint:all
		id := primitive.NewObjectID()
		if tag.Name == "_id" && ft.Type.AssignableTo(reflect.TypeOf(id)) && fv.IsZero() {
			fv.Set(reflect.ValueOf(id))
		}
		// if time.Time
		ot := time.Now()
		if tag.Name == "created_at" && ft.Type.AssignableTo(reflect.TypeOf(ot)) && fv.IsZero() {
			fv.Set(reflect.ValueOf(ot))
		}
		if tag.Name == "updated_at" && ft.Type.AssignableTo(reflect.TypeOf(ot)) && fv.IsZero() {
			fv.Set(reflect.ValueOf(ot))
		}
		// if primitive.DateTime
		now := primitive.NewDateTimeFromTime(ot)
		if tag.Name == "created_at" && ft.Type.AssignableTo(reflect.TypeOf(now)) && fv.IsZero() {
			fv.Set(reflect.ValueOf(now))
		}
		if tag.Name == "updated_at" && ft.Type.AssignableTo(reflect.TypeOf(now)) && fv.IsZero() {
			fv.Set(reflect.ValueOf(now))
		}
	}
	return v.Addr().Interface()
}

func processUpdateData(update interface{}) interface{} {
	// avoid conflict
	if strings.Contains(fmt.Sprint(update), "updated_at") {
		return update
	}
	// process bson.M
	v, ok := update.(bson.M)
	if !ok {
		// process bson.D
		d, ok := update.(bson.D)
		if !ok {
			return update
		}
		hit := false
		for i, e := range d {
			if e.Key == "$currentDate" {
				hit = true
				cd, ok := e.Value.(bson.D)
				if !ok {
					cm, ok := e.Value.(bson.M)
					if !ok {
						return update
					}
					cm["updated_at"] = true
					d[i].Value = cm
				}
				cd = append(cd, bson.E{Key: "updated_at", Value: true})
				d[i].Value = cd
			}
		}
		if !hit {
			d = append(d, bson.E{Key: "$currentDate", Value: bson.D{{Key: "updated_at", Value: true}}})
		}
		return d
	}
	current, ok := v["$currentDate"]
	if ok {
		sub, ok := current.(bson.M)
		if !ok {
			return update
		}
		sub["updated_at"] = true
		return v
	}
	v["$currentDate"] = bson.M{"updated_at": true}
	return v
}

// processUpsert add created_at field for doc which inserted before.
func (c *Collection) processUpsert(ctx context.Context, opts []*options.UpdateOptions, res *mongo.UpdateResult) {
	var opOK = false
	for _, opt := range opts {
		if opt.Upsert != nil && *opt.Upsert {
			opOK = true
		}
	}
	if opOK && res != nil && res.UpsertedID != nil {
		id, ok := res.UpsertedID.(primitive.ObjectID)
		if ok {
			u, err := c.Collection.UpdateByID(ctx, id,
				bson.M{"$set": bson.M{"created_at": primitive.NewDateTimeFromTime(id.Timestamp())}})
			if err != nil {
				logger.Errorf("add created_at when upsert failed: %s", err)
			}
			if u.ModifiedCount != 1 {
				logger.Errorf("add created_at when upsert failed. opts: %+v, res: %+v, res2: %+v", opts, res, u)
			}
		}
	}
}
