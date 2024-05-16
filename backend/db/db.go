package db

import (
	"backend/env"
	"database/sql"
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
	sqlDB *sql.DB
}

// Connect connect to mysql
func Connect(cfg env.DBConfig) (db DB, err error) {
	db.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User, cfg.Passwd, cfg.IP, cfg.Port, cfg.DbName),
	}), &gorm.Config{})
	if err != nil {
		slog.Error(fmt.Sprintf("Fail to connect mysql, reason: %s", err.Error()))
		return
	}
	db.sqlDB, err = db.DB.DB()
	if err != nil {
		return
	}
	err = db.sqlDB.Ping()
	if err != nil {
		return
	}
	err = db.AutoMigrateAllTables()
	return
}

func (db DB) Close() (err error) {
	return db.sqlDB.Close()
}

func (db DB) AutoMigrateAllTables() error {
	return db.AutoMigrate(
		new(MediaSource),
		new(MediaFile),
		new(Camera),
	)
}

func (db DB) withTransaction(f func(DB) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		db.DB = tx // db is a copied value, so change db.DB have no side effect
		return f(db)
	})
}
