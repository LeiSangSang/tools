package tools

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

type DbConfig struct {
	User string
	Pwd string
	Addr string
	Port string
	Table string
	Charset string
	MaxIdleConns int
	MaxOpenConns int
}

func InitMySql(config DbConfig) error {
	if len(config.Charset) == 0 {
		config.Charset = "utf8"
	}
	mysqlConfig := config.User + ":" + config.Pwd + "@tcp(" + config.Addr + ":" + config.Port + ")/" + config.Port + `?charset=` + config.Charset
	var err error
	Db, err = gorm.Open("mysql", mysqlConfig+"&parseTime=True&loc=Local")
	if err!=nil{
		return err
	}
	if config.MaxIdleConns != 0 {
		Db.DB().SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxOpenConns != 0 {
		Db.DB().SetMaxOpenConns(config.MaxOpenConns)
	}
	return nil
}
