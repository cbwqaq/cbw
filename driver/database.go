package driver

import (
	"GINCHAT/pkg/setting"
	"gorm.io/gorm"
	"strconv"
)

var instance *gorm.DB

type DataSource interface {
	// Name returns the name of database
	Name() string
	// String returns the details of database
	String() string
	// Register registers the database which will be used
	Register(alias ...string) (*gorm.DB, error)
}

func InitDataSource() error {
	database, err := getDatabase()
	if err != nil {
		return err
	}
	instance, err = database.Register("")
	if err != nil {
		return err
	}

	return err
}

func DB() *gorm.DB {
	return instance
}

func getDatabase() (db DataSource, err error) {
	port, _ := strconv.Atoi(setting.RemoteSetting.MysqlPort)
	db = NewMySQL(
		setting.RemoteSetting.MysqlHost,
		setting.RemoteSetting.MysqlUserName,
		setting.RemoteSetting.MysqlPassWord,
		setting.RemoteSetting.MysqlDatabase,
		port,
	)

	return
}
