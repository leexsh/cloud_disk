package models

import "time"

type UserBasic struct {
	Id        int64
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated_at"`
	DeletedAt time.Time `xorm:"deleted_at"`
}

func (user UserBasic) TableName() string {
	return "user_basic"
}
