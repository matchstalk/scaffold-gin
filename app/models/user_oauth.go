package models

import "github.com/matchstalk/scaffold-gin/common/global"

type UserOauth struct {
	global.ModelTime
	UserId  string `sql:"primary_key;type:varchar(64);index;comment:'用户ID'"`
	OauthId string `sql:"primary_key;type:varchar(64);index;comment:'三方应用ID'"`
	Oauth   Oauth  `description:"三方应用"`
	Name    string `sql:"comment:'昵称'"`
	Openid  string `sql:"type:varchar(255);comment:'openid'"`
	Unionid string `sql:"type:varchar(255);comment:'unionid'"`
	Avatar  string `sql:"type:varchar(255);comment:'头像'"`
	Extra   string `sql:"type:varchar(255);comment:'额外字段'"`
}
