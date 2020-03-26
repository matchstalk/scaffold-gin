package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/matchstalk/scaffold-gin/common/gdb"
	"github.com/matchstalk/scaffold-gin/common/migrate"
	"time"
)

type Migration struct {
	db *gorm.DB
}

func (e *Migration) getDb() *gorm.DB {
	return gdb.GetDB()
}

func (e *Migration) AddMigrateVersion(version string) error {
	if e.db == nil {
		e.getDb()
	}
	m := migrate.Migrate{
		Version:   version,
		ApplyTime: time.Now(),
	}
	return e.db.Create(&m).Error
}

func (e *Migration) RemoveMigrateVersion(version string) error {
	if e.db == nil {
		e.getDb()
	}
	m := migrate.Migrate{
		Version:   version,
	}
	return e.db.Delete(&m).Error
}