package migration

import (
	"database/sql"
	"fmt"
	// "goland/config"
	// "log"

	_ "github.com/go-sql-driver/mysql"
)

var err error

const (
	tableNameWeight = "weights"
)

// weightテーブル作成
func InitMigration(Db *sql.DB) {
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		weight VARCHAR(15) NOT NULL,
		regist_day datetime NOT NULL,
		meet int(11),
		sports int(11),
		memo varchar(255),
		updated_at datetime,
		created_at datetime default current_timestamp);`, tableNameWeight)
	Db.Exec(cmdU)
	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id int(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	// 		content varchar(255),
	// 		user_id int(11),
	// 		created_at datetime default current_timestamp);`, tableNameTodo)
	// Db.Exec(cmdT)

	// cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 			id int(255) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	// 			uuid VARCHAR(255) NOT NULL UNIQUE,
	// 			email varchar(255),
	// 			user_id int(11),
	// 			created_at datetime default current_timestamp);`, tableNameSession)

	// Db.Exec(cmdS)
}

// SQLの準備
// cmd := `insert into users (
// 	uuid,
// 	name,
// 	email,
// 	password,
// 	created_at) values (?, ?, ?, ?, ?)`

// 	_, err = Db.Exec(cmd,
// 		111,
// 		111,
// 		"111@111.com",
// 		"111",
// 		time.Now())

// if err != nil {
// 	log.Fatal(err)
// }
