package main

import (
	"goland/controller/weights"

	"goland/service"

	"goland/setups"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	// "log"
	// "fmt"
	// "os"
)

func main() {
	app := iris.New()

	// ミドルウェアの使用
	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// ログ記録（これも備え付きミドルウェア）
	ac := accesslog.File("./access.log")
	defer ac.Close()

	// cors対応
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())

	service.InitDb()

	weightsAPI := app.Party("/weights")
	mw := mvc.New(weightsAPI)
	mw.Handle(new(weights.WeightController))

	mvc.Configure(weightsAPI, setups.ConfigureWeightsControllers)
	// ポートの指定
	app.Listen(":8080")

	// app.Handle("GET", "/", func(ctx iris.Context) {
	//     ctx.JSON(iris.Map{"message": "ping"})
	// })

	// app.Listen(":8080")
}
