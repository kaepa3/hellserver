package model

import (
	"log"
	"os"
	"testing"

	"github.com/kaepa3/hellserver/testmod"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func TestMain(m *testing.M) {
	err, closeMongo, uri := testmod.SetupMongo(m)
	if err != nil {
		log.Fatal(err)
	}
	SetUri(uri)
	code := m.Run()
	// 後片付け
	closeMongo()
	// 終了
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	user := User{
		ID:       primitive.NewObjectID(),
		Name:     "kaepa3",
		Password: "hogehoge",
	}
	err := CreateUser(ctx, &user)
	if err != nil {
		t.Error(err)
	}
	count, err := CountUser(ctx, &bson.D{{Key: "name", Value: user.Name}})
	if err != nil {
		t.Error(err)
	} else if count == 0 {
		t.Error("num error:", count)
	}
}
