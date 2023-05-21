package db

import (
	"context"
	"gorm.io/gorm"
)

func getPgDbSession(db *gorm.DB, ctx context.Context) *gorm.DB {
	return db.Session(&gorm.Session{
		SkipHooks: true,
		Context:   ctx,
	})
}
