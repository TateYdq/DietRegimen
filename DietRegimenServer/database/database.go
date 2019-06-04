package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var(
	DrDatabase *gorm.DB
)
func InitMysql(dsn string) error {
	db,err := gorm.Open("mysql",dsn)
	if err != nil{
		logrus.WithError(err).Errorf("connect err")
		return err
	}
	db.SingularTable(true)
	db.LogMode(true)
	DrDatabase = db
	return nil
}


