package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(ctx context.Context) *Database {
	return &Database{ctx: ctx}
}

func (db *Database) Open(driverName, url string) error {
	var (
		d   *gorm.DB
		err error
	)
	switch driverName {
	case "mysql", "":
		d, err = gorm.Open(mysql.New(mysql.Config{
			DSN: url,
		}))
	default:
		return fmt.Errorf("unsupported database driver: %s", driverName)
	}
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	db.DB = d
	return nil
}

func (db *Database) Init() error {
	// Init Database
	initSqlSlice := parseInitSQL()
	if initSqlSlice == nil || len(initSqlSlice) == 0 {
		return nil
	}
	for _, s := range initSqlSlice {
		ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
		err := db.DB.WithContext(ctx).Exec(s).Error
		cancel()
		if err != nil {
			return fmt.Errorf("failed to init database: %w", err)
		}
	}
	return nil
}

func (db *Database) Close() error {
	return nil
}
