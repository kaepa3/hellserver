package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kaepa3/hellserver/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Users(c echo.Context) error {
	ctx := context.Background()

	users, err := model.FindUsers(ctx, &bson.D{})
	if err != nil {
		return echo.ErrNotFound
	}
	log.Println("ok")
	return c.JSON(http.StatusOK, users)
}

func IsUserExist(uid string) bool {
	id, _ := primitive.ObjectIDFromHex(uid)
	filter := bson.D{{Key: "_id", Value: id}}
	ctx := context.Background()
	_, err := model.FindUser(ctx, &filter)
	return err == nil
}

// get train data
func Train(c echo.Context) error {
	uid := userIDFromToken(c)
	ctx := context.Background()

	trains, err := model.FindTrains(ctx, &bson.D{{Key: "userid", Value: uid}})
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, trains)
}

// addnew training
func AddTrain(c echo.Context) error {
	uid := userIDFromToken(c)

	train := new(model.Train)
	if err := c.Bind(train); err != nil {
		return err
	}
	train.UserID = uid
	ctx := context.Background()
	if err := model.AddTrain(ctx, train); err != nil {
		return err
	}
	return nil
}

// get health data
func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Commin Health!")
}

func userIDFromToken(c echo.Context) string {
	auth := c.Request().Header.Get("Authorization")

	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	token, err := jwt.ParseWithClaims(auth, &TokenInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AccessToken"), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	claims := token.Claims.(*TokenInfo)
	uid := claims.UID
	r := strings.NewReplacer("ObjectID(\"", "", "\")", "")
	uid = r.Replace(uid)
	log.Println(uid)
	return uid
}
