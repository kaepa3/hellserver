package main

import (
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

	e.Static("/assets", "public/assets")

	e.File("/", "public/index.html")
	e.File("/signup", "public/signup.html") // GET /signup
	e.File("/login", "public/login.html")   // GET /signup
	e.File("/health", "public/health.html") // GET /login

	e.POST("/signup", handler.Signup) // POST /signup
	e.POST("/login", handler.Login)   // POST /login

	api := e.Group("/api")
	api.GET("/users/:id", handler.Users)
	api.GET("/health/:id", handler.Health)
	api.GET("/train/:id", handler.Train)
	api.POST("/train", handler.AddTrain)
}
