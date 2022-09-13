package model

import (
	// "github.com/go-sql-driver/mysql"
	"gopkg.in/guregu/null.v3"
	// "gopkg.in/guregu/null.v3/zero"
)

type Weight struct {
	Id        int         `json:"id" db:"id"`
	Weight    string      `json:"weight" db:"weight"`
	RegistDay null.Time   `json:"registDay" db:"regist_day"`
	Meet      null.Int    `json:"meet" db:"meet"`
	Sports    null.Int    `json:"sports" db:"sports"`
	Memo      null.String `json:"memo" db:"memo"`
	UpdatedAt null.Time   `json:"updatedAt" db:"updated_at"`
	CreatedAt null.Time   `json:"createdAt" db:"created_at"`
}
