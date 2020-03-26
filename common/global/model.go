package global

import (
	"github.com/matchstalk/scaffold-gin/tools"
	"time"
)

type Model struct {
	ModelD
	DeletedAt *time.Time `sql:"index;comment:'删除时间'"`
}

func (e *ModelD) SetId(prefix string) {
	e.Id = tools.UUID18(prefix)
}

type ModelD struct {
	Id string `sql:"primary_key;type:varchar(64);comment:'主键'"`
	ModelTime
}

type ModelTime struct {
	CreatedAt time.Time `sql:"comment:'创建时间'"`
	UpdatedAt time.Time `sql:"comment:'更新时间'"`
}
