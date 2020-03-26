package saas

import (
	"github.com/jinzhu/gorm"
	"sync"
)

var dbSaas *dbSet

type dbSet struct {
	dbs map[string]*gorm.DB
	mutex sync.Mutex
}

func GetDb(key string) (db *gorm.DB, ok bool) {
	if dbSaas == nil {
		return
	}
	dbSaas.mutex.Lock()
	db, ok = dbSaas.dbs[key]
	dbSaas.mutex.Unlock()
	return
}

func SetDb(key, driver, source string) error {
	db, err := gorm.Open(driver, source)
	if err != nil {
		return err
	}
	if dbSaas == nil {
		dbSaas = &dbSet{
			dbs: make(map[string]*gorm.DB),
		}
	}
	dbSaas.mutex.Lock()
	dbSaas.dbs[key] = db
	defer dbSaas.mutex.Unlock()
	return nil
}
