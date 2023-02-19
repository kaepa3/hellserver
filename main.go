package main

import (
	"github.com/kaepa3/hellserver/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// インスタンスを作成
	e := echo.New()
	e.Debug = true

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	routing(e)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))

}
func routing(e *echo.Echo) {

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")
	e.File("/signup", "public/signup.html") // GET /signup
	e.POST("/signup", handler.Signup)       // POST /signup
	e.File("/login", "public/login.html")   // GET /login
	e.POST("/login", handler.Login)         // POST /login

	api := e.Group("/api")
	api.GET("/users/:id", handler.Hello)
	api.GET("/data/:id", handler.Hello)
}
