package database

import (
	"context"

	"gorm.io/gorm"
)

type Database struct {
	ctx context.Context
	*gorm.DB
}
