package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mdsohelmia/go-kit/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewDatabase(cfg *config.Config) (*bun.DB, error) {
	driver := cfg.Database.Backend.Driver
	switch driver {
	case config.DatabaseDriverMysql:
		return newMysql(&cfg.Database.Backend.Mysql)
	case config.DatabaseDriverTidb:
		return newTidb(&cfg.Database.Backend.Tidb)
	default:
		return newMysql(&cfg.Database.Backend.Mysql)
	}
}

func newMysql(cfg *config.MysqlConfig) (*bun.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	sqldb, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxConnections)

	return db, nil
}

func newTidb(cfg *config.TidbConfig) (*bun.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	sqldb, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxConnections)

	return db, nil
}
