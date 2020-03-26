package models

import (
	"github.com/matchstalk/scaffold-gin/common/enums"
	"github.com/matchstalk/scaffold-gin/common/global"
	"time"
)

type User struct {
	global.Model
	Username         string       `sql:"type:varchar(100);index;comment:'用户名'"`
	Mobile           string       `sql:"type:varchar(50);comment:'手机号'"`
	Sex              enums.Sex    `sql:"comment:'性别'"`
	Email            string       `sql:"type:varchar(100);comment:'邮箱'"`
	Status           enums.Status `sql:"comment:'状态'"`
	DepartmentId     string       `sql:"type:varchar(64);index;comment:'部门ID'"`
	Title            string       `sql:"type:varchar(255);comment:'头衔'"`
	LastLoginTime    time.Time    `sql:"comment:'最后一次登录时间'"`
	FaceIcon         string       `sql:"comment:'头像'"`
	PasswordHash     string       `sql:"comment:'密码hash值'"`
	Salt             string       `sql:"comment:'密码加盐值'"`
	RestPasswordHash string       `sql:"comment:'重置密码hash值'"`
}
