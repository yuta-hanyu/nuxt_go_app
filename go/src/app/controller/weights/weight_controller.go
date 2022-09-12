package weights

import (
	// "github.com/kataras/iris/v12"
	"goland/model"
	// "github.com/kataras/iris/v12/mvc"
	"goland/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type WeightController struct {
	// WeightServices service.WeightService
	// Ctx iris.Context
}

/*
* 体重一覧取得
* GET: http://localhost:8080/books
 */
func (c *WeightController) Get() ([]model.Weight, error) {
	dbMap := service.InitDb()
	defer dbMap.Db.Close()

	var weights []model.Weight
	// ユーザーを全取得
	_, err := dbMap.Select(&weights, `SELECT * FROM weights`)
	if err != nil {
		return []model.Weight{}, err
	}

	log.Println(weights)

	return weights, err
}

// // POST: http://localhost:8080/books
// func (c *BookController) Post(b Book) int {
// 	println("Received Book: " + b.Title)

// 	return iris.StatusCreated
// }
