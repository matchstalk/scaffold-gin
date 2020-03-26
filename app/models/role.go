package models

import "github.com/matchstalk/scaffold-gin/common/global"

type Role struct {
	global.Model
	ApplicationId string `sql:"type:varchar(64);index;comment:'应用ID'"`
	Name          string `sql:"type:varchar(100);comment:'名称'"`
	Description   string `sql:"type:text;comment:'描述'"`
	Remark        string `sql:"type:text;comment:'备注'"`
}
