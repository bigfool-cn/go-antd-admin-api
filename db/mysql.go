package db

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
	"wechat-bot-api/configs"
)

var Eloquent *gorm.DB

func init()  {
	var conn bytes.Buffer
	conn.WriteString(configs.Conf.Db.User)
	conn.WriteString(":")
	conn.WriteString(configs.Conf.Db.Password)
	conn.WriteString("@tcp(")
	conn.WriteString(configs.Conf.Db.Host)
	conn.WriteString(":")
	conn.WriteString(configs.Conf.Db.Port)
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(configs.Conf.Db.Database)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")

	var db Database
	db = new(Mysql)

	var err error
	Eloquent, err = db.Open(conn.String())

	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	} else {
		log.Println("mysql connect success!")
	}

	if Eloquent.Error != nil {
		log.Fatalf("database error %v", Eloquent.Error)
	}

	Eloquent.LogMode(configs.Conf.App.Debug)
}

type Database interface {
	Open(conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open("mysql", conn)
	return eloquent, err
}
