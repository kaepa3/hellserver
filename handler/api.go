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
	log.Println("comming handler")
	uid := userIDFromToken(c)
	ctx := context.Background()
	col, err := model.GetUserCollection(ctx)
	if err != nil {
		log.Println("not correction")
		return echo.ErrNotFound
	}

	id, _ := primitive.ObjectIDFromHex(uid)
	filter := bson.D{{Key: "_id", Value: id}}
	user, err := col.FindUser(ctx, &filter)
	if err != nil {
		log.Println(err)
	}
	log.Println(user.Name)

	users, err := col.FindUsers(ctx, &bson.D{})
	if err != nil {
		return echo.ErrNotFound
	}
	log.Println("ok")
	return c.JSON(http.StatusOK, users)
}

// get train data
func Train(c echo.Context) error {
	trains := []string{}
	return c.JSON(http.StatusOK, trains)
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
