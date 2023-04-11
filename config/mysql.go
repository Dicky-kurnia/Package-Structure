package config

//import (
//	"boilerplate/exception"
//	"context"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"os"
//	"strconv"
//	"time"
//)
//
//func MysqlConnection() *gorm.DB {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	sqlDB, err := sql.Open("mysql", os.Getenv("MYSQL_HOST"))
//	exception.PanicIfNeeded(err)
//
//	err = sqlDB.PingContext(ctx)
//	exception.PanicIfNeeded(err)
//
//	mysqlPoolMax, err := strconv.Atoi(os.Getenv("MYSQL_POOL_MAX"))
//	exception.PanicIfNeeded(err)
//
//	mysqlIdleMax, err := strconv.Atoi(os.Getenv("MYSQL_IDLE_MAX"))
//	exception.PanicIfNeeded(err)
//
//	mysqlMaxLifeTime, err := strconv.Atoi(os.Getenv("MYSQL_MAX_LIFE_TIME_MINUTE"))
//	exception.PanicIfNeeded(err)
//
//	mysqlMaxIdleTime, err := strconv.Atoi(os.Getenv("MYSQL_MAX_IDLE_TIME_MINUTE"))
//	exception.PanicIfNeeded(err)
//
//	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
//	sqlDB.SetMaxIdleConns(mysqlIdleMax)
//
//	// SetMaxOpenConns sets the maximum number of open connections to the database.
//	sqlDB.SetMaxOpenConns(mysqlPoolMax)
//
//	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
//	sqlDB.SetConnMaxLifetime(time.Duration(mysqlMaxLifeTime) * time.Minute)
//
//	sqlDB.SetConnMaxIdleTime(time.Duration(mysqlMaxIdleTime) * time.Minute)
//
//	conf := &gorm.Config{}
//
//	if os.Getenv("MYSQL_LOG_QUERY") == "1" {
//		conf.Logger = logger.Default.LogMode(logger.Info)
//	}
//
//	gormDB, err := gorm.Open(mysql.New(mysql.Config{
//		Conn: sqlDB,
//	}), conf)
//
//	exception.PanicIfNeeded(err)
//	return gormDB
//}
