package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/matchstalk/scaffold-gin/app/dto"
	"github.com/matchstalk/scaffold-gin/app/models"
)

type ApplicationDao struct {
	Dao
}

func NewApplicationDao(tx *gorm.DB) *ApplicationDao {
	return &ApplicationDao{
		Dao{
			Tx: tx,
		},
	}
}

func (e *ApplicationDao) Create(m *models.Application) *gorm.DB {
	e.ifOrmNil()
	return e.Tx.Create(&m)
}

func (e *ApplicationDao) Update(m *models.Application) *gorm.DB {
	e.ifOrmNil()
	return e.Tx.Model(m).Update(m)
}

func (e *ApplicationDao) Save(m *models.Application) *gorm.DB {
	e.ifOrmNil()
	return e.Tx.Save(m)
}

func (e *ApplicationDao) Read(g dto.GeneralGetDto) (m *models.Application) {
	e.ifOrmNil()
	e.Tx.First(m, g.Id)
	return
}

func (e *ApplicationDao) Delete(g dto.GeneralDelDto) *gorm.DB {
	e.ifOrmNil()
	return e.Tx.Delete(&models.Application{}, "id = ?", g.Id)
}

func (e *ApplicationDao) Query() *gorm.DB {
	e.ifOrmNil()
	return e.Tx.Model(&models.Application{})
}
