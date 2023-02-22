package driver

import (
	"fmt"
	ms "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysql struct {
	host     string
	port     int
	usr      string
	pwd      string
	database string
}

func NewMySQL(host, usr, pwd, database string, port int) DataSource {
	return &mysql{
		host:     host,
		port:     port,
		usr:      usr,
		pwd:      pwd,
		database: database,
	}
}

func (m *mysql) Register(alias ...string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.usr, m.pwd, m.host, m.port, m.database)
	dail := ms.Open(dsn)
	return gorm.Open(dail, &gorm.Config{})
}

func (m *mysql) String() string {
	return fmt.Sprintf("type-%s host-%s port-%d user-%s database-%s",
		m.Name(), m.host, m.port, m.usr, m.database)
}

func (m *mysql) Name() string {
	return "MySQL"
}
