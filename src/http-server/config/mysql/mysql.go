package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     int32
	Password string
	Driver   string
	DBName   string
	User     string
}

func (c *Config) GetClient() (*sql.DB, error) {
	driver := c.Driver
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName)
	log.Info("driver:", driver, ",dsn:", dsn)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
