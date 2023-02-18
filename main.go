package main

import (
	"net/http"

	"github.com/kaepa3/hellserver/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	routing(e)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))

}
func routing(e *echo.Echo) {
	e.GET("/", hello)
	e.File("/signup", "public/signup.html") // GET /signup
	e.POST("/signup", handler.Signup)       // POST /signup
	e.File("/login", "public/login.html")   // GET /login
	e.POST("/login", handler.Login)         // POST /login

	api := e.Group("/api")
	api.POST("/users/New", hello)
	api.GET("/users/:id", hello)
	api.GET("/data/:id", hello)
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
