package setups

import (
	"goland/controller/weights"

	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/mvc"
)

// ConfigureUsersControllers
// main.go で定義したURLのグループを表す変数（ここではusers）が引数に入り，
// それに対してDI（依存性注入）を行っている
func ConfigureWeightsControllers(app *mvc.Application) {

	// ログを取得してくれる機能のDIもここで行う
	app.Register(accesslog.GetFields)
	// URLが "/users" から始まるリクエストを受け取った際に，
	// 以下の Controllerを使用させるという指示
	app.Handle(new(weights.WeightController))
}
