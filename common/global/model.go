package global

import (
	"github.com/matchstalk/scaffold-gin/tools"
	"time"
)

type Model struct {
	Id        string     `sql:"primary_key;type:varchar(64);comment:'主键'"`
	CreatedAt time.Time  `sql:"comment:'创建时间'"`
	UpdatedAt time.Time  `sql:"comment:'更新时间'"`
	DeletedAt *time.Time `sql:"index;comment:'删除时间'"`
}

func (e *Model) SetId(prefix string) {
	e.Id = tools.UUID18(prefix)
}
