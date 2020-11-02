package models

import (
	"app/config"
	"database/sql"
	"errors"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"time"
)

type Map map[string]func(env *config.Config) *sql.DB

var SQLGraph Map = Map{
	"sqlite": func(env *config.Config) *sql.DB {
		dsn := env.GetKey(`database`, `db_name`) + `.db`
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sql, _ := db.DB()
		return sql
	},
	"mysql": func(env *config.Config) *sql.DB {
		dsn := `root:123456789@tcp(127.0.0.1:3306)/cfadmin?charset=utf8mb4&parseTime=True&loc=Local`
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sql, _ := db.DB()
		return sql
	},
	"pgsql": func(env *config.Config) *sql.DB {
		dsn := `user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai`
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sql, _ := db.DB()
		return sql
	},
}

func CreateDB() *sql.DB {
	env := config.Default()
	db_type := env.GetKey(`database`, `db_driver`)
	if db_type == `` {
		panic(errors.New("Invalid db_type."))
	}
	f := SQLGraph[db_type]
	return f(env)
}

var Model *sql.DB = nil

func InitModels() {
	Model = CreateDB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Model.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Model.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Model.SetConnMaxLifetime(time.Hour)
}
