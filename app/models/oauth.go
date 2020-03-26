package models

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	"github.com/matchstalk/scaffold-gin/common/global"
)

type Oauth struct {
	global.ModelD
	Image       string       `sql:"type:text;comment:'缩略图'"`
	Name        string       `sql:"type:varchar(50);index;comment:'应用名称'"`
	Oauth       enums.Oauth  `sql:"type:char(20);comment:'三方应用类型'"`
	Description string       `sql:"type:text;comment:'描述'"`
	Version     string       `sql:"type:varchar(50);comment:'版本'"`
	Status      enums.Status `sql:"comment:'状态'"`
	AppId       string       `sql:"type:varchar(255);comment:'三方应用appId'"`
	AppSecret   string       `sql:"type:varchar(255);comment:'三方应用appSecret'"`
	Site        string       `sql:"type:varchar(255);comment:'官网地址'"`
	Extra       string       `sql:"type:varchar(255);comment:'额外字段'"`
}
