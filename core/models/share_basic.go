package models

import "time"

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string `xorm:"user_repository_identity"`
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"updated_at"`
	DeletedAt              time.Time `xorm:"deleted_at"`
}

func (s *ShareBasic) TableName() string {
	return "share_basic"
}
