package database

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
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
	DrDatabase = db
	return nil
}


