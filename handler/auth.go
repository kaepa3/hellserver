package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kaepa3/hellserver/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Commin Health!")
}

func Signup(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	ctx := context.Background()
	filter := bson.D{{Key: "name", Value: user.Name}}

	col, err := model.GetUserCollection(ctx)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	if count, err := col.CountUser(ctx, &filter); err != nil && count != 0 {
		return errors.Join(&echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "count err",
		}, err)
	}

	user.ID = primitive.NewObjectID()
	if err := col.CreateUser(ctx, &user); err != nil {
		return errors.Join(err, &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "db error",
		})
	}

	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	c.Logger().Debug("login----------------------------------------")
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	ctx := context.Background()
	col, err := model.GetUserCollection(ctx)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "name", Value: u.Name}}
	user, err := col.FindUser(ctx, &filter)
	if err != nil {
		return errors.Join(&echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "user not found",
		}, err)
	}

	if user.ID.IsZero() || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}
	token, err := createToken(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

var signingKey = []byte("secret")

type TokenInfo struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func createToken(user *model.User) (string, error) {
	claims := &TokenInfo{
		user.ID.String(),
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
