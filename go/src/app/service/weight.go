package service

import (
	"goland/model"
	// "github.com/go-sql-driver/mysql"
)

type WeightService struct{}

func (WeightService) GetWeightList() ([]model.Weight, error) {
	// initialize the DbMap
	dbMap := InitDb()
	defer dbMap.Db.Close()

	var weights []model.Weight

	// ユーザーを全取得
	_, err := dbMap.Select(&weights, `SELECT * FROM weights`)
	if err != nil {
		return []model.Weight{}, err
	}

	return weights, err
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
