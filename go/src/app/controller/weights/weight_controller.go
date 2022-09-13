package weights

import (
	"goland/model"

	"goland/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type WeightController struct {
	WeightServices service.WeightService
	Ctx            iris.Context
}

/*
* 体重一覧取得
* GET: http://localhost:8080/books
 */
func (c *WeightController) Get() mvc.Response {

	weights, err := c.WeightServices.GetWeightList()
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError, // エラーハンドリング
		}
	}

	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{
		Code:   iris.StatusOK,
		Object: weights,
	}
}

/*
* 体重登録
* POST: http://localhost:8080/books
 */
func (c *WeightController) Post() mvc.Response {
	// リクエストボディのjsonデータを構造体（struct）に格納する
	var weight model.Weight
	err := c.Ctx.ReadJSON(&weight)

	// エラーハンドリング（Iris 備え付きのもので作れます）
	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 新規作成
	err = c.WeightServices.CreateWeight(&weight)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError, // エラーハンドリング
		}
	}

	// Iris 備え付きのレスポンス用構造体（struct）
	return mvc.Response{Code: iris.StatusCreated}
	// // 新規作成
	// cmd := `insert into weights (
	// 	weight,
	// 	regist_day,
	// 	meet,
	// 	sports,
	// 	memo) value (?, ?, ?, ?, ?)`
	// _, err = dbMap.Db.Exec(cmd, w.Weight, w.RegistDay, w.Meet, w.Sports, w.Memo)
	// if err != nil {
	// 	return mvc.Response{
	// 		Code: iris.StatusInternalServerError, // エラーハンドリング
	// 	}
	// }
}
