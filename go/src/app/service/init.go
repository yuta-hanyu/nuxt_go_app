package service

import (
	"database/sql"
	"fmt"
	"goland/config"
	"goland/migration"
	"goland/model"

	"github.com/go-gorp/gorp"

	// "goland/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// gorp初期化処理
func InitDb() *gorp.DbMap {
	Db, err := sql.Open(
		config.Config.DbHost,
		fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
			config.Config.DbUserName,
			config.Config.DbPassWord,
			config.Config.DbHost,
			config.Config.DbPort,
			config.Config.DbName,
		),
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("conect database")
	}

	// テーブル作成
	migration.InitMigration(Db)

	dbmap := &gorp.DbMap{Db: Db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTableWithName(model.Weight{}, "weights").SetKeys(true, "Id")
	// err = dbmap.CreateTablesIfNotExists()
	// if err != nil {
	// 	fmt.Printf("error! can't craete table: %v", err)
	// }

	// ログの取得
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-nuxt-app:", log.Lmicroseconds))

	return dbmap
}
