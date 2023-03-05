package model

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

const (
	uri = "mongodb://mongo:pass@127.0.0.1:27017/"
)
const (
	dbName        = "hellsee"
	userColName   = "user"
	trainColName  = "train"
	healthColName = "health"
)

var setUri = ""

func SetUri(v string) {
	setUri = v
}

func GetUri() string {
	if setUri != "" {
		return setUri
	}
	return uri
}

func getCollection(ctx context.Context, tp CollectionType) (*mongo.Collection, error) {
	name := typeToName(tp)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetUri()))
	if err != nil {
		return nil, err
	}
	col := client.Database(dbName).Collection(name)

	return col, nil
}

type CollectionType int

const (
	TypeUser CollectionType = iota
	TypeHealth
	TypeTrain
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
}

func typeToName(tp CollectionType) string {
	if tp == TypeUser {
		return userColName
	} else if tp == TypeHealth {
		return healthColName
	} else if tp == TypeTrain {
		return trainColName
	}
	return ""
}

func getUserCollection(ctx context.Context) (*mongo.Collection, error) {
	return getCollection(ctx, TypeUser)
}

// create new user
func CreateUser(ctx context.Context, user *User) error {
	col, err := getUserCollection(ctx)
	if err != nil {
		return err
	}
	_, err = col.InsertOne(ctx, user)
	return err
}

// check add user
func CountUser(ctx context.Context, user *bson.D) (int, error) {
	col, err := getUserCollection(ctx)
	if err != nil {
		return 0, err
	}
	count, err := col.CountDocuments(ctx, user)
	return int(count), err
}

// search user
func FindUser(ctx context.Context, filter *bson.D) (*User, error) {
	col, err := getUserCollection(ctx)
	if err != nil {
		return nil, err
	}
	var result User
	r := col.FindOne(ctx, filter)
	err = r.Decode(&result)
	return &result, err
}

func FindUsers(ctx context.Context, filter *bson.D) (*[]User, error) {
	col, err := getUserCollection(ctx)
	if err != nil {
		return nil, err
	}
	cur, err := col.Find(ctx, filter)
	if err != nil {
		log.Println("finder")
		return nil, err
	}
	defer cur.Close(ctx)

	var result []User
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type HealthCollection struct {
	col *mongo.Collection
}

// 健康情報の取得
func GetHealthCollection(ctx context.Context) (*HealthCollection, error) {
	c, err := getCollection(ctx, TypeHealth)
	if err != nil {
		return nil, err
	}
	return &HealthCollection{col: c}, err
}

type TrainType int

const (
	BenchPress TrainType = iota
)

type Train struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     string             `json:"userid"`
	TrainID    TrainType          `json:"trainid"`
	Weight     float64            `json:"weight"`
	Repetition int                `json:"repetition"`
}

func getTrainCollection(ctx context.Context) (*mongo.Collection, error) {
	return getCollection(ctx, TypeHealth)
}

func FindTrains(ctx context.Context, filter *bson.D) (*[]Train, error) {
	col, err := getTrainCollection(ctx)
	if err != nil {
		return nil, err
	}
	cur, e := col.Find(ctx, filter)
	if e != nil {
		return nil, e
	}
	defer cur.Close(ctx)

	var result []Train
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func AddTrain(ctx context.Context, train *Train) error {
	col, err := getTrainCollection(ctx)
	if err != nil {
		return err
	}
	_, err = col.InsertOne(ctx, train)
	return err
}
