package models

type UserScope struct {
	UserId        string `sql:"primary_key;type:varchar(64);index;comment:'用户ID'"`
	ApplicationId string `sql:"primary_key;type:varchar(64);index;comment:'应用ID'"`
}
