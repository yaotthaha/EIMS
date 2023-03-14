package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	S "eims/sql"
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
	return db.DB.WithContext(db.ctx).AutoMigrate(&S.Department{}, &S.Employee{})
}

func (db *Database) Init() error {
	runCtx, runCancel := context.WithTimeout(db.ctx, 8*time.Second)
	err := db.DB.WithContext(runCtx).Exec(S.AddDefaultDepartment).Error
	runCancel()
	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}
	//
	runCtx, runCancel = context.WithTimeout(db.ctx, 8*time.Second)
	err = db.DB.WithContext(runCtx).Exec(S.AddDefaultEmployee).Error
	runCancel()
	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}
	//
	return nil
}

func (db *Database) Close() error {
	return nil
}
