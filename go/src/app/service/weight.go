package service

import (
	// "goland/_example-mvc-api/model"
	// "fmt"
	"goland/model"
	"log"
	// "github.com/go-sql-driver/mysql"
)

type WeightService struct{}

func (WeightService) GetWeightList() ([]model.Weight, error) {
	log.Println(2222)

	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	var weights []model.Weight

	// ユーザーを全取得
	_, err := dbMap.Select(&weights, `SELECT * FROM weights`)
	if err != nil {
		return []model.Weight{}, err
	}

	// for weights {
	// 	log.Println(weights)
	// 	// var s mysql.NullTime
	// 	// err := weights.Scan(&s)
	// 	// errチェックをすること
	// 	// if s.Valid {
	// 	// 	// s.String を使う
	// 	// } else {
	// 	// 	// NULL値
	// 	// }
	// }

	return weights, nil
}

func (WeightService) CreateWeight(weight *model.Weight) error {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	// トラン ザクションを走らせながらinsert
	tx, _ := dbMap.Begin()

	err := tx.Insert(weight)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// func (UserService) UpdateUser(user *model.User) error {
// 	// initialize the DbMap
// 	dbMap := InitDb()
// 	defer dbMap.Db.Close()

// 	// トランザクションを走らせながらupdate
// 	tx, _ := dbMap.Begin()

// 	_, err := tx.Update(user)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	tx.Commit()

// 	return nil
// }

// func (UserService) DeleteUser(id int) error {
// 	// initialize the DbMap
// 	dbMap := InitDb()
// 	defer dbMap.Db.Close()

// 	// id から削除するユーザーを取得
// 	var user model.User
// 	err := dbMap.SelectOne(&user, `SELECT * FROM users WHERE id = :id`,
// 		map[string]interface{}{
// 			"id": id,
// 		})
// 	if err != nil {
// 		fmt.Printf("error! can't find user by id: %v.\n", id)
// 		return err
// 	}

// 	// トランザクションを走らせながらdelete
// 	tx, _ := dbMap.Begin()

// 	_, err := tx.Delete(&user)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	tx.Commit()

// 	return nil
// }
