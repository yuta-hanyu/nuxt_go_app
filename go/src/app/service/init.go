package service

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"goland/model"
	"goland/config"
	"github.com/go-gorp/gorp"
    
    _ "github.com/go-sql-driver/mysql"
)

// gorp初期化処理
func InitDb() *gorp.DbMap {

	fmt.Println(1111, config.Config.DbHost)
    dbHost := "mysql"
    db, err := sql.Open(
								"mysql",
								fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", 
														config.Config.DbUserName,
														config.Config.DbPassWord,
														// config.Config.DbHost,
														dbHost,
														config.Config.DbPort,
														config.Config.DbName,
													),
							)
    if err != nil {
        fmt.Println(err)
    } else {
			fmt.Println("Db接続完了")
		}
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    
   // テーブル作成
    dbmap.AddTableWithName(model.Weight{}, "weight").SetKeys(true, "Id")
    err = dbmap.CreateTablesIfNotExists()
    if err != nil {
        fmt.Printf("error! can't craete table: %v", err)
    }
    
    // ログの取得
    dbmap.TraceOn("[gorp]", log.New(os.Stdout, "go-iris-sample:", log.Lmicroseconds))
    
    return dbmap
}