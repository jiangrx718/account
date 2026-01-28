package utils

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMysqlFromViper() error {

	viper.SetDefault("mysql.dsn", "")
	viper.SetDefault("mysql.max_open_conns", 100)
	viper.SetDefault("mysql.max_idle_conns", 20)

	var err error
	dsn := viper.GetString("mysql.dsn")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	s, err := db.DB()
	if err != nil {
		return err
	}
	s.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	s.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	s.SetConnMaxIdleTime(time.Second * 5)
	s.SetConnMaxLifetime(time.Hour)

	return nil
}

func DefaultMysql() *gorm.DB {
	if db == nil {
		panic("get mysql client error")
	}
	return db
}

func NewMysql(dsn string, config *gorm.Config) *gorm.DB {
	if dsn == "" {
		viper.SetDefault("mysql.dsn", "")
		dsn = viper.GetString("mysql.dsn")
	}
	viper.SetDefault("mysql.max_open_conns", 100)
	viper.SetDefault("mysql.max_idle_conns", 20)

	if config == nil {
		config = &gorm.Config{}
	}
	client, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic(err)
	}

	s, err := client.DB()
	if err != nil {
		panic(err)
	}
	s.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	s.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	s.SetConnMaxIdleTime(time.Second * 5)
	s.SetConnMaxLifetime(time.Hour)

	if client == nil {
		panic("get mysql client error")
	}

	return client
}
