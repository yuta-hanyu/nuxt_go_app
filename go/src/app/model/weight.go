package model

import (
	"time"
)

type Weight struct {
	Id    int  `json:"id" db:"id", primarykey, autoincrement`
	Weight  string  `json:"weight" db:"weight", notnull, size:16`
	Day   time.Time  `json:"day"  db:"day"`
	Meet   int  `json:"meet"  db:"meet"`
	Sports   int  `json:"sports"  db:"sports"`
	Memo   string  `json:"memo"  db:"memo"`
	CreatedAt   time.Time  `json:"created_at"  db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"  db:"updated_at"`
}