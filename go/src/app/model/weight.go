package model

import (
	// "github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v3"
	// "gopkg.in/guregu/null.v3/zero"
)

type Weight struct {
	Id        int         `json:"id" db:"id"`
	Weight    string      `json:"weight" db:"weight"`
	RegistDay null.Time   `json:"regist_day" db:"regist_day"`
	Meet      null.Int    `json:"meet" db:"meet"`
	Sports    null.Int    `json:"sports" db:"sports"`
	Memo      null.String `json:"memo" db:"memo"`
	UpdatedAt null.Time   `json:"updated_at" db:"updated_at"`
	CreatedAt null.Time   `json:"created_at" db:"created_at"`
}
