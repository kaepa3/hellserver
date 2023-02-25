package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

const (
	dbName  = "hellsee"
	colName = "Users"
	uri     = "mongodb://mongo:pass@127.0.0.1:27017/"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
}

type UserCollection struct {
	col *mongo.Collection
}

func GetUserCollection(ctx context.Context) (*UserCollection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	col := client.Database(dbName).Collection(colName)

	return &UserCollection{col: col}, err
}

// create new user
func (c *UserCollection) CreateUser(ctx context.Context, user *User) error {
	_, err := c.col.InsertOne(ctx, user)
	return err
}

// check add user
func (c *UserCollection) CountUser(ctx context.Context, user *bson.D) (int, error) {

	count, err := c.col.CountDocuments(ctx, user)
	return int(count), err
}

// search user
func (c *UserCollection) FindUser(ctx context.Context, filter *bson.D) (*User, error) {
	var result User
	r := c.col.FindOne(ctx, filter)
	err := r.Decode(&result)
	return &result, err
}
