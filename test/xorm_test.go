package test

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"

	"xorm.io/xorm"
)

func TestXromTest(t *testing.T) {

	engine, err := xorm.NewEngine("mysql", "root:Lizheng123@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8")
	defer engine.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGorm(t *testing.T) {
	dsn := "root:Lizheng123@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
}
