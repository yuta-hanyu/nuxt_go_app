package main

import (
	"github.com/kataras/iris/v12"
	"goland/service"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"fmt"
)

func main() {
		
		
    app := iris.New()
		 // ログ記録（これも備え付きミドルウェア）
		 ac := accesslog.File("./log/access.log")
		 defer ac.Close()
		 app.UseRouter(ac.Handler)
		 app.UseRouter(recover.New())

		fmt.Println(service.InitDb())
		// return

    app.Handle("GET", "/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "ping"})
    })
	
    app.Listen(":8080")
}