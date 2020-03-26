package models

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	"github.com/matchstalk/scaffold-gin/common/global"
)

type Department struct {
	global.Model
	Name        string       `sql:"type:varchar(50);index;comment:'应用名称'"`
	Description string       `sql:"type:text;comment:'描述'"`
	ParentId    string       `sql:"type:varchar(64);comment:'父级ID'"`
	Parent      *Department  `description:"父级部门"`
	ParentIds   string       `sql:"type:text;comment:'父级ID集'"`
	ChildIds    string       `sql:"type:text;comment:'子级ID集'"`
	LeaderId    string       `sql:"type:varchar(64);comment:'部门领导Id'"`
	Sort        uint8        `sql:"comment:'排序'"`
	Status      enums.Status `sql:"comment:'状态'"`
	Remark      string       `sql:"type:text;comment:'备注'"`
}
