package models

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	"github.com/matchstalk/scaffold-gin/common/global"
)

type Application struct {
	global.ModelTime
	Id          string       `sql:"primary_key;type:varchar(50);comment:'应用名称'"`
	Domain      string       `sql:"type:varchar(255);comment:'访问地址'"`
	Version     string       `sql:"type:varchar(50);comment:'版本'"`
	Description string       `sql:"type:text;comment:'描述'"`
	Status      enums.Status `sql:"comment:'状态'"`
}
