package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated_at"`
	DeletedAt          time.Time `xorm:"deleted_at"`
}

func (r *UserRepository) TableName() string {
	return "user_repository"
}
