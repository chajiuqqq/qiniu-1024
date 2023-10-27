package xmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is same ad mongo.Collection
// InsertOne and InsertMany will auto set _id, createdAt and updatedAt fields.
// Insert auto set only takes effect when struct is inserted.
// UpdateByID, UpdateOne and UpdateMany will auto set updatedAt field.
type Collection struct {
	*mongo.Collection
}

// Collection gets a handle for a collection with the given name configured with the given CollectionOptions.
func (db *Database) Collection(name string) *Collection {
	return &Collection{db.Database.Collection(name)}
}

// InsertOne wraps origin one for auto setting field tagged by `_id`,`createdAt`,`updatedAt`.
// Only takes effect when the parameter `document` is a pointer to struct.
//
// Origin InsertOne executes an insert command to insert a single document into the collection.
//
// The document parameter must be the document to be inserted. It cannot be nil. If the document does not have an _id
// field when transformed into BSON, one will be added automatically to the marshalled document. The original document
// will not be modified. The _id can be retrieved from the InsertedID field of the returned InsertOneResult.
//
// The opts parameter can be used to specify options for the operation (see the options.InsertOneOptions documentation.)
//
// For more information about the command, see https://docs.mongodb.com/manual/reference/command/insert/.
func (c *Collection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, processInsertData(document), opts...)
}

// InsertMany wraps origin one for auto setting field tagged by `_id`,`createdAt`,`updatedAt`.
// Only takes effect when the parameter `document` is an array of pointer to struct.
//
// Origin InsertMany executes an insert command to insert multiple documents into the collection. If write errors occur
// during the operation (e.g. duplicate key error), this method returns a BulkWriteException error.
//
// The documents parameter must be a slice of documents to insert. The slice cannot be nil or empty. The elements must
// all be non-nil. For any document that does not have an _id field when transformed into BSON, one will be added
// automatically to the marshalled document. The original document will not be modified. The _id values for the inserted
// documents can be retrieved from the InsertedIDs field of the returned InsertManyResult.
//
// The opts parameter can be used to specify options for the operation (see the options.InsertManyOptions documentation.)
//
// For more information about the command, see https://docs.mongodb.com/manual/reference/command/insert/.
func (c *Collection) InsertMany(ctx context.Context, document []interface{},
	opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	a := make([]interface{}, 0, len(document))
	for _, v := range document {
		a = append(a, processInsertData(v))
	}
	return c.Collection.InsertMany(ctx, a, opts...)
}

// UpdateByID wraps origin one for auto adding `updatedAt` to parameter `update` if necessary .
// If there is an upsert operation, will add `createdAt`.
//
// Origin UpdateByID executes an update command to update the document whose _id value matches the provided ID in the collection.
// This is equivalent to running UpdateOne(ctx, bson.D{{"_id", id}}, update, opts...).
//
// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
//
// The update parameter must be a document containing update operators
// (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be
// made to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.
func (c *Collection) UpdateByID(ctx context.Context, id interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := c.Collection.UpdateByID(ctx, id, processUpdateData(update), opts...)
	if err != nil {
		return nil, err
	}
	c.processUpsert(ctx, opts, res)
	return res, nil
}

// UpdateOne wraps origin one for auto adding `updatedAt` to parameter `update` if necessary .
// If there is an upsert operation, will add `createdAt`.
//
// Origin UpdateOne executes an update command to update at most one document in the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult
// with a MatchedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the
// matched set and MatchedCount will equal 1.
//
// The update parameter must be a document containing update operators
// (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be
// made to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.
func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := c.Collection.UpdateOne(ctx, filter, processUpdateData(update), opts...)
	if err != nil {
		return nil, err
	}
	c.processUpsert(ctx, opts, res)
	return res, nil
}

// UpdateMany wraps origin one for auto adding `updatedAt` to parameter `update` if necessary .
// If there is an upsert operation, will add `createdAt`.
//
// Origin UpdateMany executes an update command to update documents in the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the documents to be
// updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult
// with a MatchedCount of 0 will be returned.
//
// The update parameter must be a document containing update operators
// (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be made
// to the selected documents. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.
func (c *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{},
	opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := c.Collection.UpdateMany(ctx, filter, processUpdateData(update), opts...)
	if err != nil {
		return nil, err
	}
	c.processUpsert(ctx, opts, res)
	return res, nil
}

// FindOneAndUpdate wraps origin one for auto adding `updatedAt` to parameter `update` if necessary .
// If there is an upsert operation, will **not** add `createdAt`.
//
// FindOneAndUpdate executes a findAndModify command to update at most one document in the collection and returns the
// document as it appeared before updating.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// updated. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to
// ErrNoDocuments wil be returned. If the filter matches multiple documents, one will be selected from the matched set.
//
// The update parameter must be a document containing update operators
// (https://www.mongodb.com/docs/manual/reference/operator/update/) and can be used to specify the modifications to be made
// to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the options.FindOneAndUpdateOptions
// documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/findAndModify/.
func (c *Collection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return c.Collection.FindOneAndUpdate(ctx, filter, processUpdateData(update), opts...)
}
